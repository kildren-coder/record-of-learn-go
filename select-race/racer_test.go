package select_race

import (
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.fackbook.com"
	fastURL := "http://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
