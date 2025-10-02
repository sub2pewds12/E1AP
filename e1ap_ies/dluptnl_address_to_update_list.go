package e1ap_ies

// DLUPTNLAddressToUpdateList From: 9_4_4_PDU_Definitions.txt:1441
type DLUPTNLAddressToUpdateList []DLUPTNLAddressToUpdateItem

// DLUPTNLAddressToUpdateItem From: 9_4_5_Information_Element_Definitions.txt:402
type DLUPTNLAddressToUpdateItem struct {
	OldTNLAdress []byte `asn1:"mandatory,ext"`
	NewTNLAdress []byte `asn1:"mandatory,ext"`
}
