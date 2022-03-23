package group

import (
	"context"
	"fmt"
	"time"

	"github.com/txchat/dtalk/pkg/naming"
	xgrpc "github.com/txchat/dtalk/pkg/net/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

type Client struct {
	conn   *grpc.ClientConn
	client GroupClient
}

func New(etcdAddr, schema, srvName string, dial time.Duration, opts ...grpc.DialOption) *Client {
	rb := naming.NewResolver(etcdAddr, schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", schema, srvName) // "schema://[authority]/service"
	fmt.Println("group rpc client call addr:", addr)

	conn, err := xgrpc.NewGRPCConnWithOpts(addr, dial, opts...)
	if err != nil {
		panic(err)
	}
	return &Client{
		conn:   conn,
		client: NewGroupClient(conn),
	}
}

func (c *Client) GetGroupIds(ctx context.Context, memberId string) ([]int64, error) {
	in := &GetGroupIdsRequest{
		MemberId: memberId,
	}
	res, err := c.client.GetGroupIds(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.GroupIds, nil
}

func (c *Client) CheckInGroup(ctx context.Context, memberId string, groupId int64) (bool, error) {
	in := &CheckInGroupRequest{
		MemberId: memberId,
		GroupId:  groupId,
	}
	res, err := c.client.CheckInGroup(ctx, in)
	if err != nil {
		return false, err
	}
	return res.IsOk, err
}

func (c *Client) GetMemberIds(ctx context.Context, groupId int64) ([]string, error) {
	in := &GetMemberIdsRequest{
		GroupId: groupId,
	}
	res, err := c.client.GetMemberIds(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.MemberIds, err
}

func (c *Client) CheckMute(ctx context.Context, memberId string, groupId int64) (bool, error) {
	in := &CheckMuteRequest{
		MemberId: memberId,
		GroupId:  groupId,
	}
	res, err := c.client.CheckMute(ctx, in)
	if err != nil {
		return false, err
	}
	return res.IsOk, nil
}

type Group struct {
	Id     int64
	Name   string
	Avatar string
}

func (c *Client) GetGroups(ctx context.Context, id string) ([]*Group, error) {
	in := &GetGroupsReq{
		Id: id,
	}
	res, err := c.client.GetGroups(ctx, in)
	if err != nil {
		return nil, err
	}
	groups := make([]*Group, len(res.Groups))
	for i, group := range res.Groups {
		groups[i] = &Group{
			Id:     group.Id,
			Name:   group.Name,
			Avatar: group.Avatar,
		}
	}

	return groups, nil
}

func (c *Client) GetMember(ctx context.Context, groupId int64, memberId string) (*GetMemberResp, error) {
	in := &GetMemberReq{
		MemberId: memberId,
		GroupId:  groupId,
	}
	res, err := c.client.GetMember(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ------------------ new proto client ----------------------------------

func (c *Client) ChangeOwner(ctx context.Context, req *ChangeOwnerReq) (*ChangeOwnerResp, error) {
	client := NewGroupClient(c.conn)
	return client.ChangeOwner(ctx, req)
}

func (c *Client) CreateGroup(ctx context.Context, req *CreateGroupReq) (*CreateGroupResp, error) {
	client := NewGroupClient(c.conn)
	return client.CreateGroup(ctx, req)
}

func (c *Client) GetGroupList(ctx context.Context, req *GetGroupListReq) (*GetGroupListResp, error) {
	client := NewGroupClient(c.conn)
	return client.GetGroupList(ctx, req)
}

func (c *Client) GetGroupMemberInfo(ctx context.Context, req *GetGroupMemberInfoReq) (*GetGroupMemberInfoResp, error) {
	client := NewGroupClient(c.conn)
	return client.GetGroupMemberInfo(ctx, req)
}

func (c *Client) GetGroupMemberList(ctx context.Context, req *GetGroupMemberListReq) (*GetGroupMemberListResp, error) {
	client := NewGroupClient(c.conn)
	return client.GetGroupMemberList(ctx, req)
}

func (c *Client) GetMuteList(ctx context.Context, req *GetMuteListReq) (*GetMuteListResp, error) {
	client := NewGroupClient(c.conn)
	return client.GetMuteList(ctx, req)
}

func (c *Client) GetPriGroupInfo(ctx context.Context, req *GetPriGroupInfoReq) (*GetPriGroupInfoResp, error) {
	client := NewGroupClient(c.conn)
	return client.GetPriGroupInfo(ctx, req)
}

func (c *Client) GetPubGroupInfo(ctx context.Context, req *GetPubGroupInfoReq) (*GetPubGroupInfoResp, error) {
	client := NewGroupClient(c.conn)
	return client.GetPubGroupInfo(ctx, req)
}

func (c *Client) GroupDisband(ctx context.Context, req *GroupDisbandReq) (*GroupDisbandResp, error) {
	client := NewGroupClient(c.conn)
	return client.GroupDisband(ctx, req)
}

func (c *Client) GroupExit(ctx context.Context, req *GroupExitReq) (*GroupExitResp, error) {
	client := NewGroupClient(c.conn)
	return client.GroupExit(ctx, req)
}

func (c *Client) GroupRemove(ctx context.Context, req *GroupRemoveReq) (*GroupRemoveResp, error) {
	client := NewGroupClient(c.conn)
	return client.GroupRemove(ctx, req)
}

func (c *Client) InviteGroupMembers(ctx context.Context, req *InviteGroupMembersReq) (*InviteGroupMembersResp, error) {
	client := NewGroupClient(c.conn)
	return client.InviteGroupMembers(ctx, req)
}

func (c *Client) SetAdmin(ctx context.Context, req *SetAdminReq) (*SetAdminResp, error) {
	client := NewGroupClient(c.conn)
	return client.SetAdmin(ctx, req)
}

func (c *Client) UpdateGroupAvatar(ctx context.Context, req *UpdateGroupAvatarReq) (*UpdateGroupAvatarResp, error) {
	client := NewGroupClient(c.conn)
	return client.UpdateGroupAvatar(ctx, req)
}

func (c *Client) UpdateGroupFriendType(ctx context.Context, req *UpdateGroupFriendTypeReq) (*UpdateGroupFriendTypeResp, error) {
	client := NewGroupClient(c.conn)
	return client.UpdateGroupFriendType(ctx, req)
}

func (c *Client) UpdateGroupJoinType(ctx context.Context, req *UpdateGroupJoinTypeReq) (*UpdateGroupJoinTypeResp, error) {
	client := NewGroupClient(c.conn)
	return client.UpdateGroupJoinType(ctx, req)
}

func (c *Client) UpdateGroupMemberMuteTime(ctx context.Context, req *UpdateGroupMemberMuteTimeReq) (*UpdateGroupMemberMuteTimeResp, error) {
	client := NewGroupClient(c.conn)
	return client.UpdateGroupMemberMuteTime(ctx, req)
}

func (c *Client) UpdateGroupMemberName(ctx context.Context, req *UpdateGroupMemberNameReq) (*UpdateGroupMemberNameResp, error) {
	client := NewGroupClient(c.conn)
	return client.UpdateGroupMemberName(ctx, req)
}

func (c *Client) UpdateGroupMuteType(ctx context.Context, req *UpdateGroupMuteTypeReq) (*UpdateGroupMuteTypeResp, error) {
	client := NewGroupClient(c.conn)
	return client.UpdateGroupMuteType(ctx, req)
}

func (c *Client) UpdateGroupName(ctx context.Context, req *UpdateGroupNameReq) (*UpdateGroupNameResp, error) {
	client := NewGroupClient(c.conn)
	return client.UpdateGroupName(ctx, req)
}
