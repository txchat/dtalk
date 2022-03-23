package api

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/address"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/util"
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
			pubKey, err := VerifyAddress(sig)
			if err != nil {
				log.Debug().Err(err).Msg("VerifyAddress failed")
				context.Set(ReqError, err)
				context.Abort()
				return
			}
			addr := address.PublicKeyToAddress(address.NormalVer, pubKey)
			if addr == "" {
				log.Debug().Msg("PublicKeyToAddress addr is empty")
				context.Set(ReqError, err)
				context.Abort()
				return
			}
			//set val
			context.Set(Signature, sig)
			context.Set(Address, addr)
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

//get pubKey
func VerifyAddress(str string) ([]byte, error) {
	//<signature>#<msg>#<address>; <>
	ss := strings.SplitN(str, "#", -1)
	if len(ss) < 3 {
		log.Debug().Err(fmt.Errorf("need length:%v,got:%v", 3, len(ss))).Str("sig", str).Msg("split signature failed")
		return nil, xerror.NewError(xerror.SignatureInvalid)
	}
	sigData := ss[0]
	msgData := ss[1]
	pubKeyData := ss[2]

	msg := strings.SplitN(msgData, "*", -1)
	if len(msg) < 2 {
		log.Debug().Err(fmt.Errorf("need msg length:%v,got:%v", 2, len(ss))).Str("msgData", msgData).Msg("split msg data failed")
		return nil, xerror.NewError(xerror.SignatureInvalid)
	}
	time, err := strconv.ParseInt(msg[0], 10, 64)
	if err != nil {
		log.Debug().Err(err).Str("datetime", msg[0]).Msg("ParseInt datetime failed")
		return nil, xerror.NewError(xerror.SignatureInvalid)
	}
	//secp256
	sig, err := base64.StdEncoding.DecodeString(sigData)
	if err != nil {
		log.Debug().Err(err).Str("sigData", sigData).Msg("base64 decode sig data failed")
		return nil, xerror.NewError(xerror.SignatureInvalid)
	}
	pubKey, err := util.HexDecode(pubKeyData)
	if err != nil {
		log.Debug().Err(err).Str("pubKeyData", pubKeyData).Msg("hex decode pubKey failed")
		return nil, xerror.NewError(xerror.SignatureInvalid)
	}
	msg256 := sha256.Sum256([]byte(msgData))
	if !util.Secp256k1Verify(msg256[:], sig, pubKey) {
		log.Debug().Err(err).Str("msgData", msgData).Bytes("sig", sig).Bytes("pubKey", pubKey).
			Msg("Secp256k1Verify failed")
		return nil, xerror.NewError(xerror.SignatureInvalid)
	}
	//检查时间是否过期
	if util.CheckTimeOut(time, HeaderTimeOut) {
		log.Debug().Err(err).Int64("time", time).Msg("verify timeout")
		return nil, xerror.NewError(xerror.SignatureExpired)
	}
	return pubKey, nil
}
