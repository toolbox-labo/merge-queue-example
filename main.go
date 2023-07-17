package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%s", greeting("John", &now))
}

func getMSG(d *time.Time) string {
	if d.Hour() <= 12 {
		return "Good morning"
	} else if d.Hour() >= 18 {
		return "Good evening"
	}
	return "Good afternoon"
}

func greeting(name string, d *time.Time) string {
	return fmt.Sprintf("%s, %s", getMSG(d), name)
}
