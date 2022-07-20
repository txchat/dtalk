package service

// 组合各种数据访问来构建业务逻辑。

import (
	"time"

	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/service/group/model/biz"
	"github.com/txchat/dtalk/service/group/model/db"
)

func (s *Service) getNowTime() int64 {
	return time.Now().UnixNano() / 1e6
}

// CheckInGroup 查询并检查该成员是否在群中
func (s *Service) CheckInGroup(memberId string, groupId int64) (bool, error) {
	memberType, err := s.dao.GetMemberTypeMemberIdAndGroupId(memberId, groupId)
	if err != nil {
		return false, err
	}
	if memberType == biz.GroupMemberTypeOther {
		return false, nil
	}
	return true, nil
}

// InsertGroupMembers 批量插入群成员
func (s *Service) InsertGroupMembers(tx *mysql.MysqlTx, groupMembers []*db.GroupMember) error {
	err := s.dao.InsertGroupMembers(tx, groupMembers)
	if err != nil {
		return err
	}
	return nil
}
