package grpc

import (
	"context"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/group/logic"
	"time"

	"github.com/txchat/dtalk/pkg/interceptor/logger"
	"github.com/txchat/dtalk/pkg/interceptor/trace"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/types"
	"github.com/txchat/dtalk/service/group/service"
	"google.golang.org/grpc"
)

func New(cfg *xgrpc.ServerConfig, svr *service.Service) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(cfg.Timeout))
	logServerInterceptor := logger.NewServerInterceptor(svr.GetLog(), []string{
		"/dtalk.group.Group/GetGroupIds",
		"/dtalk.group.Group/CheckInGroup",
		"/dtalk.group.Group/GetMemberIds",
		"/dtalk.group.Group/CheckMute",
		"/dtalk.group.Group/GetGroups",
		"/dtalk.group.Group/GetMember",
		"/dtalk.group.Group/GetGroupInfo",
	})

	var ws = xgrpc.NewServer(
		cfg,
		connectionTimeout,
		grpc.ChainUnaryInterceptor(
			xerror.ErrInterceptor,
			trace.ServerUnaryInterceptor,
			logServerInterceptor.Unary,
		),
	)
	pb.RegisterGroupServer(ws.Server(), &server{svc: svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	pb.UnimplementedGroupServer
	svc *service.Service
}

// GetGroupIds .
func (s *server) GetGroupIds(ctx context.Context, req *pb.GetGroupIdsRequest) (*pb.GetGroupIdsReply, error) {
	memberId := req.MemberId
	groupIds, err := s.svc.GetGroupIdsByMemberId(memberId)
	if err != nil {
		return nil, err
	}

	return &pb.GetGroupIdsReply{GroupIds: groupIds}, nil
}

// CheckInGroup .
func (s *server) CheckInGroup(ctx context.Context, req *pb.CheckInGroupRequest) (*pb.CheckInGroupReply, error) {
	memberId := req.MemberId
	groupId := req.GroupId
	isOk, err := s.svc.CheckInGroup(memberId, groupId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckInGroupReply{IsOk: isOk}, nil
}

// GetMemberIds 单独开一个得到 id 的
func (s *server) GetMemberIds(ctx context.Context, req *pb.GetMemberIdsRequest) (*pb.GetMemberIdsReply, error) {
	groupId := req.GroupId
	groupMembers, err := s.svc.GetMembersByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	groupMemberIds := make([]string, len(groupMembers), len(groupMembers))
	for i := range groupMembers {
		groupMemberIds[i] = groupMembers[i].GroupMemberId
	}
	return &pb.GetMemberIdsReply{MemberIds: groupMemberIds}, nil
}

// CheckMute .
func (s *server) CheckMute(ctx context.Context, req *pb.CheckMuteRequest) (*pb.CheckMuteReply, error) {
	groupId := req.GroupId
	memberId := req.MemberId
	muteTime, err := s.svc.GetGroupMemberMuteTime(groupId, memberId)
	if err != nil {
		return nil, err
	}
	nowTime := time.Now().UnixNano() / 1e6
	res := &pb.CheckMuteReply{IsOk: false}
	if muteTime > nowTime {
		res.IsOk = true
	}
	return res, nil
}

// GetGroups 获得群列表
func (s *server) GetGroups(ctx context.Context, req *pb.GetGroupsReq) (*pb.GetGroupsResp, error) {
	getGroupsReq := &types.GetGroupListRequest{
		PersonId: req.Id,
	}

	getGroupsResp, err := s.svc.GetGroupListSvc(ctx, getGroupsReq)
	if err != nil {
		return nil, err
	}

	groupsInfo := make([]*pb.GroupInfo, len(getGroupsResp.Groups))
	for i, group := range getGroupsResp.Groups {
		groupsInfo[i] = &pb.GroupInfo{
			Id:     group.Id,
			Name:   group.Name,
			Avatar: group.Avatar,
		}
	}

	return &pb.GetGroupsResp{
		Groups: groupsInfo,
	}, nil

}

// GetMember .
func (s *server) GetMember(ctx context.Context, req *pb.GetMemberReq) (*pb.GetMemberResp, error) {
	groupId := req.GroupId
	memberId := req.MemberId

	mem, err := s.svc.GetMemberByMemberIdAndGroupId(ctx, memberId, groupId)
	if err != nil {
		if err == model.ErrRecordNotExist {
			return &pb.GetMemberResp{
				GroupMemberType: biz.GroupMemberTypeOther,
			}, nil
		}
		return nil, err
	}

	return &pb.GetMemberResp{
		GroupId:               mem.GroupId,
		GroupMemberId:         mem.GroupMemberId,
		GroupMemberName:       mem.GroupMemberName,
		GroupMemberType:       mem.GroupMemberType,
		GroupMemberJoinTime:   mem.GroupMemberJoinTime,
		GroupMemberUpdateTime: 0,
	}, nil
}

// ----------------- new ----------------

// GetGroupInfo 根据一个群 id 查询一个群信息
// oa 使用, 返回群信息和判断群是否被解散
func (s *server) GetGroupInfo(ctx context.Context, req *pb.GetGroupInfoReq) (*pb.GetGroupInfoResp, error) {
	groupId := req.GroupId

	group, err := s.svc.GetGroupInfoByGroupId(ctx, groupId)
	if err != nil {
		if err.Error() == xerror.NewError(xerror.GroupStatusDisBand).Error() {
			return &pb.GetGroupInfoResp{
				GroupId:      groupId,
				GroupExist:   false,
				GroupName:    "",
				GroupAvatar:  "",
				GroupOwnerId: "",
			}, nil
		} else {
			return nil, err
		}
	}

	return &pb.GetGroupInfoResp{
		GroupId:      group.GroupId,
		GroupExist:   true,
		GroupName:    group.GroupPubName,
		GroupAvatar:  group.GroupAvatar,
		GroupOwnerId: group.GroupOwnerId,
	}, nil
}

// ForceAddMember 增加群成员
// oa 专属
func (s *server) ForceAddMember(ctx context.Context, req *pb.ForceAddMemberReq) (*pb.ForceAddMemberResp, error) {
	l := logic.NewForceAddMemberLogic(ctx, s.svc)
	return l.ForceAddMember(req)
}

// ForceDeleteMember 删除群成员
// oa 专属, 退部门的时候退群
func (s *server) ForceDeleteMember(ctx context.Context, req *pb.ForceDeleteMemberReq) (*pb.ForceDeleteMemberResp, error) {
	l := logic.NewForceDeleteMemberLogic(ctx, s.svc)
	return l.ForceDeleteMember(req)
}

// ForceDisbandGroup .
func (s *server) ForceDisbandGroup(ctx context.Context, req *pb.ForceDisbandGroupReq) (*pb.ForceDisbandGroupResp, error) {
	err := s.svc.GroupDisband(ctx, req.GroupId, req.OpeId)
	if err != nil {
		return nil, err
	}

	return &pb.ForceDisbandGroupResp{}, nil
}

// ForceUpdateGroupType .
func (s *server) ForceUpdateGroupType(ctx context.Context, req *pb.ForceUpdateGroupTypeReq) (*pb.ForceUpdateGroupTypeResp, error) {
	err := s.svc.UpdateGroupType(req.GroupId, req.GroupType)
	if err != nil {
		return nil, err
	}

	return &pb.ForceUpdateGroupTypeResp{}, nil
}

// ForceAddMembers .
func (s *server) ForceAddMembers(ctx context.Context, req *pb.ForceAddMembersReq) (*pb.ForceAddMembersResp, error) {
	l := logic.NewForceAddMembersLogic(ctx, s.svc)
	return l.ForceAddMembers(req)
}

// ForceDeleteMembers .
func (s *server) ForceDeleteMembers(ctx context.Context, req *pb.ForceDeleteMembersReq) (*pb.ForceDeleteMembersResp, error) {
	l := logic.NewForceDeleteMembersLogic(ctx, s.svc)
	return l.ForceDeleteMembers(req)
}

// ForceJoinGroups .
func (s *server) ForceJoinGroups(ctx context.Context, req *pb.ForceJoinGroupsReq) (*pb.ForceJoinGroupsResp, error) {
	l := logic.NewForceJoinGroupsLogic(ctx, s.svc)
	return l.ForceJoinGroups(req)
}

// ForceExitGroups .
func (s *server) ForceExitGroups(ctx context.Context, req *pb.ForceExitGroupsReq) (*pb.ForceExitGroupsResp, error) {
	l := logic.NewForceExitGroupsLogic(ctx, s.svc)
	return l.ForceExitGroups(req)
}

// ForceChangeOwner .
func (s *server) ForceChangeOwner(ctx context.Context, req *pb.ForceChangeOwnerReq) (*pb.ForceChangeOwnerResp, error) {
	l := logic.NewForceChangeOwnerLogic(ctx, s.svc)
	return l.ForceChangeOwner(req)
}

// CreateGroup .
func (s *server) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	l := logic.NewCreateGroupLogic(ctx, s.svc)
	return l.CreateGroup(req)
}

// ChangeOwner .
func (s *server) ChangeOwner(ctx context.Context, req *pb.ChangeOwnerReq) (*pb.ChangeOwnerResp, error) {
	l := logic.NewChangeOwnerLogic(ctx, s.svc)
	return l.ChangeOwner(req)
}

func (s *server) GetGroupList(ctx context.Context, req *pb.GetGroupListReq) (*pb.GetGroupListResp, error) {
	l := logic.NewGetGroupListLogic(ctx, s.svc)
	return l.GetGroupList(req)
}

func (s *server) GetGroupMemberInfo(ctx context.Context, req *pb.GetGroupMemberInfoReq) (*pb.GetGroupMemberInfoResp, error) {
	l := logic.NewGetGroupMemberInfoLogic(ctx, s.svc)
	return l.GetGroupMemberInfo(req)
}

func (s *server) GetGroupMemberList(ctx context.Context, req *pb.GetGroupMemberListReq) (*pb.GetGroupMemberListResp, error) {
	l := logic.NewGetGroupMemberListLogic(ctx, s.svc)
	return l.GetGroupMemberList(req)
}

func (s *server) GetMuteList(ctx context.Context, req *pb.GetMuteListReq) (*pb.GetMuteListResp, error) {
	l := logic.NewGetMuteListLogic(ctx, s.svc)
	return l.GetMuteList(req)
}

func (s *server) GetPriGroupInfo(ctx context.Context, req *pb.GetPriGroupInfoReq) (*pb.GetPriGroupInfoResp, error) {
	l := logic.NewGetPriGroupInfoLogic(ctx, s.svc)
	return l.GetPriGroupInfo(req)
}

func (s *server) GetPubGroupInfo(ctx context.Context, req *pb.GetPubGroupInfoReq) (*pb.GetPubGroupInfoResp, error) {
	l := logic.NewGetPubGroupInfoLogic(ctx, s.svc)
	return l.GetPubGroupInfo(req)
}

func (s *server) GroupDisband(ctx context.Context, req *pb.GroupDisbandReq) (*pb.GroupDisbandResp, error) {
	l := logic.NewGroupDisbandLogic(ctx, s.svc)
	return l.GroupDisband(req)
}

func (s *server) GroupExit(ctx context.Context, req *pb.GroupExitReq) (*pb.GroupExitResp, error) {
	l := logic.NewGroupExitLogic(ctx, s.svc)
	return l.GroupExit(req)
}

func (s *server) GroupRemove(ctx context.Context, req *pb.GroupRemoveReq) (*pb.GroupRemoveResp, error) {
	l := logic.NewGroupRemoveLogic(ctx, s.svc)
	return l.GroupRemove(req)
}

func (s *server) InviteGroupMembers(ctx context.Context, req *pb.InviteGroupMembersReq) (*pb.InviteGroupMembersResp, error) {
	l := logic.NewInviteGroupMembersLogic(ctx, s.svc)
	return l.InviteGroupMembers(req)
}

func (s *server) SetAdmin(ctx context.Context, req *pb.SetAdminReq) (*pb.SetAdminResp, error) {
	l := logic.NewSetAdminLogic(ctx, s.svc)
	return l.SetAdmin(req)
}

func (s *server) UpdateGroupAvatar(ctx context.Context, req *pb.UpdateGroupAvatarReq) (*pb.UpdateGroupAvatarResp, error) {
	l := logic.NewUpdateGroupAvatarLogic(ctx, s.svc)
	return l.UpdateGroupAvatar(req)
}

func (s *server) UpdateGroupFriendType(ctx context.Context, req *pb.UpdateGroupFriendTypeReq) (*pb.UpdateGroupFriendTypeResp, error) {
	l := logic.NewUpdateGroupFriendTypeLogic(ctx, s.svc)
	return l.UpdateGroupFriendType(req)
}

func (s *server) UpdateGroupJoinType(ctx context.Context, req *pb.UpdateGroupJoinTypeReq) (*pb.UpdateGroupJoinTypeResp, error) {
	l := logic.NewUpdateGroupJoinTypeLogic(ctx, s.svc)
	return l.UpdateGroupJoinType(req)
}

func (s *server) UpdateGroupMemberMuteTime(ctx context.Context, req *pb.UpdateGroupMemberMuteTimeReq) (*pb.UpdateGroupMemberMuteTimeResp, error) {
	l := logic.NewUpdateGroupMemberMuteTimeLogic(ctx, s.svc)
	return l.UpdateGroupMemberMuteTime(req)
}

func (s *server) UpdateGroupMemberName(ctx context.Context, req *pb.UpdateGroupMemberNameReq) (*pb.UpdateGroupMemberNameResp, error) {
	l := logic.NewUpdateGroupMemberNameLogic(ctx, s.svc)
	return l.UpdateGroupMemberName(req)
}

func (s *server) UpdateGroupMuteType(ctx context.Context, req *pb.UpdateGroupMuteTypeReq) (*pb.UpdateGroupMuteTypeResp, error) {
	l := logic.NewUpdateGroupMuteTypeLogic(ctx, s.svc)
	return l.UpdateGroupMuteType(req)
}

func (s *server) UpdateGroupName(ctx context.Context, req *pb.UpdateGroupNameReq) (*pb.UpdateGroupNameResp, error) {
	l := logic.NewUpdateGroupNameLogic(ctx, s.svc)
	return l.UpdateGroupName(req)
}
