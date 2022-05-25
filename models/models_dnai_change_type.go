package models

type DnaiChangeType string

// List of DnaiChangeType
const (
	DnaiChangeType_EARLY      DnaiChangeType = "EARLY"
	DnaiChangeType_EARLY_LATE DnaiChangeType = "EARLY_LATE"
	DnaiChangeType_LATE       DnaiChangeType = "LATE"
)
