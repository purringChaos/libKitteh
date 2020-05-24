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
	return strconv.Itoa(x) + suffix
}

func timeStr(t time.Time, c PrettyConfig) string {
	var end string
	var hour int
	if t.Hour() >= 12 {
		if t.Hour() >= 13 && !c.Use12HourTime {
			hour = t.Hour() - 12
		} else {
			hour = t.Hour()
		}
		end = "pm"
	} else {
		hour = t.Hour()
		end = "am"
	}

	if c.Use12HourTime {
		end = ""
	}

	str := fmt.Sprintf("%02d:%02d",
		hour,
		t.Minute())

	if (t.Second() != 0 || !c.RemoveEmptySeconds) && !c.HideSeconds {
		str = str + fmt.Sprintf(":%02d", t.Second())
	}
	str = str + end

	return str
}

func Pretty(t time.Time, c PrettyConfig) string {
	return fmt.Sprintf("%s the %s of %s in %d at %s",
		t.Weekday().String(),
		ordinal(t.Day()),
		t.Month().String(),
		t.Year(),
		timeStr(t, c),
	)
}