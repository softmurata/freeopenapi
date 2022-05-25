package models

type PolicyControlRequestTrigger string

// List of PolicyControlRequestTrigger
const (
	PolicyControlRequestTrigger_PLMN_CH              PolicyControlRequestTrigger = "PLMN_CH"
	PolicyControlRequestTrigger_RES_MO_RE            PolicyControlRequestTrigger = "RES_MO_RE"
	PolicyControlRequestTrigger_AC_TY_CH             PolicyControlRequestTrigger = "AC_TY_CH"
	PolicyControlRequestTrigger_UE_IP_CH             PolicyControlRequestTrigger = "UE_IP_CH"
	PolicyControlRequestTrigger_UE_MAC_CH            PolicyControlRequestTrigger = "UE_MAC_CH"
	PolicyControlRequestTrigger_AN_CH_COR            PolicyControlRequestTrigger = "AN_CH_COR"
	PolicyControlRequestTrigger_US_RE                PolicyControlRequestTrigger = "US_RE"
	PolicyControlRequestTrigger_APP_STA              PolicyControlRequestTrigger = "APP_STA"
	PolicyControlRequestTrigger_APP_STO              PolicyControlRequestTrigger = "APP_STO"
	PolicyControlRequestTrigger_AN_INFO              PolicyControlRequestTrigger = "AN_INFO"
	PolicyControlRequestTrigger_CM_SES_FAIL          PolicyControlRequestTrigger = "CM_SES_FAIL"
	PolicyControlRequestTrigger_PS_DA_OFF            PolicyControlRequestTrigger = "PS_DA_OFF"
	PolicyControlRequestTrigger_DEF_QOS_CH           PolicyControlRequestTrigger = "DEF_QOS_CH"
	PolicyControlRequestTrigger_SE_AMBR_CH           PolicyControlRequestTrigger = "SE_AMBR_CH"
	PolicyControlRequestTrigger_QOS_NOTIF            PolicyControlRequestTrigger = "QOS_NOTIF"
	PolicyControlRequestTrigger_NO_CREDIT            PolicyControlRequestTrigger = "NO_CREDIT"
	PolicyControlRequestTrigger_PRA_CH               PolicyControlRequestTrigger = "PRA_CH"
	PolicyControlRequestTrigger_SAREA_CH             PolicyControlRequestTrigger = "SAREA_CH"
	PolicyControlRequestTrigger_SCNN_CH              PolicyControlRequestTrigger = "SCNN_CH"
	PolicyControlRequestTrigger_RE_TIMEOUT           PolicyControlRequestTrigger = "RE_TIMEOUT"
	PolicyControlRequestTrigger_RES_RELEASE          PolicyControlRequestTrigger = "RES_RELEASE"
	PolicyControlRequestTrigger_SUCC_RES_ALLO        PolicyControlRequestTrigger = "SUCC_RES_ALLO"
	PolicyControlRequestTrigger_RAT_TY_CH            PolicyControlRequestTrigger = "RAT_TY_CH"
	PolicyControlRequestTrigger_REF_QOS_IND_CH       PolicyControlRequestTrigger = "REF_QOS_IND_CH"
	PolicyControlRequestTrigger_NUM_OF_PACKET_FILTER PolicyControlRequestTrigger = "NUM_OF_PACKET_FILTER"
	PolicyControlRequestTrigger_UE_STATUS_RESUME     PolicyControlRequestTrigger = "UE_STATUS_RESUME"
	PolicyControlRequestTrigger_UE_TZ_CH             PolicyControlRequestTrigger = "UE_TZ_CH"
)