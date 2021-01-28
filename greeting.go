package greetings

import "fmt"

// Hello returns a greeting message.
func Hello(msg string) string {
	message := fmt.Sprintf("Greeting, %v", msg)
	return message
}
