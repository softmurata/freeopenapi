package models

type QosResourceType string

// List of QosResourceType
const (
	QosResourceType_NON_GBR          QosResourceType = "NON_GBR"
	QosResourceType_NON_CRITICAL_GBR QosResourceType = "NON_CRITICAL_GBR"
	QosResourceType_CRITICAL_GBR     QosResourceType = "CRITICAL_GBR"
)
