package record

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/txchat/dtalk/gateway/api/v1/internal/logic/record"
	"github.com/txchat/dtalk/gateway/api/v1/internal/svc"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	comet "github.com/txchat/im/api/comet/grpc"
	"io/ioutil"
)

// PushToUid
// @Summary 推送消息
// @Description comet.Proto由接口组装，客户端只需传入comet.Proto的body部分
// @Author dld@33.cn
// @Tags record 消息模块
// @Accept       mpfd
// @Produce      json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param message body string true  "消息协议序列化"
// @Success 200 {object} model.GeneralResponse{}
// @Router	/record/push [post]
func PushToUid(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("MultipartForm"+err.Error()))
			return
		}
		files := form.File["message"]

		if len(files) < 1 {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("file len less than 1"))
			return
		}

		//file, err := c.FormFile("")
		file := files[0]
		f, err := file.Open()
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("Open File"+err.Error()))
			return
		}
		defer f.Close()

		body, err := ioutil.ReadAll(f)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ReadAll"+err.Error()))
			return
		}
		if len(body) == 0 {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("error message length 0"))
			return
		}
		p := comet.Proto{
			Ver:  1,
			Op:   int32(comet.Op_SendMsg),
			Seq:  0,
			Ack:  0,
			Body: body,
		}
		data, err := proto.Marshal(&p)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.SendMsgFailed).SetExtMessage(err.Error()))
			return
		}
		uid := c.MustGet(api.Address).(string)
		l := record.NewLogic(c.Request.Context(), ctx)
		mid, createTime, err := l.Push("", uid, data)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.SendMsgFailed).SetExtMessage(err.Error()))
			return
		}
		ret := map[string]interface{}{
			"logId":    mid,
			"datetime": createTime,
		}
		c.Set(api.ReqResult, ret)
		c.Set(api.ReqError, nil)
	}
}

// PushToUid2
// @Summary 推送消息2
// @Description comet.Proto由客户端传入
// @Author dld@33.cn
// @Tags record 消息模块
// @Accept       mpfd
// @Produce      json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param message body string true  "消息协议序列化"
// @Success 200 {object} model.GeneralResponse{}
// @Router	/record/push2 [post]
func PushToUid2(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("MultipartForm"+err.Error()))
			return
		}
		files := form.File["message"]

		if len(files) < 1 {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("file len less than 1"))
			return
		}

		//file, err := c.FormFile("")
		file := files[0]
		f, err := file.Open()
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("Open File"+err.Error()))
			return
		}
		defer f.Close()

		body, err := ioutil.ReadAll(f)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("ReadAll"+err.Error()))
			return
		}
		if len(body) == 0 {
			c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage("error message length 0"))
			return
		}
		uid := c.MustGet(api.Address).(string)
		l := record.NewLogic(c.Request.Context(), ctx)
		mid, createTime, err := l.Push("", uid, body)
		if err != nil {
			c.Set(api.ReqError, xerror.NewError(xerror.SendMsgFailed).SetExtMessage(err.Error()))
			return
		}
		ret := map[string]interface{}{
			"logId":    mid,
			"datetime": createTime,
		}
		c.Set(api.ReqResult, ret)
		c.Set(api.ReqError, nil)
	}
}
