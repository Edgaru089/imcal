package sync

import (
	"fmt"
	"time"

	"edgaru089.ink/go/imcal/internal/calendar"
	"edgaru089.ink/go/imcal/internal/util"
	"github.com/emersion/go-ical"
	"github.com/emersion/go-webdav"
	"github.com/emersion/go-webdav/caldav"
)

type Client struct {
	c             *caldav.Client
	authed_client webdav.HTTPClient

	principal, home_set string
	Calendars           []caldav.Calendar
}

func NewClient(username, password, url string) (c *Client, err error) {
	c = &Client{}

	c.authed_client = webdav.HTTPClientWithBasicAuth(nil, username, password)
	c.c, err = caldav.NewClient(c.authed_client, url)
	if err != nil {
		return nil, err
	}

	c.principal, err = c.c.FindCurrentUserPrincipal()
	if err != nil {
		return nil, err
	}
	fmt.Println("FindCurrentUserPrincipal =", c.principal)

	c.home_set, err = c.c.FindCalendarHomeSet(c.principal)
	if err != nil {
		return nil, err
	}
	fmt.Println("FindCalendarHomeSet =", c.home_set)

	fmt.Println("FindCalendars")
	c.Calendars, err = c.c.FindCalendars(c.home_set)
	if err != nil {
		return nil, err
	}

	for _, t := range c.Calendars {
		fmt.Printf("NAME=%s PATH=%s DESC=%s FEATs=%v\n", t.Name, t.Path, t.Description, t.SupportedComponentSet)
	}

	return
}

func (c *Client) PullCalendar(path string) (cal calendar.Calendar, err error) {

	query := &caldav.CalendarQuery{}
	obj, err := c.c.QueryCalendar(path, query)
	if err != nil {
		return
	}

	for _, c0 := range obj {
		for _, e := range c0.Data.Events() {

			fmt.Printf("NAME=%s TITLE=%s@%s START=%v END=%v\n", e.Name, util.Whatever(e.Props.Text(ical.PropSummary)), util.Whatever(e.Props.Text(ical.PropLocation)), util.Unwrap(e.DateTimeStart(time.Local)), util.Unwrap(e.DateTimeEnd(time.Local)))

			cal.Events = append(cal.Events, calendar.Event{
				Begin: util.Whatever(e.DateTimeStart(time.Local)),
				End:   util.Whatever(e.DateTimeStart(time.Local)),
				Title: util.Whatever(e.Props.Text(ical.PropSummary)),
			})
		}
	}

	return
}
