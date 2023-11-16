package calendar

// Calendar is a collection of (either) events (or) todos.
type Calendar struct {
	Title       string
	Description string
	Events      []Event
	Totos       []Todo
}
