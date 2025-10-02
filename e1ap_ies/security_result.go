package e1ap_ies

// SecurityResult From: 9_4_5_Information_Element_Definitions.txt:2163
type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult       `asn1:"mandatory,ext"`
	ConfidentialityProtectionResult ConfidentialityProtectionResult `asn1:"mandatory,ext"`
}
