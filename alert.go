package sdk

import "time"

// AlertNotification is a grafana alert notification channel.
type AlertNotification struct {
	ID           int    `json:"id"`
	UID          string `json:"uid"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	IsDefault    bool   `json:"isDefault"`
	SendReminder bool   `json:"sendReminder"`
	Settings     struct {
		Addresses string `json:"addresses"`
	} `json:"settings"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
