package logger

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/txchat/dtalk/pkg/logger"

	"github.com/txchat/dtalk/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	xerror "github.com/txchat/dtalk/pkg/error"
)

type Middleware struct {
	log zerolog.Logger
}

func NewMiddleware(log zerolog.Logger) *Middleware {
	return &Middleware{
		log: log,
	}
}

func (m *Middleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := dumpRequest(ctx.Request)

		start := time.Now()

		ctx.Next()

		tlog := logger.NewLogWithCtx(ctx, m.log)

		latency := time.Since(start)
		if latency > time.Minute {
			latency = latency - latency%time.Second
		}
		reqURI := ctx.Request.RequestURI
		if reqURI == "" {
			reqURI = ctx.Request.URL.RequestURI()
		}

		tlog.Info().
			Str("clientIP", ctx.ClientIP()).
			Str("method", ctx.Request.Method).
			Str("Path", reqURI).
			Dur("span", latency).
			Str("|body", body).
			Int("status", ctx.Writer.Status()).
			Msg("http req")

		err, ok := ctx.Get(api.ReqError)
		if ok {
			code, msg := parseErr(err)
			if code != 0 {
				tlog.Error().
					Int("code", code).
					Str("msg", msg).
					Msg("http err")
			}
		}
	}
}

func ParseErr(err interface{}) (code int, msg string) {
	return parseErr(err)
}

func parseErr(err interface{}) (code int, msg string) {
	if err != nil {
		switch ty := err.(type) {
		case *xerror.Error:
			code = ty.Code()
			msg = ty.Error()
		case error:
			code = xerror.CodeInnerError
			msg = err.(error).Error()
		default:
			e := xerror.NewError(xerror.CodeInnerError)
			code = e.Code()
			msg = e.Error()
		}
		return
	}

	code = xerror.CodeOK

	return
}

func DupReadCloser(reader io.ReadCloser) (io.ReadCloser, io.ReadCloser) {
	var buf bytes.Buffer
	tee := io.TeeReader(reader, &buf)
	return ioutil.NopCloser(tee), ioutil.NopCloser(&buf)
}

// dumpRequest 格式化请求样式
func dumpRequest(req *http.Request) string {
	var dup io.ReadCloser
	var err error
	//req.Body, dup = iox.DupReadCloser(req.Body)
	req.Body, dup = DupReadCloser(req.Body)

	var b bytes.Buffer

	//reqURI := req.RequestURI
	//if reqURI == "" {
	//	reqURI = req.URL.RequestURI()
	//}

	//fmt.Fprintf(&b, "%s - %s - HTTP/%d.%d - OperaotrId:%s - ReqBody:", req.Method,
	//	reqURI, req.ProtoMajor, req.ProtoMinor, operatorId)

	chunked := len(req.TransferEncoding) > 0 && req.TransferEncoding[0] == "chunked"
	if req.Body != nil {
		var n int64
		var dest io.Writer = &b
		if chunked {
			dest = httputil.NewChunkedWriter(dest)
		}
		n, err = io.Copy(dest, req.Body)
		if chunked {
			dest.(io.Closer).Close()
		}
		if n > 0 {
			//io.WriteString(&b, "\n")
		}
	}

	req.Body = dup
	if err != nil {
		return err.Error()
	}

	return b.String()
}
