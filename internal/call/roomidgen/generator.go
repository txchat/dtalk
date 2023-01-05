package roomidgen

import (
	"context"
	"sync"
)

// RoomIDGenerator roomId 生成器, 获得临时唯一的 roomId
type RoomIDGenerator struct {
	// TODO 换个更科学的
	id   int32
	max  int32
	node int32
	sync.Mutex
}

func NewRoomIDGen(node int32) *RoomIDGenerator {
	if node > 9 {
		node = 9
	}
	room := &RoomIDGenerator{
		id:  0,
		max: 1000000,
	}
	room.node = node * room.max
	return room
}

// GetID 递增 1 获得 roomId, 到达 max 后取模归零
func (r *RoomIDGenerator) GetID(ctx context.Context) (int64, error) {
	r.Lock()
	defer r.Unlock()
	r.id = (r.id + 1) % r.max
	return int64(r.node + r.id), nil
}
