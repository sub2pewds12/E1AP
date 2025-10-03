package e1ap_ies

// SecurityInformation From: 9_4_5_Information_Element_Definitions.txt:2152
// ASN.1 Data Type: SEQUENCE
type SecurityInformation struct {
	SecurityAlgorithm SecurityAlgorithm `aper:"mandatory,ext"`
	UPSecuritykey     UPSecuritykey     `aper:"mandatory,ext"`
}
