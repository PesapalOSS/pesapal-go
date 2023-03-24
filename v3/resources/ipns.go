package resources

import "time"

type IPNsResources service

type IPN struct {
	Id        string    `json:"ipn_id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_date"`
}

func (*IPNsResources) List() []IPN {
	ipns := []IPN{}

	return ipns
}

type RegisterIPN struct {
	Url              string `json:"url"`
	NotificationType string `json:"ipn_notification_type"`
}

func (*IPNsResources) Register(r *RegisterIPN) (*IPN, error) {
	ipn := &IPN{}

	return ipn, nil
}
