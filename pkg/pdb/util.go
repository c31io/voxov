package pdb

import "time"

func TimeToUs(t time.Time) int64 {
	return t.UnixMicro()
}

func UsToTime(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Microsecond))
}
