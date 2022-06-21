package dao

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/txchat/dtalk/service/device/model"
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

//hash key:userId
//hash val: key=device val=json boj
func (d *Dao) AddDeviceInfo(device *model.Device) error {
	val, err := json.Marshal(device)
	if err != nil {
		return err
	}
	key := keyDevice(device.Uid)
	conn := d.redis.Get()
	defer conn.Close()
	if err := conn.Send("HSET", key, device.ConnectId, val); err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.Send(HSET %s,%s,%s)", key, device.ConnectId, val))
		return err
	}
	//添加deviceToken-uid索引
	keyDT := keyDeviceToken(device.DeviceToken)
	if err := conn.Send("SET", keyDT, device.Uid); err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.Send(SET %s,%s)", keyDT, device.Uid))
		return err
	}
	if err := conn.Flush(); err != nil {
		d.log.Error().Err(err).Msg("conn.Flush()")
		return err
	}
	n := 2
	for i := 0; i < n; i++ {
		if _, err := conn.Receive(); err != nil {
			d.log.Error().Err(err).Msg("conn.Receive()")
			return err
		}
	}
	return d.setExpire(device)
}

//hash key:userId
//hash val: key=device val=json boj
func (d *Dao) setExpire(device *model.Device) error {
	key := keyDevice(device.Uid)
	conn := d.redis.Get()
	defer conn.Close()
	nMap, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.DO(HGETALL %s)", key))
		return err
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
				d.log.Error().Err(err).Msg(fmt.Sprintf("conn.Send(HDEL %s,%s)", keyDevice(item.Uid), item.ConnectId))
				return err
			}
		}
	}
	if err := conn.Flush(); err != nil {
		d.log.Error().Err(err).Msg("conn.Flush()")
		return err
	}
	for i := 0; i < n; i++ {
		if _, err := conn.Receive(); err != nil {
			d.log.Error().Err(err).Msg("conn.Receive()")
			return err
		}
	}
	return nil
}

//hash key:userId
//hash val: key=device val=json boj
func (d *Dao) EnableDevice(uid, connId string) error {
	//读出Device
	key := keyDevice(uid)
	conn := d.redis.Get()
	defer conn.Close()
	data, err := redis.Bytes(conn.Do("HGET", key, connId))
	if err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.DO(HGET %s,%s)", key, connId))
		return err
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
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.Send(HSET %s,%s,%s)", key, device.ConnectId, val))
		return err
	}
	if err := conn.Flush(); err != nil {
		d.log.Error().Err(err).Msg("conn.Flush()")
		return err
	}
	n := 1
	for i := 0; i < n; i++ {
		if _, err := conn.Receive(); err != nil {
			d.log.Error().Err(err).Msg("conn.Receive()")
			return err
		}
	}
	return nil
}

func (d *Dao) GetDevicesByConnID(uid, connID string) (*model.Device, error) {
	key := keyDevice(uid)
	conn := d.redis.Get()
	defer conn.Close()
	jsonItem, err := redis.String(conn.Do("HGET", key, connID))
	if err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.DO(HGET %s %s)", key, connID))
		return nil, err
	}

	item := model.Device{}
	err = json.Unmarshal([]byte(jsonItem), &item)
	if err != nil {
		return nil, err
	}
	//获取deviceToken所属的uid
	item.DTUid, err = d.GetTokenDevice(item.DeviceToken)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

//key:userId; val:json
func (d *Dao) GetAllDevices(uid string) ([]*model.Device, error) {
	key := keyDevice(uid)
	conn := d.redis.Get()
	defer conn.Close()
	nMap, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.DO(HGETALL %s)", key))
		return nil, err
	}
	devices := make([]*model.Device, 0)
	for _, val := range nMap {
		item := model.Device{}
		err := json.Unmarshal([]byte(val), &item)
		if err != nil {
			return nil, err
		}
		//获取deviceToken所属的uid
		item.DTUid, err = d.GetTokenDevice(item.DeviceToken)
		if err != nil {
			continue
		}
		devices = append(devices, &item)
	}
	return devices, nil
}

//key:userId; val:json
func (d *Dao) GetTokenDevice(deviceToken string) (string, error) {
	key := keyDeviceToken(deviceToken)
	conn := d.redis.Get()
	defer conn.Close()
	uid, err := redis.String(conn.Do("GET", key))
	if err != nil {
		d.log.Error().Err(err).Msg(fmt.Sprintf("conn.DO(GET %s)", key))
		return "", err
	}
	return uid, nil
}
