package datetime

import (
	"fmt"
	"time"
	"strconv"
)

type PrettyConfig struct {
	Use12HourTime bool
	RemoveEmptySeconds bool
	HideSeconds bool
}

func ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return suffix
}

type TimeContainer struct {
    Hour string
    Minutes string
    Seconds string
    Ending string
}

func timeContainerToString(tc TimeContainer) string {
	var s string
	s = tc.Hour + ":" + tc.Minutes
	if tc.Seconds != "" {
		s = s + ":" + tc.Seconds
	}
	s = s + tc.Ending
	return s
}

func fmt2Int(i int) string {
    return fmt.Sprintf("%02d", i)
}

func timeFromConfig(t time.Time, c PrettyConfig) TimeContainer {
    var hour int
    var seconds string
	var end string
	if t.Hour() >= 12 {
		if t.Hour() >= 13 && c.Use12HourTime {
			hour = t.Hour() - 12
		} else {
			hour = t.Hour()
		}
		end = "pm"
	} else {
		hour = t.Hour()
		end = "am"
	}

	if !c.Use12HourTime {
		end = ""
	}

	if (t.Second() != 0 || !c.RemoveEmptySeconds) && !c.HideSeconds {
		seconds = fmt2Int(t.Second())
	}

	return TimeContainer{fmt2Int(hour), fmt2Int(t.Minute()), seconds, end}
}

func Pretty(t time.Time, c PrettyConfig) string {
	return fmt.Sprintf("%s the %s of %s in %d at %s",
		t.Weekday().String(),
		strconv.Itoa(t.Day()) + ordinal(t.Day()),
		t.Month().String(),
		t.Year(),
		timeContainerToString(timeFromConfig(t, c)),
	)
}

type TimeDateContainer struct {
	TimeContainer
	Weekday string
	Day string
	DayOrdinal string
	Month string
	Hour string
	Year string
}

func PrettyStruct(t time.Time, c PrettyConfig) TimeDateContainer {
	tdc := TimeDateContainer{}
	tdc.Weekday = t.Weekday().String()
	tdc.DayOrdinal = ordinal(t.Day())
	tdc.Day = strconv.Itoa(t.Day())
	tdc.Month = t.Month().String()
	tdc.Year = strconv.Itoa(t.Year())
	tc := timeFromConfig(t, c)
	tdc.Hour = tc.Hour
	tdc.Minutes = tc.Minutes
	tdc.Seconds = tc.Seconds
	tdc.Ending = tc.Ending
	return tdc
}
