package calendar

import "time"

// Event is a event with a begin/end time,
// a title and other things like description & location.
type Event struct {
	Begin, End time.Time
	Title      string

	Others map[string]string
}

// Todo is another type of event with only a date
// and no other times.
type Todo struct {
	Date  time.Time
	Title string

	Others map[string]string
}
