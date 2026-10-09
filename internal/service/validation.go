package service

import "fmt"

func ValidateChannel(channel string) error {

	if channel == "" {
		return fmt.Errorf("channel cannot be empty")
	}

	if channel != "EMAIL" && channel != "SMS" {
		return fmt.Errorf("invalid channel: %s", channel)
	}

	return nil
}
