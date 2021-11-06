package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retireived message with the name
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of random greeting formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"hail, %v! Well met!",
	}

	// Return a randomly selected greeting format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
