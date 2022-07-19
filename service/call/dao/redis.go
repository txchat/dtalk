package dao

import (
	"fmt"

	"github.com/txchat/dtalk/service/call/model"
)

const (
	_prefixCallSession = "session:%d"
	_prefixCallStatus  = "status:%s"
)

func keySession(traceId int64) string {
	return fmt.Sprintf(_prefixCallSession, traceId)
}

func keyStatus(userId string) string {
	return fmt.Sprintf(_prefixCallStatus, userId)
}

// GetSession get session from traceId
func (d *Dao) GetSession(traceId int64) (*model.Session, error) {
	key := keySession(traceId)
	if ok, err := d.redis.Exists(key); err != nil {
		return nil, err
	} else if !ok {
		return nil, model.ErrSessionNotExist
	}

	res := &model.Session{}
	if err := d.redis.Read(key, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Dao) SaveSession(session *model.Session) error {
	key := keySession(session.TraceId)
	if err := d.redis.Write(key, session, model.SESSIONMAXTIME); err != nil {
		return err
	}
	return nil
}

// 弃用
func (d *Dao) GetStatus(userId string) (int, error) {
	d.redis.RLock()
	defer d.redis.RUnlock()

	key := keyStatus(userId)
	if ok, err := d.redis.Exists(key); err != nil {
		return 0, err
	} else if !ok {
		return 0, nil
	}

	exist, err := d.redis.GetInt(key)
	if err != nil {
		return 0, err
	}
	return exist, nil
}

// 弃用
func (d *Dao) SetStatus(userId string, isBusy int) error {
	d.redis.Lock()
	defer d.redis.Unlock()

	key := keyStatus(userId)
	if err := d.redis.Set(key, isBusy); err != nil {
		return err
	}
	return nil
}
