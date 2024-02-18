package commands

import "testing"

func TestShutdownParameters(t *testing.T) {
	seconds, message, err := shutdownParameters("10 10 seconds later, server will be shutdown")
	if err != "" {
		t.Errorf("Error: %s", err)
	}

	if seconds != 10 {
		t.Errorf("Expected 10, got %d", seconds)
	}

	if message != "10_seconds_later,_server_will_be_shutdown" {
		t.Errorf("Expected 10_seconds_later,_server_will_be_shutdown, got %s", message)
	}
}
