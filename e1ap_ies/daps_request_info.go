package e1ap_ies

// DAPSRequestInfo represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:307
type DAPSRequestInfo struct {
	DapsIndicator DAPSRequestInfoDapsIndicator `asn1:"mandatory,ext"`
}
