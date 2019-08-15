package utils

import (
	"strings"
	"time"
)

func PrintableDate(t time.Time) string {
	stamp := t.String()
	fields := strings.Split(stamp, " ")
	ret := strings.Join(fields[:2], "__")
	ret = strings.Replace(ret, ":", "_", -1)
	return ret
}
