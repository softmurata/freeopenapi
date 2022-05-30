package models

type RequestedQosMonitoringParameter string

const (
	RequestedQosMonitoringParameter_DOWNLINK   RequestedQosMonitoringParameter = "DOWNLINK"
	RequestedQosMonitoringParameter_UPLINK     RequestedQosMonitoringParameter = "UPLINK"
	RequestedQosMonitoringParameter_ROUND_TRIP RequestedQosMonitoringParameter = "ROUND_TRIP"
)
