package model

import (
	"testing"
	"time"
)

func TestOrder_Validate(t *testing.T) {
	type fields struct {
		HotelID string
		RoomID  string
		Email   string
		From    time.Time
		To      time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Empty body", fields{}, false},
		{"Correct body", fields{"reddison", "123", "123@test.com", time.Now(), time.Now().Add(24 * time.Hour)}, true},
		{"Incorrect time range", fields{"reddison", "123", "123@test.com", time.Now().Add(24 * time.Hour), time.Now()}, false},
		{"Incorrect time range", fields{"reddison", "123", "123@test.com", time.Now(), time.Time{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Order{
				HotelID: tt.fields.HotelID,
				RoomID:  tt.fields.RoomID,
				Email:   tt.fields.Email,
				From:    tt.fields.From,
				To:      tt.fields.To,
			}
			if got := r.Validate(); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
