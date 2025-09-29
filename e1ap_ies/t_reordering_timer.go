package e1ap_ies

// TReorderingTimer represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2304
type TReorderingTimer struct {
	TReordering TReordering `asn1:"mandatory,ext"`
}
