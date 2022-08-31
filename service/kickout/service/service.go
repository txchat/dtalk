package service

import (
	"context"
	"time"

	"github.com/robfig/cron"
	"github.com/rs/zerolog/log"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/interceptor/trace"
	"github.com/txchat/dtalk/pkg/slg"
	group "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/kickout/config"
	"google.golang.org/grpc"
)

type Service struct {
	GroupClient *group.Client
	SlgClient   *slg.Client

	closer context.CancelFunc
}

func New(c *config.Config) *Service {
	s := &Service{
		GroupClient: group.New(c.GroupRPCClient.RegAddrs,
			c.GroupRPCClient.Schema,
			c.GroupRPCClient.SrvName,
			time.Duration(c.GroupRPCClient.Dial),
			grpc.WithChainUnaryInterceptor(xerror.ErrClientInterceptor, trace.UnaryClientInterceptor),
		),
		SlgClient: slg.NewClient(slg.NewHTTPClient(c.SlgHTTPClient.Host, c.SlgHTTPClient.Salt)),
	}
	return s
}

func (s *Service) Shutdown(ctx context.Context) {
	down := make(chan struct{})
	go func() {
		s.closer()
		close(down)
	}()

	select {
	case <-ctx.Done():
		return
	case <-down:
		return
	}
}

func (s *Service) Run(ctx context.Context, spec string) {
	round := 0
	ctx, closer := context.WithCancel(ctx)
	s.closer = closer

	c := cron.New()
	err := c.AddFunc(spec, func() {
		select {
		case <-ctx.Done():
			log.Info().Int("last round", round).Msg("task stop")
			c.Stop()
			return
		default:
			// 定时任务处理逻辑
			round++
			s.kickOut(ctx, round)
		}
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}

func (s *Service) kickOut(ctx context.Context, round int) {
	roundLog := log.With().Int("round", round).Logger()
	//获取所有藏品群的信息
	resp, err := s.GroupClient.GetNFTGroupsExtInfo(ctx, &group.GetNFTGroupsExtInfoReq{})
	if err != nil {
		roundLog.Error().Err(err).Msg("GroupClient.GetNFTGroupsExtInfo")
		return
	}
	roundLog.Info().Int("total nft groups number", len(resp.Conditions)).Msg("")
	for i, condition := range resp.Conditions {
		groupId := condition.GetGroupId()
		roundLog = roundLog.With().Int64("groupId", groupId).Int("times", i).Logger()

		//获取群主uid
		gInfo, err := s.GroupClient.GetGroupInfo(ctx, &group.GetGroupInfoReq{
			GroupId: groupId,
		})
		if err != nil {
			roundLog.Error().Err(err).Msg("GroupClient.GetGroupInfo")
			continue
		}
		//获取群成员id
		members, err := s.GroupClient.GetMemberIds(ctx, groupId)
		if err != nil {
			roundLog.Error().Err(err).Msg("GroupClient.GetMemberIds")
			continue
		}
		//筛选出不满足入群条件的成员
		filteredMembers, err := s.getGroupNFTNotHandleMembers(gInfo.GetGroupOwnerId(), condition.GetCondition(), members)
		if err != nil {
			roundLog.Error().Err(err).Msg("getGroupNFTNotHandleMembers")
			continue
		}
		roundLog.Info().Strs("id", filteredMembers).Msg("need kick out members")
		resp2, err := s.GroupClient.GroupRemove(ctx, &group.GroupRemoveReq{
			GroupId:   groupId,
			PersonId:  gInfo.GetGroupOwnerId(),
			MemberIds: filteredMembers,
		})
		if err != nil {
			roundLog.Error().Err(err).Msg("GroupClient.GroupRemove")
			continue
		}
		roundLog.Info().Int32("now group members number", resp2.GetMemberNum()).Strs("kicked out members", resp2.GetMemberIds()).Msg("after kick out")
	}
}

// getGroupNFTNotHandleMembers 获取未持有相应藏品成员
func (s *Service) getGroupNFTNotHandleMembers(groupOwner string, condition *group.Condition, members []string) ([]string, error) {
	//请求上链购接口判断 conditionsRequest.GetType() conditionsRequest.GetNft()
	ids := make([]string, len(condition.GetNft()))
	for i, nft := range condition.GetNft() {
		ids[i] = nft.GetId()
	}
	conditions := make([]*slg.UserCondition, len(members))
	for i, tarId := range members {
		//群主无需判断
		if tarId == groupOwner {
			continue
		}
		item := &slg.UserCondition{
			UID:        tarId,
			HandleType: condition.GetType(),
			Conditions: ids,
		}
		conditions[i] = item
	}
	gps, err := s.SlgClient.LoadGroupPermission(conditions)
	if err != nil {
		return nil, err
	}
	filteredMembers := make([]string, 0)
	for _, memberId := range members {
		if !gps.IsPermission(memberId) {
			//群主无需判断
			if memberId == groupOwner {
				continue
			}
			filteredMembers = append(filteredMembers, memberId)
		}
	}
	return filteredMembers, nil
}
