package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/service/discovery/model"
)

const (
	cNode = "find:chat_node"
	dNode = "find:chain33_node"
)

//key:name; val:json
func (d *Dao) SetCNode(key string, node *model.CNode) error {
	val, err := json.Marshal(node)
	if err != nil {
		return err
	}
	conn := d.redis.Get()
	defer conn.Close()
	if err := conn.Send("HSET", cNode, key, val); err != nil {
		d.log.Error(fmt.Sprintf("conn.Send(HSET %s,%s,%s)", cNode, key, val), "err", err)
		return err
	}
	if err := conn.Flush(); err != nil {
		d.log.Error("conn.Flush()", "err", err)
		return err
	}
	if _, err := conn.Receive(); err != nil {
		d.log.Error("conn.Receive()", "err", err)
		return err
	}
	return nil
}

func (d *Dao) GetCNodes() ([]*model.CNode, error) {
	conn := d.redis.Get()
	defer conn.Close()
	nMap, err := redis.StringMap(conn.Do("HGETALL", cNode))
	if err != nil {
		d.log.Error(fmt.Sprintf("conn.DO(HGETALL %s)", cNode), "err", err)
		return nil, err
	}
	nodes := make([]*model.CNode, 0)
	for _, v := range nMap {
		item := model.CNode{}
		err := json.Unmarshal([]byte(v), &item)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, &item)
	}
	return nodes, nil
}

//key:name; val:json
func (d *Dao) SetDNode(key string, node *model.DNode) error {
	val, err := json.Marshal(node)
	if err != nil {
		return err
	}
	conn := d.redis.Get()
	defer conn.Close()
	if err := conn.Send("HSET", dNode, key, val); err != nil {
		d.log.Error(fmt.Sprintf("conn.Send(HSET %s,%s,%s)", dNode, key, val), "err", err)
		return err
	}
	if err := conn.Flush(); err != nil {
		d.log.Error("conn.Flush()", "err", err)
		return err
	}
	if _, err := conn.Receive(); err != nil {
		d.log.Error("conn.Receive()", "err", err)
		return err
	}
	return nil
}

func (d *Dao) GetDNodes() ([]*model.DNode, error) {
	conn := d.redis.Get()
	defer conn.Close()
	nMap, err := redis.StringMap(conn.Do("HGETALL", dNode))
	if err != nil {
		d.log.Error(fmt.Sprintf("conn.DO(HGETALL %s)", dNode), "err", err)
		return nil, err
	}
	nodes := make([]*model.DNode, 0)
	for _, v := range nMap {
		item := model.DNode{}
		err := json.Unmarshal([]byte(v), &item)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, &item)
	}
	return nodes, nil
}
