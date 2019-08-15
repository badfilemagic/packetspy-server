package utils

import (
	"github.com/google/gopacket"
	"log"
	"strconv"
	"time"
)


func converttime(t string) time.Time {
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return time.Unix(i,0).UTC()
}

func MakeCaptureInfo(fields []string) gopacket.CaptureInfo {

	ts := converttime(fields[0])
	cl, _ := strconv.Atoi(fields[1])
	len,_ := strconv.Atoi(fields[2])
	idx, _ := strconv.Atoi(fields[3])

	i := gopacket.CaptureInfo{ts, cl, len, idx, nil}

	return i
}
