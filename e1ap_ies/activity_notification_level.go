package e1ap_ies

// ActivityNotificationLevel represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:96
type ActivityNotificationLevel int32

const (
	ActivityNotificationLevel_DRB        ActivityNotificationLevel = 0
	ActivityNotificationLevel_PDUSession ActivityNotificationLevel = 1
	ActivityNotificationLevel_Ue         ActivityNotificationLevel = 2
)
