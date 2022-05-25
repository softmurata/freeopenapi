package models

type UpPathChgEvent struct {
	NotificationUri string `json:"notificationUri" yaml:"notificationUri" bson:"notificationUri" mapstructure:"NotificationUri"`
	// It is used to set the value of Notification Correlation ID in the notification sent by the SMF.
	NotifCorreId string         `json:"notifCorreId" yaml:"notifCorreId" bson:"notifCorreId" mapstructure:"NotifCorreId"`
	DnaiChgType  DnaiChangeType `json:"dnaiChgType" yaml:"dnaiChgType" bson:"dnaiChgType" mapstructure:"DnaiChgType"`
}
