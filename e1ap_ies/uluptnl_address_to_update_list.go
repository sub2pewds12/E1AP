package e1ap_ies

// ULUPTNLAddressToUpdateList From: 9_4_4_PDU_Definitions.txt:1461
type ULUPTNLAddressToUpdateList []ULUPTNLAddressToUpdateItem

// ULUPTNLAddressToUpdateItem From: 9_4_5_Information_Element_Definitions.txt:2382
type ULUPTNLAddressToUpdateItem struct {
	OldTNLAdress []byte `asn1:"mandatory,ext"`
	NewTNLAdress []byte `asn1:"mandatory,ext"`
}
