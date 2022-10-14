package log

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger(skipPath []string, isTestLog bool) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(skipPath); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range skipPath {
			skip[path] = struct{}{}
		}
	}

	if isTestLog {
		return func(c *gin.Context) {
			// Start timer
			start := time.Now()
			path := c.Request.URL.Path
			raw := c.Request.URL.RawQuery

			data, _ := c.GetRawData()
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

			// Process request
			c.Next()

			// Log only when path is not being skipped
			if _, ok := skip[path]; !ok {

				timeStamp := time.Now()
				if raw != "" {
					path = path + "?" + raw
				}

				Infof("%v | %3d | %13v | %15s | %-7s %#v | %s\n%s",
					timeStamp.Format("2006/01/02 - 15:04:05"),
					c.Writer.Status(),
					timeStamp.Sub(start),
					c.ClientIP(),
					c.Request.Method,
					path,
					string(data),
					c.Errors.ByType(gin.ErrorTypePrivate).String(),
				)
			}
		}
	}
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			timeStamp := time.Now()
			if raw != "" {
				path = path + "?" + raw
			}

			Infof("%v | %3d | %13v | %15s | %-7s %#v\n%s",
				timeStamp.Format("2006/01/02 - 15:04:05"),
				c.Writer.Status(),
				timeStamp.Sub(start),
				c.ClientIP(),
				c.Request.Method,
				path,
				c.Errors.ByType(gin.ErrorTypePrivate).String(),
			)
		}
	}
}
