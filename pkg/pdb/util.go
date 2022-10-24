package pdb

import "time"

func TimeToMs(t time.Time) int64 {
	return t.UnixMicro()
}

func MsToTime(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Microsecond))
}
