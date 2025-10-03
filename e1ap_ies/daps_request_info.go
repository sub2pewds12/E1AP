package e1ap_ies

// DAPSRequestInfo From: 9_4_5_Information_Element_Definitions.txt:307
// ASN.1 Data Type: SEQUENCE
type DAPSRequestInfo struct {
	DapsIndicator DAPSRequestInfoDapsIndicator `aper:"mandatory,ext"`
}
