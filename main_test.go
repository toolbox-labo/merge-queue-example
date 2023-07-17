package main

import (
	"testing"
	"time"
)

func TestGetMSG(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want string
	}{
		{
			name: "morning",
			time: time.Date(2023, time.July, 17, 12, 0, 0, 0, time.UTC),
			want: "Good morning",
		},
		{
			name: "afternoon",
			time: time.Date(2023, time.July, 17, 13, 0, 0, 1, time.UTC),
			want: "Good afternoon",
		},
		{
			name: "evening",
			time: time.Date(2023, time.July, 17, 18, 0, 0, 1, time.UTC),
			want: "Good evening",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getMSG(&tt.time)

			if tt.want != got {
				t.Errorf("Unexpected result. want: %s, got: %s", tt.want, got)
			}
		})
	}

}

func TestGreeting(t *testing.T) {
	evening := time.Date(2023, time.July, 17, 18, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		want string
	}{
		{
			name: "John",
			want: "Good evening, John",
		},
		{
			name: "Tarou",
			want: "Good evening, Tarou",
		},
		{
			name: "TooLongName!", // over 11 chars => Longname
			want: "Good evening, Longname",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := greeting(tt.name, &evening)
			if tt.want != got {
				t.Errorf("Unexpected result. want: %s, got: %s", tt.want, got)
			}
		})
	}
}
