package utils

import "time"

const formatYYYYMMDDHHmmss = "2006-01-02 15:04:05"
const formatYYYYMMDD = "2006-01-02"

func TimeToStringDateTime(t time.Time) string {
	return t.Format(formatYYYYMMDDHHmmss)
}

func TimeToStringDate(t time.Time) string {
	return t.Format(formatYYYYMMDD)
}
