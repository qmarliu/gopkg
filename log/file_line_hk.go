package log

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type fileHook struct{}

func newFileHook() *fileHook {
	return &fileHook{}
}

func (f *fileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (f *fileHook) Fire(entry *logrus.Entry) error {
	var s string
	_, file, line, _ := runtime.Caller(8)
	i := strings.SplitAfter(file, "/")
	if len(i) > 3 {
		s = i[len(i)-3] + i[len(i)-2] + i[len(i)-1] + ":" + strconv.FormatInt(int64(line), 10)
	}
	entry.Data["FilePath"] = s
	return nil
}
