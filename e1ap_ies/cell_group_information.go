package e1ap_ies

// CellGroupInformation From: 9_4_5_Information_Element_Definitions.txt:214
type CellGroupInformation []CellGroupInformationItem

// CellGroupInformationItem From: 9_4_5_Information_Element_Definitions.txt:216
type CellGroupInformationItem struct {
	CellGroupID     int64            `asn1:"mandatory,ext"`
	ULConfiguration *ULConfiguration `asn1:"optional,ext"`
	DLTXStop        *DLTXStop        `asn1:"optional,ext"`
	RATType         *RATType         `asn1:"optional,ext"`
}
