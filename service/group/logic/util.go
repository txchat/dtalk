package logic

import (
	"strings"

	xerror "github.com/txchat/dtalk/pkg/error"
	pb "github.com/txchat/dtalk/service/group/api"
	"github.com/txchat/dtalk/service/group/model/biz"
)

func FilteredMemberIds(memberIds []string) []string {
	newMemberIds := make([]string, 0, len(memberIds))
	for _, memberId := range memberIds {
		memberId, err := FilteredMemberId(memberId)
		if err != nil {
			continue
		}

		newMemberIds = append(newMemberIds, memberId)
	}

	return newMemberIds
}

func FilteredMemberId(memberId string) (string, error) {
	memberId = strings.TrimSpace(memberId)
	if memberId != "" && len(memberId) < 40 {
		return memberId, nil
	}
	return "", xerror.NewError(xerror.ParamsError)
}

func NewRPCGroupInfo(do *biz.GroupInfo) *pb.GroupBizInfo {
	if do == nil {
		return &pb.GroupBizInfo{}
	}

	return &pb.GroupBizInfo{
		Id:            do.GroupId,
		MarkId:        do.GroupMarkId,
		Name:          do.GroupName,
		Avatar:        do.GroupAvatar,
		MemberNum:     do.GroupMemberNum,
		MemberMaximum: do.GroupMaximum,
		Introduce:     do.GroupIntroduce,
		Status:        pb.GroupStatus(do.GroupStatus),
		OwnerId:       do.GroupOwnerId,
		CreateTime:    do.GroupCreateTime,
		UpdateTime:    do.GroupUpdateTime,
		JoinType:      pb.GroupJoinType(do.GroupJoinType),
		MuteType:      pb.GroupMuteType(do.GroupMuteType),
		FriendType:    pb.GroupFriendType(do.GroupFriendType),
		MuteNum:       do.MuteNum,
		AdminNum:      do.AdminNum,
		AESKey:        do.AESKey,
		PubName:       do.GroupPubName,
		Type:          pb.GroupType(do.GroupType),
		Owner:         nil,
		Person:        nil,
		Members:       nil,
	}
}

func NewRPCGroupInfos(dos []*biz.GroupInfo) []*pb.GroupBizInfo {
	dtos := make([]*pb.GroupBizInfo, 0, len(dos))
	for _, do := range dos {
		dtos = append(dtos, NewRPCGroupInfo(do))
	}

	return dtos
}

func NewRPCGroupMemberInfo(do *biz.GroupMember) *pb.GroupMemberBizInfo {
	if do == nil {
		return &pb.GroupMemberBizInfo{}
	}

	return &pb.GroupMemberBizInfo{
		GroupId:  do.GroupId,
		Id:       do.GroupMemberId,
		Name:     do.GroupMemberName,
		Type:     pb.GroupMemberType(do.GroupMemberType),
		MuteTime: do.GroupMemberMuteTime,
		JoinTime: do.GroupMemberJoinTime,
	}
}

func NewRPCGroupMemberInfos(dos []*biz.GroupMember) []*pb.GroupMemberBizInfo {
	dtos := make([]*pb.GroupMemberBizInfo, 0, len(dos))
	for _, do := range dos {
		dtos = append(dtos, NewRPCGroupMemberInfo(do))
	}

	return dtos
}
