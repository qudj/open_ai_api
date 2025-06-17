package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qudj/open_ai_api/utils"
	"net/http"
	"net/http/httputil"
	"regexp"
	"time"
)

// Logger 自定义日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		re := regexp.MustCompile(`/health(_[^/]*)?$`)
		logger := utils.SugarLogger()
		if !re.MatchString(path) {
			query := c.Request.URL.RawQuery
			c.Next()
			latency := time.Since(start)
			if len(c.Errors) > 0 {
				logger.Errorf("errs: %+v", c.Errors.Errors())
			} else {
				msg := "[Access-Log]"
				stressTest := c.Query("stress_test")
				if stressTest != "" {
					msg = fmt.Sprintf("%s - StressTest (%s)", msg, stressTest)
				}
				logger.Infow(msg,
					"method", c.Request.Method,
					"referer", c.Request.Referer(),
					"path", path,
					"router_path", fmt.Sprintf("%s:%s", c.Request.Method, c.FullPath()),
					"status", c.Writer.Status(),
					"response_size", c.Writer.Size(),
					"query", query,
					"ip", c.ClientIP(),
					"ua", c.Request.UserAgent(),
					"latency", latency.Milliseconds(),
				)
			}
		}
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := utils.SugarLogger()
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				logger.Errorw("[Recovery from panic]",
					"err", err,
					"req", string(httpRequest),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
