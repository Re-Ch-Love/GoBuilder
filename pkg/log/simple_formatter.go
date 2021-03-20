package log

import (
	"fmt"
	"gitee.com/KongchengPro/GoBuilder/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

/*
SimpleFormatter is a simple but beautiful log formatter
Usage:
	log.SetFormatter(new(SimpleFormatter))
	log.SetReportCaller(true)  // Must do this, or no caller info.
*/
type SimpleFormatter struct {
	EnableDebug bool
}

func (s *SimpleFormatter) Format(e *log.Entry) (bytes []byte, retErr error) {
	var b strings.Builder
	var msg string
	if e.Message != "" {
		msg = e.Message
	} else {
		msg = "nil"
	}
	level := strings.ToUpper(e.Level.String())
	defer utils.ReturnErrorFromPanic(&retErr, nil)
	b.WriteString("[")
	b.WriteString(level)
	b.WriteString("] ")
	if s.EnableDebug {
		// debug output format
		// [LEVEL] MST 2006-01-02 15:04:05 >> message >> { field1: xxx, field2: xxx } >> funcName:lineNumber
		t := time.Now().Format("MST 2006-01-02 15:04:05")
		b.WriteString(t)
		b.WriteString(" >> ")
		b.WriteString(msg)
		b.WriteString(" >> ")
		b.WriteString("{ ")
		fields2String(e.Data, &b)
		b.WriteString(" }")
		b.WriteString(" >> ")
		if e.HasCaller() {
			b.WriteString(e.Caller.Function + ":" + strconv.Itoa(e.Caller.Line))
		} else {
			b.WriteString("nil:nil")
		}
	} else {
		// no debug output format
		// [LEVEL] message { field1: xxx, field2: xxx }
		b.WriteString(msg)
		b.WriteString(" { ")
		fields2String(e.Data, &b)
		b.WriteString(" }")
	}
	b.WriteString("\n")
	return []byte(b.String()), nil
}

func fields2String(f log.Fields, b *strings.Builder) {
	l := len(f)
	if l != 0 {
		var n int
		for k, v := range f {
			n++
			b.WriteString(k)
			b.WriteString(": ")
			vStr := fmt.Sprintf("%v", v)
			if vStr == "<nil>" {
				vStr = strings.Trim(vStr, "<>")
			} else {
				b.WriteString(vStr)
			}
			if n != l {
				b.WriteString(", ")
			}
		}
	} else {
		b.WriteString("nil")
	}
}
