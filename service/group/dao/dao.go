package dao

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/api/trace"
	"github.com/txchat/dtalk/pkg/mysql"
	"github.com/txchat/dtalk/pkg/rand"
	"github.com/txchat/dtalk/pkg/redis"
	"github.com/txchat/dtalk/service/group/config"
)

var srvName = "group/dao"

type Dao struct {
	log  zerolog.Logger
	conn *mysql.MysqlConn
	// redis 连接
	redis *redis.Pool
	// redis 过期时间
	redisExpire time.Duration
}

func New(c *config.Config) *Dao {
	d := &Dao{
		log:         zlog.Logger.With().Str("service", srvName).Logger(),
		conn:        newDB(c.MySQL),
		redis:       redis.New(c.Redis),
		redisExpire: time.Duration(c.Redis.Expire),
	}
	//log init
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Conf.Env == "debug" {
		d.log = zlog.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Str("service", srvName).Logger()
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if config.Conf.Env == "benchmark" {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	return d
}

func newDB(cfg *config.MySQL) *mysql.MysqlConn {
	c, err := mysql.NewMysqlConn(cfg.Host, fmt.Sprintf("%v", cfg.Port),
		cfg.User, cfg.Pwd, cfg.Db, "UTF8MB4")
	if err != nil {
		panic(err)
	}
	return c
}

func (d *Dao) NewTx() (*mysql.MysqlTx, error) {
	return d.conn.NewTx()
}

func (d *Dao) GetLogWithTrace(ctx context.Context) zerolog.Logger {
	logId := d.GetTrace(ctx)
	return d.log.With().Str("trace", logId).Logger()
}

func (d *Dao) GetTrace(ctx context.Context) string {
	return trace.NewTraceIdWithContext(ctx)
}

func (d *Dao) GetRandRedisExpire() time.Duration {
	return (time.Duration(rand.RandInt(0, 100))*time.Second + time.Duration(d.redisExpire)) / time.Second
}

func (d *Dao) getNowTime() int64 {
	return time.Now().UnixNano() / 1e6
}
