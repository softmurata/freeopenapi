package models

type FlowStatus string

// List of FlowStatus
const (
	FlowStatus_ENABLED_UPLINK   FlowStatus = "ENABLED-UPLINK"
	FlowStatus_ENABLED_DOWNLINK FlowStatus = "ENABLED-DOWNLINK"
	FlowStatus_ENABLED          FlowStatus = "ENABLED"
	FlowStatus_DISABLED         FlowStatus = "DISABLED"
	FlowStatus_REMOVED          FlowStatus = "REMOVED"
)