package models

type AccessTech string

// List of AccessTech
const (
	AccessTech_NR                                AccessTech = "NR"
	AccessTech_EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE AccessTech = "EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE"
	AccessTech_EUTRAN_IN_NBS1_MODE_ONLY          AccessTech = "EUTRAN_IN_NBS1_MODE_ONLY"
	AccessTech_EUTRAN_IN_WBS1_MODE_ONLY          AccessTech = "EUTRAN_IN_WBS1_MODE_ONLY"
	AccessTech_UTRAN                             AccessTech = "UTRAN"
	AccessTech_GSM_AND_ECGSM_IO_T                AccessTech = "GSM_AND_ECGSM_IoT"
	AccessTech_GSM_WITHOUT_ECGSM_IO_T            AccessTech = "GSM_WITHOUT_ECGSM_IoT"
	AccessTech_ECGSM_IO_T_ONLY                   AccessTech = "ECGSM_IoT_ONLY"
	AccessTech_CDMA_1X_RTT                       AccessTech = "CDMA_1xRTT"
	AccessTech_CDMA_HRPD                         AccessTech = "CDMA_HRPD"
	AccessTech_GSM_COMPACT                       AccessTech = "GSM_COMPACT"
)
