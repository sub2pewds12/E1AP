package e1ap_ies

// Criticality represents the ASN.1 definition from 9_4_6_Common_Definitions.txt:34
type Criticality int32

const (
	Criticality_Reject Criticality = 0
	Criticality_Ignore Criticality = 1
	Criticality_Notify Criticality = 2
)
