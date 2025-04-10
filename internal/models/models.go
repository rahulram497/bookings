package models

import "time"

// User is the user model DB
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AccessLevel int
}

// Room is the rooms DB model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction model for DB
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservation is the reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
}

// RoomRestriction is the room restriction model
type RoomRestriction struct {
	ID             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	StartDate      time.Time
	EndDate        time.Time
	RoomID         int
	ReservationID  int
	RestrictionsID int
	Room           Room
	Reservation    Reservation
	Restriction    Reservation
}
