package models

type MeteringMethod string

// List of MeteringMethod
const (
	MeteringMethod_DURATION        MeteringMethod = "DURATION"
	MeteringMethod_VOLUME          MeteringMethod = "VOLUME"
	MeteringMethod_DURATION_VOLUME MeteringMethod = "DURATION_VOLUME"
	MeteringMethod_EVENT           MeteringMethod = "EVENT"
)
