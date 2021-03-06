package holidays

import (
	"testing"
	"time"
)

// fakeDate create date for testing.
func fakeDate(y, m, d int) time.Time {
	if y == 0 && m == 0 && d == 0 {
		return midnight(time.Time{})
	}
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

// formatDate returns shorter date representation.
func formatDate(date time.Time) string {
	return date.Format(ISO8601DateFormat)
}

func TestNew(t *testing.T) {
	got, _ := New("hr", "testdata")
	want := 29
	if len(got) != want {
		t.Errorf("len(New(hr, testdata)) = %v; want %v",
			len(got),
			want,
		)
	}
}

func TestMidnight(t *testing.T) {
	in := time.Date(2016, 1, 23, 1, 2, 3, 4, time.UTC)
	want := time.Date(2016, 1, 23, 0, 0, 0, 0, time.UTC)
	if got := midnight(in); got != want {
		t.Errorf("midnight(%v) = %v; want %v", in, got, want)
	}
}

func TestParseDate(t *testing.T) {
	// invalid
	in := "abc"
	want := fakeDate(0, 0, 0)
	if got, err := parseDate(in); got != want || err == nil {
		t.Errorf("parseDate(%q) = %v, %v; want %v, error",
			in, formatDate(got), err, formatDate(want))
	}

	// valid
	in = "2016-01-23"
	want = fakeDate(2016, 1, 23)
	if got, err := parseDate(in); got != want || err != nil {
		t.Errorf("parseDate(%q) = %v, %v; want %v, nil",
			in, formatDate(got), err, formatDate(want))
	}
}

func TestDaysBetween(t *testing.T) {
	since := fakeDate(2012, 2, 3)
	var tests = []struct {
		now  time.Time
		want int
	}{
		{fakeDate(2014, 3, 21), 777},
		{fakeDate(2014, 7, 10), 888},
		{fakeDate(2014, 10, 29), 999},
	}
	for _, tt := range tests {
		if got := DaysBetween(since, tt.now); got != tt.want {
			t.Errorf("DaysBetween(%v, %v) = %v; want %v",
				formatDate(since),
				formatDate(tt.now), got, tt.want)
		}
	}

}

func TestNextFestivus(t *testing.T) {
	today := fakeDate(2016, 12, 20)
	want := fakeDate(2016, 12, 23)
	if got := NextFestivus(today); got != want {
		t.Errorf("NextFestivus(today: %v) = %v; want %v",
			formatDate(today),
			formatDate(got),
			formatDate(want),
		)
	}

	// today if after festivus in current year
	today = fakeDate(2016, 12, 29)
	want = fakeDate(2017, 12, 23)
	if got := NextFestivus(today); got != want {
		t.Errorf("NextFestivus(today: %v) = %v; want %v",
			formatDate(today),
			formatDate(got),
			formatDate(want),
		)
	}
}

func TestDaysToFestivus(t *testing.T) {
	today := fakeDate(2016, 12, 20)
	want := 3
	if got := DaysToFestivus(today); got != want {
		t.Errorf("DaysToFestivus(today: %v) = %v; want %v",
			formatDate(today),
			got,
			want,
		)
	}
}

func TestByYear(t *testing.T) {
	all, _ := New("hr", "testdata")
	today := fakeDate(2016, 12, 20)
	want := 14
	if got := ByYear(all, today); len(got) != want {
		t.Errorf("len(ByYear(all, today: %v)) = %v; want %v",
			formatDate(today),
			len(got),
			want,
		)
	}
}
