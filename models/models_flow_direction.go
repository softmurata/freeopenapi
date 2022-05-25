package models

type FlowDirection string

// List of FlowDirection
const (
	FlowDirection_DOWNLINK      FlowDirection = "DOWNLINK"
	FlowDirection_UPLINK        FlowDirection = "UPLINK"
	FlowDirection_BIDIRECTIONAL FlowDirection = "BIDIRECTIONAL"
	FlowDirection_UNSPECIFIED   FlowDirection = "UNSPECIFIED"
)
