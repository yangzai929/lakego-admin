package datebin

import (
	"testing"
	"time"
)

func Test_String(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index  string
		layout string
		date   string
	}{
		{
			index:  "index-1",
			layout: "2006-01-02 15:04:05",
			date:   "2024-01-03 21:15:12",
		},
		{
			index:  "index-2",
			layout: "20060102150405",
			date:   "20240103211512",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		parseTt := NewDatebin().WithTime(tt).String()

		eq(parseTt, tt.Format("2006-01-02 15:04:05"), "failed String, index "+td.index)
	}
}

func Test_GoString(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index  string
		layout string
		date   string
	}{
		{
			index:  "index-1",
			layout: "2006-01-02 15:04:05",
			date:   "2024-01-03 21:15:12",
		},
		{
			index:  "index-2",
			layout: "20060102150405",
			date:   "20240103211512",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		parseTt := NewDatebin().WithTime(tt).GoString()

		eq(parseTt, tt.GoString(), "failed GoString, index "+td.index)
	}
}

func Test_ToStdtime(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index  string
		layout string
		date   string
	}{
		{
			index:  "index-1",
			layout: "2006-01-02 15:04:05",
			date:   "2024-01-03 21:15:12",
		},
		{
			index:  "index-2",
			layout: "20060102150405",
			date:   "20240103211512",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		parseTt := NewDatebin().WithTime(tt).ToStdTime()

		eq(parseTt, tt, "failed ToStdTime, index "+td.index)
	}
}

func Test_ToString(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index  string
		layout string
		date   string
	}{
		{
			index:  "index-1",
			layout: "2006-01-02 15:04:05",
			date:   "2024-01-03 21:15:12",
		},
		{
			index:  "index-2",
			layout: "20060102150405",
			date:   "20240103211512",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		parseTt := NewDatebin().WithTime(tt).ToString()

		eq(parseTt, tt.String(), "failed ToString, index "+td.index)
	}
}

func Test_ToStarString(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index  string
		layout string
		date   string
		star   string
	}{
		{
			index:  "index-1",
			layout: "2006-01-02 15:04:05",
			date:   "2024-01-03 21:15:12",
			star:   "Capricorn",
		},
		{
			index:  "index-2",
			layout: "2006-01-02 15:04:05",
			date:   "2024-01-21 21:15:12",
			star:   "Aquarius",
		},
		{
			index:  "index-3",
			layout: "2006-01-02 15:04:05",
			date:   "2024-02-21 21:15:12",
			star:   "Pisces",
		},
		{
			index:  "index-4",
			layout: "2006-01-02 15:04:05",
			date:   "2024-03-22 21:15:12",
			star:   "Aries",
		},
		{
			index:  "index-5",
			layout: "2006-01-02 15:04:05",
			date:   "2024-04-22 21:15:12",
			star:   "Taurus",
		},
		{
			index:  "index-6",
			layout: "2006-01-02 15:04:05",
			date:   "2024-05-22 21:15:12",
			star:   "Gemini",
		},
		{
			index:  "index-7",
			layout: "2006-01-02 15:04:05",
			date:   "2024-06-25 21:15:12",
			star:   "Cancer",
		},
		{
			index:  "index-8",
			layout: "2006-01-02 15:04:05",
			date:   "2024-07-25 21:15:12",
			star:   "Leo",
		},
		{
			index:  "index-9",
			layout: "2006-01-02 15:04:05",
			date:   "2024-08-25 21:15:12",
			star:   "Virgo",
		},
		{
			index:  "index-10",
			layout: "2006-01-02 15:04:05",
			date:   "2024-09-25 21:15:12",
			star:   "Libra",
		},
		{
			index:  "index-11",
			layout: "2006-01-02 15:04:05",
			date:   "2024-10-27 21:15:12",
			star:   "Scorpio",
		},
		{
			index:  "index-12",
			layout: "2006-01-02 15:04:05",
			date:   "2024-11-27 21:15:12",
			star:   "Sagittarius",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		star := NewDatebin().WithTime(tt).ToStarString()

		eq(star, td.star, "failed ToStarString, index "+td.index)
	}
}

func Test_ToSeasonString(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index  string
		layout string
		date   string
		season string
	}{
		{
			index:  "index-1",
			layout: "2006-01-02 15:04:05",
			date:   "2024-05-03 21:15:12",
			season: "Spring",
		},
		{
			index:  "index-2",
			layout: "2006-01-02 15:04:05",
			date:   "2024-07-21 21:15:12",
			season: "Summer",
		},
		{
			index:  "index-3",
			layout: "2006-01-02 15:04:05",
			date:   "2024-09-21 21:15:12",
			season: "Autumn",
		},
		{
			index:  "index-4",
			layout: "2006-01-02 15:04:05",
			date:   "2024-12-22 21:15:12",
			season: "Winter",
		},
		{
			index:  "index-5",
			layout: "2006-01-02 15:04:05",
			date:   "2024-11-22 21:15:12",
			season: "Autumn",
		},
		{
			index:  "index-6",
			layout: "2006-01-02 15:04:05",
			date:   "2024-08-22 21:15:12",
			season: "Summer",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		season := NewDatebin().WithTime(tt).ToSeasonString()

		eq(season, td.season, "failed ToSeasonString, index "+td.index)
	}
}

