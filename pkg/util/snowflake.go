package util

import (
	"errors"
	"sync"
	"time"
)

// +-----------------------------------------------------------------------+
// | UNUSED(1BIT) | TIMESTAMP(41BIT) | NODE-ID(10BIT) | SEQUENCE-ID(12BIT) |
// +-----------------------------------------------------------------------+

// snowflake算法，用于生成分布式系统下全局唯一id（64位）
// 1位符号位
// 41位时间戳（毫秒级），为当前时间戳-初始时间戳，大约能使用从初始时间戳开始至初始时间戳69年后
// 10位节点id
// 12位序列id

const (
	nodeBits       uint8 = 10                        // 节点id所占的位数，10位表示最多可以有2^10=1024个节点
	sequenceBits   uint8 = 12                        // 序列id所占的位数，12位表示每台机器1毫秒内最多可以生成2^12=4096个唯一id
	nodeMax        int64 = -1 ^ (-1 << nodeBits)     // 节点id的最大值，位运算用于防止溢出，默认为1023
	sequenceMax    int64 = -1 ^ (-1 << sequenceBits) // 序列id的最大值，位运算用于防止溢出，默认为4095
	timestampShift uint8 = nodeBits + sequenceBits   // 时间戳左移位数，默认为22
	nodeShift      uint8 = sequenceBits              // 节点id左移位数，默认为12
	epoch          int64 = 1590940800000             // 初始时间戳（默认为2020-06-01 00:00:00的毫秒级时间戳，正式使用后不可修改）
)

// Snowflake 定义一个Snowflake节点所需要的基本参数
type Snowflake struct {
	sync.Mutex       // 互斥锁，用于确保并发安全
	timestamp  int64 // 当前时间戳（毫秒级）
	nodeId     int64 // 当前节点的id
	sequenceId int64 // 当前毫秒下已经生成的序列id，默认在1毫秒内最多生成4096个
}

// NewSnowflake 实例化一个Snowflake节点
func NewSnowflake(nodeId int64, location ...*time.Location) (*Snowflake, error) {
	// 判断nodeId是否非法
	if nodeId < 0 || nodeId > nodeMax {
		return nil, errors.New("node id excess of quantity")
	}
	return &Snowflake{
		timestamp:  TimeNowUnixMilli(location...),
		nodeId:     nodeId,
		sequenceId: 0,
	}, nil
}

// NextId 获取唯一id
func (s *Snowflake) NextId(location ...*time.Location) int64 {
	s.Lock()
	defer s.Unlock()

	// 获取当前时间的毫秒级时间戳
	now := TimeNowUnixMilli(location...)

	// 发生时钟回拨时，等待时间重回timestamp记录的时间
	if s.timestamp > now {
		for now <= s.timestamp {
			SleepMilli(s.timestamp - now)
			now = TimeNowUnixMilli(location...)
		}
	}

	if s.timestamp == now {
		s.sequenceId++
		// 判断当前节点是否在1毫秒内生成了sequenceMax个id
		if s.sequenceId > sequenceMax {
			// 如果当前节点在1毫秒内生成的id超过了上限，则等待1毫秒后再继续生成id
			for now <= s.timestamp {
				SleepMilli(1)
				now = TimeNowUnixMilli(location...)
			}
		}
	} else {
		// 如果当前时间戳与节点上一次生成id的时间戳不一致，则重置节点的序列id
		s.sequenceId = 0
		s.timestamp = now // 将当前节点上一次生成id的时间戳更新为当前时间戳
	}
	id := int64((now-epoch)<<timestampShift | (s.nodeId << nodeShift) | (s.sequenceId))
	return id
}
