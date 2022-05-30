package models

type ReportingFrequency string

const (
	ReportingFrequency_EVENT_TRIGGERED ReportingFrequency = "EVENT_TRIGGERED"
	ReportingFrequency_PERIODIC        ReportingFrequency = "PERIODIC"
	ReportingFrequency_SESSION_RELEASE ReportingFrequency = "SESSION_RELEASE"
)
