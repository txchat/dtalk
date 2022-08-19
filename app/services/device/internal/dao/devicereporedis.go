package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/app/services/device/internal/model"
	xredis "github.com/txchat/dtalk/pkg/redis"
	xproto "github.com/txchat/imparse/proto"
)

const (
	_prefixDevice      = "device:%v"
	_prefixDeviceToken = "device_token:%v"
)

func keyDevice(uid string) string {
	return fmt.Sprintf(_prefixDevice, uid)
}

func keyDeviceToken(deviceToken string) string {
	return fmt.Sprintf(_prefixDeviceToken, deviceToken)
}

type DeviceRepositoryRedis struct {
	redis *redis.Pool
}

func NewDeviceRepositoryRedis(c xredis.Config) *DeviceRepositoryRedis {
	return &DeviceRepositoryRedis{
		redis: xredis.NewPool(c),
	}
}

//hash key:userId
//hash val: key=device val=json boj
func (repo *DeviceRepositoryRedis) AddDeviceInfo(device *model.Device) error {
	val, err := json.Marshal(device)
	if err != nil {
		return err
	}
	key := keyDevice(device.Uid)
	conn := repo.redis.Get()
	defer conn.Close()
	if err := conn.Send("HSET", key, device.ConnectId, val); err != nil {
		return fmt.Errorf("conn.Send(HSET %s,%s,%s) failed:%v", key, device.ConnectId, val, err)
	}
	//添加deviceToken-uid索引
	keyDT := keyDeviceToken(device.DeviceToken)
	if err := conn.Send("SET", keyDT, device.Uid); err != nil {
		return fmt.Errorf("conn.Send(SET %s,%s) failed:%v", keyDT, device.Uid, err)
	}
	if err := conn.Flush(); err != nil {
		return fmt.Errorf("conn.Flush() failed:%v", err)
	}
	n := 2
	for i := 0; i < n; i++ {
		if _, err := conn.Receive(); err != nil {
			return fmt.Errorf("conn.Receive() failed:%v", err)
		}
	}
	return repo.setExpire(device)
}

//hash key:userId
//hash val: key=device val=json boj
func (repo *DeviceRepositoryRedis) setExpire(device *model.Device) error {
	key := keyDevice(device.Uid)
	conn := repo.redis.Get()
	defer conn.Close()
	nMap, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return fmt.Errorf("conn.DO(HGETALL %s) failed:%v", key, err)
	}
	n := 0
	for _, val := range nMap {
		item := model.Device{}
		err := json.Unmarshal([]byte(val), &item)
		if err != nil {
			return err
		}
		if item.ConnectId != device.ConnectId {
			switch xproto.Device(item.DeviceType) {
			case xproto.Device_IOS, xproto.Device_Android:
				if xproto.Device(device.DeviceType) != xproto.Device_Android &&
					xproto.Device(device.DeviceType) != xproto.Device_IOS {
					continue
				}
			default:
				if item.DeviceType != device.DeviceType || item.DeviceUuid != device.DeviceUuid {
					continue
				}
			}
			n++
			if err := conn.Send("HDEL", keyDevice(item.Uid), item.ConnectId); err != nil {
				return fmt.Errorf("conn.Send(HDEL %s,%s) failed:%v", keyDevice(item.Uid), item.ConnectId, err)
			}
		}
	}
	if err := conn.Flush(); err != nil {
		return fmt.Errorf("conn.Flush() failed:%v", err)
	}
	for i := 0; i < n; i++ {
		if _, err := conn.Receive(); err != nil {
			return fmt.Errorf("conn.Receive() failed:%v", err)
		}
	}
	return nil
}

//hash key:userId
//hash val: key=device val=json boj
func (repo *DeviceRepositoryRedis) EnableDevice(uid, connId string) error {
	//读出Device
	key := keyDevice(uid)
	conn := repo.redis.Get()
	defer conn.Close()
	data, err := redis.Bytes(conn.Do("HGET", key, connId))
	if err != nil {
		return fmt.Errorf("conn.DO(HGET %s,%s) failed:%v", key, connId, err)
	}
	var device model.Device
	err = json.Unmarshal(data, &device)
	if err != nil {
		return err
	}
	device.IsEnabled = true

	//写回Device
	val, err := json.Marshal(device)
	if err != nil {
		return err
	}
	if err := conn.Send("HSET", key, device.ConnectId, val); err != nil {
		return fmt.Errorf("conn.Send(HSET %s,%s,%s) failed:%v", key, device.ConnectId, val, err)
	}
	if err := conn.Flush(); err != nil {
		return fmt.Errorf("conn.Flush() failed:%v", err)
	}
	n := 1
	for i := 0; i < n; i++ {
		if _, err := conn.Receive(); err != nil {
			return fmt.Errorf("conn.Receive() failed:%v", err)
		}
	}
	return nil
}

func (repo *DeviceRepositoryRedis) GetDevicesByConnID(uid, connID string) (*model.Device, error) {
	key := keyDevice(uid)
	conn := repo.redis.Get()
	defer conn.Close()
	jsonItem, err := redis.String(conn.Do("HGET", key, connID))
	if err != nil {
		return nil, fmt.Errorf("conn.DO(HGET %s %s) failed:%v", key, connID, err)
	}

	item := model.Device{}
	err = json.Unmarshal([]byte(jsonItem), &item)
	if err != nil {
		return nil, err
	}
	//获取deviceToken所属的uid
	item.DTUid, err = repo.GetTokenDevice(item.DeviceToken)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

//key:userId; val:json
func (repo *DeviceRepositoryRedis) GetAllDevices(uid string) ([]*model.Device, error) {
	key := keyDevice(uid)
	conn := repo.redis.Get()
	defer conn.Close()
	nMap, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return nil, fmt.Errorf("conn.DO(HGETALL %s) failed:%v", key, err)
	}
	devices := make([]*model.Device, 0)
	for _, val := range nMap {
		item := model.Device{}
		err := json.Unmarshal([]byte(val), &item)
		if err != nil {
			return nil, err
		}
		//获取deviceToken所属的uid
		item.DTUid, err = repo.GetTokenDevice(item.DeviceToken)
		if err != nil {
			continue
		}
		devices = append(devices, &item)
	}
	return devices, nil
}

//key:userId; val:json
func (repo *DeviceRepositoryRedis) GetTokenDevice(deviceToken string) (string, error) {
	key := keyDeviceToken(deviceToken)
	conn := repo.redis.Get()
	defer conn.Close()
	uid, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", fmt.Errorf("conn.DO(GET %s) failed:%v", key, err)
	}
	return uid, nil
}
