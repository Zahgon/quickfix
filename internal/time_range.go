package internal

import (
	"time"
)

// TimeOfDay represents the time of day.
type TimeOfDay struct {
	hour, minute, second int
	d                    time.Duration
}

const shortForm = "15:04:05"

// NewTimeOfDay returns a newly initialized TimeOfDay.
func NewTimeOfDay(hour, minute, second int) TimeOfDay {
	_ = "STUB: not implemented"
	return *new(TimeOfDay)
}

// ParseTimeOfDay parses a TimeOfDay from a string in the format HH:MM:SS.
func ParseTimeOfDay(str string) (TimeOfDay, error) {
	_ = "STUB: not implemented"
	return *new(TimeOfDay), nil
}

// TimeRange represents a time band in a given time zone.
type TimeRange struct {
	startTime, endTime TimeOfDay
	weekdays           []time.Weekday
	startDay, endDay   *time.Weekday
	loc                *time.Location
}

// NewUTCTimeRange returns a time range in UTC.
func NewUTCTimeRange(start, end TimeOfDay, weekdays []time.Weekday) (*TimeRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewTimeRangeInLocation returns a time range in a given location.
func NewTimeRangeInLocation(start, end TimeOfDay, weekdays []time.Weekday, loc *time.Location) (*TimeRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewUTCWeekRange returns a weekly TimeRange.
func NewUTCWeekRange(startTime, endTime TimeOfDay, startDay, endDay time.Weekday) (*TimeRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWeekRangeInLocation returns a time range in a given location.
func NewWeekRangeInLocation(startTime, endTime TimeOfDay, startDay, endDay time.Weekday, loc *time.Location) (*TimeRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TimeRange) isInWeekdays(day time.Weekday) bool { _ = "STUB: not implemented"; return false }

func (r *TimeRange) addWeekdayOffset(day time.Weekday, offset int) time.Weekday {
	_ = "STUB: not implemented"
	return *new(time.Weekday)
}

func (r *TimeRange) isInTimeRange(t time.Time) bool { _ = "STUB: not implemented"; return false }

func (r *TimeRange) isInWeekRange(t time.Time) bool { _ = "STUB: not implemented"; return false }

// IsInRange returns true if time t is within in the time range.
func (r *TimeRange) IsInRange(t time.Time) bool { _ = "STUB: not implemented"; return false }

// IsInSameRange determines if two points in time are in the same time range.
func (r *TimeRange) IsInSameRange(t1, t2 time.Time) bool { _ = "STUB: not implemented"; return false }
