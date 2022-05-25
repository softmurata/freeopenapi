package models

type RequestedRuleDataType string

// List of RequestedRuleDataType
const (
	RequestedRuleDataType_CH_ID         RequestedRuleDataType = "CH_ID"
	RequestedRuleDataType_MS_TIME_ZONE  RequestedRuleDataType = "MS_TIME_ZONE"
	RequestedRuleDataType_USER_LOC_INFO RequestedRuleDataType = "USER_LOC_INFO"
	RequestedRuleDataType_RES_RELEASE   RequestedRuleDataType = "RES_RELEASE"
	RequestedRuleDataType_SUCC_RES_ALLO RequestedRuleDataType = "SUCC_RES_ALLO"
)
