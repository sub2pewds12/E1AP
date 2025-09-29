package e1ap_ies

// SliceSupportItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2176
type SliceSupportItem struct {
	SNSSAI SNSSAI `asn1:"mandatory"`
}
