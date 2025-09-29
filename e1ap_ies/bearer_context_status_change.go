package e1ap_ies

// BearerContextStatusChange represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:131
type BearerContextStatusChange int32

const (
	BearerContextStatusChange_Suspend BearerContextStatusChange = 0
	BearerContextStatusChange_Resume  BearerContextStatusChange = 1
)
