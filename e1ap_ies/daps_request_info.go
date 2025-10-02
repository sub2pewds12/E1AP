package e1ap_ies

// DAPSRequestInfo From: 9_4_5_Information_Element_Definitions.txt:307
type DAPSRequestInfo struct {
	DapsIndicator DAPSRequestInfoDapsIndicator `asn1:"mandatory,ext"`
}
