package e1ap_ies

// IntegrityProtectionResult represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1258
type IntegrityProtectionResult int32

const (
	IntegrityProtectionResult_Performed    IntegrityProtectionResult = 0
	IntegrityProtectionResult_NotPerformed IntegrityProtectionResult = 1
)
