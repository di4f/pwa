package app

import (
	"testing"

	"github.com/omnipunk/pwa/v9/errors"
	"github.com/omnipunk/pwa/v9/logs"
)

func TestLog(t *testing.T) {
	DefaultLogger = t.Logf
	Log("hello", "world")
	Logf("hello %v", "Maxoo")
}

func TestServerLog(t *testing.T) {
	testSkipWasm(t)
	testLogger(t, serverLog)
}

func TestClientLog(t *testing.T) {
	testSkipNonWasm(t)
	testLogger(t, clientLog)
}

func testLogger(t *testing.T, l func(string, ...any)) {
	utests := []struct {
		scenario string
		value    any
	}{
		{
			scenario: "log",
			value:    logs.New("test").WithTag("type", "log"),
		},
		{
			scenario: "error",
			value:    errors.New("test").WithTag("type", "error"),
		},
	}

	for _, u := range utests {
		t.Run(u.scenario, func(t *testing.T) {
			l("%v", u.value)
		})
	}
}
