package e1ap_ies

// DLUPTNLAddressToUpdateItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:402
type DLUPTNLAddressToUpdateItem struct {
	OldTNLAdress []byte `asn1:"lb:1,ub:160,mandatory,ext"`
	NewTNLAdress []byte `asn1:"lb:1,ub:160,mandatory,ext"`
}