func Test_ToWeekdayString(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index   string
		layout  string
		date    string
		weekday string
	}{
		{
			index:   "index-1",
			layout:  "2006-01-02 15:04:05",
			date:    "2024-05-03 21:15:12",
			weekday: "Friday",
		},
		{
			index:   "index-2",
			layout:  "2006-01-02 15:04:05",
			date:    "2024-07-21 21:15:12",
			weekday: "Sunday",
		},
		{
			index:   "index-3",
			layout:  "2006-01-02 15:04:05",
			date:    "2024-09-21 21:15:12",
			weekday: "Saturday",
		},
		{
			index:   "index-4",
			layout:  "2006-01-02 15:04:05",
			date:    "2023-12-26 21:15:12",
			weekday: "Tuesday",
		},
		{
			index:   "index-5",
			layout:  "2006-01-02 15:04:05",
			date:    "2023-11-13 21:15:12",
			weekday: "Monday",
		},
		{
			index:   "index-6",
			layout:  "2006-01-02 15:04:05",
			date:    "2024-08-22 21:15:12",
			weekday: "Thursday",
		},
		{
			index:   "index-7",
			layout:  "2006-01-02 15:04:05",
			date:    "2023-08-16 21:15:12",
			weekday: "Wednesday",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		weekday := NewDatebin().WithTime(tt).ToWeekdayString()

		eq(weekday, td.weekday, "failed ToWeekdayString, index "+td.index)
	}
}

func Test_ToLayoutString(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index    string
		layout   string
		date     string
		toLayout string
		toDate   string
	}{
		{
			index:    "index-1",
			layout:   "2006-01-02 15:04:05",
			date:     "2024-05-03 21:15:12",
			toLayout: "2006-01-02",
			toDate:   "2024-05-03",
		},
		{
			index:    "index-2",
			layout:   "2006-01-02 15:04:05",
			date:     "2024-07-21 21:15:12",
			toLayout: "2006-01-02 15",
			toDate:   "2024-07-21 21",
		},
		{
			index:    "index-3",
			layout:   "2006-01-02 15:04:05",
			date:     "2024-09-21 21:15:12",
			toLayout: "2006",
			toDate:   "2024",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		toDate := NewDatebin().WithTime(tt).ToLayoutString(td.toLayout)

		eq(toDate, td.toDate, "failed ToLayoutString, index "+td.index)
	}
}

func Test_ToFormatString(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index    string
		layout   string
		date     string
		toLayout string
		toDate   string
	}{
		{
			index:    "index-1",
			layout:   "2006-01-02 15:04:05",
			date:     "2024-05-03 21:15:12",
			toLayout: "Y-m-d",
			toDate:   "2024-05-03",
		},
		{
			index:    "index-2",
			layout:   "2006-01-02 15:04:05",
			date:     "2024-07-21 21:15:12",
			toLayout: "Y-m-d H",
			toDate:   "2024-07-21 21",
		},
		{
			index:    "index-3",
			layout:   "2006-01-02 15:04:05",
			date:     "2024-09-21 21:15:12",
			toLayout: "Y",
			toDate:   "2024",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		toDate := NewDatebin().WithTime(tt).ToFormatString(td.toLayout)

		eq(toDate, td.toDate, "failed ToFormatString, index "+td.index)
	}
}

func Test_ToOtherFormatString(t *testing.T) {
	eq := assertEqualT(t)

	tests := []struct {
		index    string
		layout   string
		date     string
		toLayout func(Datebin) string
		toDate   string
	}{
		{
			index:  "index-1",
			layout: "2006-01-02 15:04:05",
			date:   "2023-05-03 21:15:12",
			toLayout: func(d Datebin) string {
				return d.ToAnsicString()
			},
			toDate: "Wed May  3 21:15:12 2023",
		},
		{
			index:  "index-2",
			layout: "2006-01-02 15:04:05",
			date:   "2023-05-03 21:15:12",
			toLayout: func(d Datebin) string {
				return d.ToUnixDateString()
			},
			toDate: "Wed May  3 21:15:12 CST 2023",
		},
		{
			index:  "index-3",
			layout: "2006-01-02 15:04:05",
			date:   "2023-05-03 21:15:12",
			toLayout: func(d Datebin) string {
				return d.ToRubyDateString()
			},
			toDate: "Wed May 03 21:15:12 +0800 2023",
		},
		{
			index:  "index-4",
			layout: "2006-01-02 15:04:05",
			date:   "2023-05-03 21:15:12",
			toLayout: func(d Datebin) string {
				return d.ToRFC850String()
			},
			toDate: "Wednesday, 03-May-23 21:15:12 CST",
		},
		{
			index:  "index-5",
			layout: "2006-01-02 15:04:05",
			date:   "2023-05-03 21:15:12",
			toLayout: func(d Datebin) string {
				return d.ToRssString()
			},
			toDate: "Wed, 03 May 2023 21:15:12 +0800",
		},
		{
			index:  "index-6",
			layout: "2006-01-02 15:04:05",
			date:   "2023-05-03 21:15:12",
			toLayout: func(d Datebin) string {
				return d.ToISO8601String()
			},
			toDate: "2023-05-03T21:15:12+08:00",
		},
		{
			index:  "index-7",
			layout: "2006-01-02 15:04:05",
			date:   "2023-05-03 21:15:12",
			toLayout: func(d Datebin) string {
				return d.ToW3CString()
			},
			toDate: "2023-05-03T21:15:12+08:00",
		},
		{
			index:  "index-8",
			layout: "2006-01-02 15:04:05",
			date:   "2023-05-03 21:15:12",
			toLayout: func(d Datebin) string {
				return d.ToDatetimeString()
			},
			toDate: "2023-05-03 21:15:12",
		},
	}

	for _, td := range tests {
		tt, err := time.ParseInLocation(td.layout, td.date, time.Local)
		if err != nil {
			t.Fatal(err)
		}

		toDate := td.toLayout(NewDatebin().WithTime(tt))

		eq(toDate, td.toDate, "failed ToOtherFormatString, index "+td.index)
	}
}