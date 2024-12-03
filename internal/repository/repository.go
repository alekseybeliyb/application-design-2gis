package repository

import (
	"app-booking/internal/handler"
	"app-booking/internal/model"
	"errors"
	"maps"
	"sync"
	"time"
)

var (
	ErrNoAvailability = errors.New("no availability")
)

type repository struct {
	mutex        sync.Mutex
	orders       []model.Order
	availability map[model.HRKey]map[time.Time]int
}

func NewInMemoryRepository() handler.OrderRepository {
	return &repository{
		orders:       make([]model.Order, 0),
		availability: make(map[model.HRKey]map[time.Time]int),
	}
}

func (r *repository) Migrate() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	initAvailability := []model.RoomAvailability{
		{"reddison", "lux", date(2024, 1, 1), 1},
		{"reddison", "lux", date(2024, 1, 2), 1},
		{"reddison", "lux", date(2024, 1, 3), 1},
		{"reddison", "lux", date(2024, 1, 4), 1},
		{"reddison", "lux", date(2024, 1, 5), 0},
	}

	availability := make(map[model.HRKey]map[time.Time]int)

	for _, room := range initAvailability {
		key := model.HRKey{HotelId: room.HotelID, RoomId: room.RoomID}
		_, ok := availability[key]
		if !ok {
			availability[key] = make(map[time.Time]int)
		}
		availability[key][room.Date] += room.Quota
	}

	r.availability = availability

	return nil
}

func (r *repository) CreateOrder(newOrder *model.Order) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err := r.decrementQuota(newOrder)
	if err != nil {
		return err
	}
	r.orders = append(r.orders, *newOrder)

	return nil
}

func (r *repository) decrementQuota(newOrder *model.Order) error {
	copyAvailability := make(map[model.HRKey]map[time.Time]int, len(r.availability))
	maps.Copy(copyAvailability, r.availability)
	for currentDate := newOrder.From; currentDate.Before(newOrder.To); currentDate = currentDate.AddDate(0, 0, 1) {
		key := model.HRKey{HotelId: newOrder.HotelID, RoomId: newOrder.RoomID}
		if dateQuota, ok := copyAvailability[key]; ok {
			if day, dayExists := dateQuota[currentDate]; dayExists && day >= 1 {
				dateQuota[currentDate]--
			} else {
				return ErrNoAvailability
			}
		} else {
			return ErrNoAvailability
		}
	}
	r.availability = copyAvailability

	return nil
}

func (r *repository) GetAllOrders() ([]model.Order, error) {
	return r.orders, nil
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}