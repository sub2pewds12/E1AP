package e1ap_ies

// ConfidentialityProtectionResult represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:259
type ConfidentialityProtectionResult int32

const (
	ConfidentialityProtectionResult_Performed    ConfidentialityProtectionResult = 0
	ConfidentialityProtectionResult_NotPerformed ConfidentialityProtectionResult = 1
)
