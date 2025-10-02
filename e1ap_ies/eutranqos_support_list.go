package e1ap_ies

// EUTRANQOSSupportList From: 9_4_5_Information_Element_Definitions.txt:1042
type EUTRANQOSSupportList []EUTRANQOSSupportItem

// EUTRANQOSSupportItem From: 9_4_5_Information_Element_Definitions.txt:1044
type EUTRANQOSSupportItem struct {
	EUTRANQOS EUTRANQOS `asn1:"mandatory"`
}
