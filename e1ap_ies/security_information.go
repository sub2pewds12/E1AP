package e1ap_ies

// SecurityInformation From: 9_4_5_Information_Element_Definitions.txt:2152
type SecurityInformation struct {
	SecurityAlgorithm SecurityAlgorithm `asn1:"mandatory,ext"`
	UPSecuritykey     UPSecuritykey     `asn1:"mandatory,ext"`
}
