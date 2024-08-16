package main

import (
	"testing"
	"time"

	"poiesis.interimme.net/internal/assert"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 7, 13, 13, 33, 33, 33, time.UTC),
			want: "13 Jul 2024 at 13:33",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 7, 13, 13, 33, 33, 33, time.FixedZone("CET", 1*60*60)),
			want: "13 Jul 2024 at 12:33",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
