package model

import "time"

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

type HRKey struct {
	HotelId string
	RoomId  string
}

type DateQuota map[time.Time]int
