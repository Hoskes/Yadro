package time_parser

import "time"

// CheckTimeDeviation Отдает true если разница между endTime и startTime не превышает допустимой delta.
func CheckTimeDeviation(startTime time.Time, endTime time.Time, delta string) (bool, error) {
	duration, err := ParseStrToDuration(delta)
	if err != nil {
		return false, err
	}
	return endTime.Sub(startTime) > duration, nil
}
