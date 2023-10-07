package helper

import "time"

func GetTime() time.Time {
	return time.Now().UTC()
}

func GetUnixTime() int64 {
	return GetTime().Unix()
}

func GetUnixTimeNano() int64 {
	return GetTime().UnixNano()
}

func GetUnixTimeMili() int64 {
	return GetTime().UnixMilli()
}

func GetUnixTimeMicro() int64 {
	return GetTime().UnixMicro()
}
