package utils

import "time"

const DateTimeFormat = time.RFC3339 // ISO 8601 format

// Tarihi string olarak döner
func FormatDateTime(t time.Time) string {
	return t.UTC().Format(DateTimeFormat)
}

// Tarihi time.Time'e çevirir
func ParseDateTime(dateTimeStr string) (time.Time, error) {
	return time.Parse(DateTimeFormat, dateTimeStr)
}
