package models

type FlowDirectionRm string

// List of FlowDirectionRm
const (
	FlowDirectionRm_DOWNLINK      FlowDirectionRm = "DOWNLINK"
	FlowDirectionRm_UPLINK        FlowDirectionRm = "UPLINK"
	FlowDirectionRm_BIDIRECTIONAL FlowDirectionRm = "BIDIRECTIONAL"
	FlowDirectionRm_UNSPECIFIED   FlowDirectionRm = "UNSPECIFIED"
)
