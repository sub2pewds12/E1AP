package e1ap_ies

// NRCGISupportList From: 9_4_5_Information_Element_Definitions.txt:1502
type NRCGISupportList []NRCGISupportItem

// NRCGISupportItem From: 9_4_5_Information_Element_Definitions.txt:1504
type NRCGISupportItem struct {
	NRCGI NRCGI `asn1:"mandatory"`
}
