package e1ap_ies

// PrivacyIndicator represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1926
type PrivacyIndicator int32

const (
	PrivacyIndicator_ImmediateMDT PrivacyIndicator = 0
	PrivacyIndicator_LoggedMDT    PrivacyIndicator = 1
)
