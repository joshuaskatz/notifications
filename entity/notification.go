package entity

import "time"

type Notification interface {
	isNotification()
}

type BaseNotification struct {
	CreatedAt time.Time `json: "createdAt"`
}

func (BaseNotification) isNotification() {}

type UnreadWorkRequest struct {
	BaseNotification
	WorkID int    `json:"workID"`
	Title  string `json:"title"`
}

type UnreadMessagesNotification struct {
	BaseNotification
	Count int `json:"count"`
}
