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
	afternoon := time.Date(2023, time.July, 17, 18, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		want string
	}{
		{
			name: "John",
			want: "Good afternoon, John",
		},
		{
			name: "Tarou",
			want: "Good afternoon, Tarou",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := greeting(tt.name, &afternoon)
			if tt.want != got {
				t.Errorf("Unexpected result. want: %s, got: %s", tt.want, got)
			}
		})
	}
}
