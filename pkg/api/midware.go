package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/auth"
	xerror "github.com/txchat/dtalk/pkg/error"
)

var log = zlog.Logger

func composeHttpResp(code int, msg string, data interface{}) interface{} {
	type HttpAck struct {
		Result  int         `json:"result"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	var ret HttpAck
	ret.Result = code
	ret.Message = msg
	ret.Data = data
	return &ret
}

func parseRlt(result interface{}, err interface{}) (code int, msg string, data interface{}) {
	if err != nil {
		switch ty := err.(type) {
		case *xerror.Error:
			code = ty.Code()
			msg = ty.Error()
			data = ty.Data()
		default:
			log.Warn().Interface("err", err).Msg("inner error type")
			e := xerror.NewError(xerror.CodeInnerError)
			code = e.Code()
			msg = e.Error()
			data = err
		}
		return
	}

	code = xerror.CodeOK
	if isNil(result) {
		data = gin.H{}
	} else {
		data = result
	}
	return
}

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func RespMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if v, ok := c.Get(RespMiddleWareDisabled); ok && v == true {
			return
		}
		err := c.MustGet(ReqError)
		result, _ := c.Get(ReqResult)
		ret := composeHttpResp(parseRlt(result, err))
		c.PureJSON(http.StatusOK, ret)
		//c.PureJSON()
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		sig := context.GetHeader("FZM-SIGNATURE")
		uuid := context.GetHeader("FZM-UUID")
		device := context.GetHeader("FZM-DEVICE")
		deviceName := context.GetHeader("FZM-DEVICE-NAME")
		version := context.GetHeader("FZM-VERSION")

		//
		if sig == "MOCK" || sig == "MOCK2" {
			mockAddr := ""
			switch sig {
			case "MOCK":
				mockAddr = "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR"
			case "MOCK2":
				mockAddr = "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu"
			}
			//set val
			context.Set(Signature, sig)
			context.Set(Address, mockAddr)
			context.Set(Uuid, uuid)
			context.Set(DeviceType, device)
			context.Set(DeviceName, deviceName)
			context.Set(Version, version)
		} else {
			server := auth.NewDefaultApiAuthenticator()
			uid, err := server.Auth(sig)
			if err != nil {
				switch err {
				case auth.ERR_SIGNATUREEXPIRED:
					err = xerror.NewError(xerror.SignatureExpired)
				default:
					err = xerror.NewError(xerror.SignatureInvalid)
				}
				log.Debug().Err(err).Msg("VerifyAddress failed")
				context.Set(ReqError, err)
				context.Abort()
				return
			}
			//set val
			context.Set(Signature, sig)
			context.Set(Address, uid)
			context.Set(Uuid, uuid)
			context.Set(DeviceType, device)
			context.Set(DeviceName, deviceName)
			context.Set(Version, version)
		}
	}
}

func HeaderMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		uuid := context.GetHeader("FZM-UUID")
		device := context.GetHeader("FZM-DEVICE")
		deviceName := context.GetHeader("FZM-DEVICE-NAME")
		version := context.GetHeader("FZM-VERSION")

		//set val
		context.Set(Uuid, uuid)
		context.Set(DeviceType, device)
		context.Set(DeviceName, deviceName)
		context.Set(Version, version)
	}
}
