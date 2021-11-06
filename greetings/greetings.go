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

func init() {
	rand.Seed(time.Now().Unix())
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
