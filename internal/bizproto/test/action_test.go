package test

import (
	"bytes"
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/internal/bizproto"
	"github.com/txchat/im/api/protocol"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/proto/common"
)

type msg []byte

var sourceData = []msg{}

func TestMain(t *testing.M) {
	sourceData = genTestFrames()
	os.Exit(t.Run())
}

func genTestFrames() []msg {
	p := protocol.Proto{
		Ver:  0,
		Op:   int32(protocol.Op_SendMsg),
		Seq:  1,
		Ack:  0,
		Body: nil,
	}

	pro := common.Proto{
		EventType: common.Proto_common,
		Body:      nil,
	}

	comm := common.Common{
		ChannelType: common.Channel_ToUser,
		Mid:         0,
		Seq:         "client-msg",
		From:        "client-from",
		Target:      "client-target",
		MsgType:     common.MsgType_Text,
		Msg:         nil,
	}

	data2, err := proto.Marshal(&comm)
	if err != nil {
		panic(err)
	}
	pro.Body = data2

	data1, err := proto.Marshal(&pro)
	if err != nil {
		panic(err)
	}
	p.Body = data1

	data, err := proto.Marshal(&p)
	if err != nil {
		panic(err)
	}
	return []msg{data}
}

type testDB struct {
}

func (db *testDB) GetMsg(ctx context.Context, from, seq string) (*imparse.MsgIndex, error) {
	return nil, nil
}

func (db *testDB) AddMsg(ctx context.Context, uid string, m *imparse.MsgIndex) error {
	return nil
}
func (db *testDB) GetMid(ctx context.Context) (id int64, err error) {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	id = int64(rand.Intn(100))
	return
}

type testExec struct {
}

func (e *testExec) Transport(ctx context.Context, id int64, key, from, target string, ch imparse.Channel, frameType imparse.FrameType, data []byte) error {
	return nil
}

func (e *testExec) RevAck(ctx context.Context, id int64, keys []string, data []byte) error {
	return nil
}

func TestStandard_Filter(t *testing.T) {
	//全局初始化
	var db testDB
	var e testExec
	var exec = imparse.NewStandardAnswer(&db, &e, nil, nil)
	var parser bizproto.StandardParse

	//局部初始化
	//来源模式 1. ws推送；2. http推送；3. 测试命令行
	from := "server-client"
	key := "server-key"

	for _, m := range sourceData {
		ctx := context.Background()
		f, err := parser.NewFrame(key, from, bytes.NewReader(m))
		if err != nil {
			t.Error(err)
			return
		}
		_, err = exec.Filter(ctx, f)
		if err != nil {
			t.Error(err)
			continue
		}
		err = exec.Transport(ctx, f)
		if err != nil {
			t.Error(err)
			continue
		}
		_, err = exec.Ack(ctx, f)
		if err != nil {
			t.Error(err)
			continue
		}
	}
}

func BenchmarkCreateFrame(b *testing.B) {
	//全局初始化
	var db testDB
	var e testExec
	var exec = imparse.NewStandardAnswer(&db, &e, nil, nil)
	var parser bizproto.StandardParse

	data := sourceData[0]
	//局部初始化
	//来源模式 1. ws推送；2. http推送；3. 测试命令行
	from := "server-client"
	key := "server-key"

	for n := 0; n < b.N; n++ {
		ctx := context.Background()
		f, err := parser.NewFrame(key, from, bytes.NewReader(data))
		if err != nil {
			b.Error(err)
			return
		}
		_, err = exec.Filter(ctx, f)
		if err != nil {
			b.Error(err)
			continue
		}
		err = exec.Transport(ctx, f)
		if err != nil {
			b.Error(err)
			continue
		}
		_, err = exec.Ack(ctx, f)
		if err != nil {
			b.Error(err)
			continue
		}
	}
}
