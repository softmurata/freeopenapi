package models

type N1NotificationRequestData struct {
	N1MessageContainer *N1MessageContainer  `json:"n1MessageContainer,omitempty"`
	N1NotifySubscriptionId string  `json:"n1NotifySubscriptionId,omitempty"`
}

