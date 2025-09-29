package e1ap_ies

// NRCGISupportItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1504
type NRCGISupportItem struct {
	NRCGI NRCGI `asn1:"mandatory"`
}
