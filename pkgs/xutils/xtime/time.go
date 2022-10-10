package xtime

import (
	// "crypto/rand"
	"math/rand"
	"time"
)

const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
)

func GetRandomTime(from, to time.Time) time.Time {
	t1 := from.UnixMilli()
	t2 := to.UnixMilli()
	if t1 > t2 {
		t1, t2 = t2, t1
	}
	delta := t2 - t1
	randomTimestampMs := rand.Int63n(delta) + t1
	return time.UnixMilli(randomTimestampMs)
}

/**
 * get random time from 'from' to 'to'. It will parse the string type time to time.Time
 * use the layout "2006-01-02 15:04:05". return empty time.Time{} and error msg if parse
 * error.
 */
func GetRandomTimeString(from, to string) (time.Time, error) {
	t1, err := time.Parse(TIME_LAYOUT, from)
	if err != nil {
		return time.Time{}, err
	}
	t2, err := time.Parse(TIME_LAYOUT, to)
	if err != nil {
		return time.Time{}, err
	}
	return GetRandomTime(t1, t2), nil
}
