package db

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

type testpair struct {
	value    string
	expected int
}

var tests = []testpair{
	{"trace", 6},
	{"debug", 5},
	{"info", 4},
	{"warning", 3},
	{"error", 2},
}

func TestPgxLogLevel(t *testing.T) {
	for _, pair := range tests {
		ll, _ := log.ParseLevel(pair.value)
		v := pgxLogLevel(ll)
		if int(v) != pair.expected {
			t.Error(
				"For", pair.value,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
