package e1ap_ies

// EUTRANQOSSupportItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1044
type EUTRANQOSSupportItem struct {
	EUTRANQOS EUTRANQOS `asn1:"mandatory"`
}
