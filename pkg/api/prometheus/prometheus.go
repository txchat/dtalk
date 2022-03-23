package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"reflect"
)

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//path := c.FullPath()
		//timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		//c.Next()
		//
		//// 统计错误 code
		//err := c.MustGet(api.ReqError)
		//code, _, _ := parseErr(nil, err)
		////if code != 0 {
		////	responseCode.WithLabelValues(strconv.Itoa(code)).Inc()
		////}
		//
		//status := c.Writer.Status()
		////responseStatus.WithLabelValues(strconv.Itoa(status)).Inc()
		//totalRequests.WithLabelValues(path, strconv.Itoa(status), strconv.Itoa(code)).Inc()
		//timer.ObserveDuration()
	}
}

func init() {
	_ = prometheus.Register(totalRequests)
	//_ = prometheus.Register(responseStatus)
	_ = prometheus.Register(httpDuration)
	//_ = prometheus.Register(responseCode)
}

// 统计所有 url 访问次数
var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path", "status", "code"},
)

//
//var responseStatus = prometheus.NewCounterVec(
//	prometheus.CounterOpts{
//		Name: "http_response_status",
//		Help: "Status of HTTP response",
//	},
//	[]string{"status"},
//)
//
////
//var responseCode = prometheus.NewCounterVec(
//	prometheus.CounterOpts{
//		Name: "http_response_code",
//		Help: "Status of HTTP response",
//	},
//	[]string{"code"},
//)

//
var httpDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	},
	[]string{"path"},
)
