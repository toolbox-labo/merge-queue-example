package main

import "testing"

func TestGetMSG(t *testing.T) {
	want := "Hi"
	got := getMSG()

	if want != got {
		t.Errorf("Unexpected result. want: %s, got: %s", want, got)
	}
}

func TestGreeting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "John",
			want: "Hi, John",
		},
		{
			name: "Tarou",
			want: "Hi, Tarou",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := greeting(tt.name)
			if tt.want != got {
				t.Errorf("Unexpected result. want: %s, got: %s", tt.want, got)
			}
		})
	}
}
