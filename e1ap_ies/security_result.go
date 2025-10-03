package e1ap_ies

// SecurityResult From: 9_4_5_Information_Element_Definitions.txt:2163
// ASN.1 Data Type: SEQUENCE
type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult       `aper:"mandatory,ext"`
	ConfidentialityProtectionResult ConfidentialityProtectionResult `aper:"mandatory,ext"`
}
