package model

import "time"

type Order struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Email   string    `json:"email"`
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
}

func (r *Order) GetRangeFromOrder() []time.Time {
	if r.From.After(r.To) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := normalizeDate(r.From); !d.After(normalizeDate(r.To)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}
	return days
}

func normalizeDate(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
