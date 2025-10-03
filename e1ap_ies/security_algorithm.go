package e1ap_ies

// SecurityAlgorithm From: 9_4_5_Information_Element_Definitions.txt:2129
// ASN.1 Data Type: SEQUENCE
type SecurityAlgorithm struct {
	CipheringAlgorithm           CipheringAlgorithm            `aper:"mandatory,ext"`
	IntegrityProtectionAlgorithm *IntegrityProtectionAlgorithm `aper:"optional,ext"`
}
