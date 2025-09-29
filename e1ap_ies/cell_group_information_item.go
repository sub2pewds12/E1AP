package e1ap_ies

// CellGroupInformationItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:216
type CellGroupInformationItem struct {
	CellGroupID     int64            `asn1:"lb:0,ub:3,mandatory,ext"`
	ULConfiguration *ULConfiguration `asn1:"optional,ext"`
	DLTXStop        *DLTXStop        `asn1:"optional,ext"`
	RATType         *RATType         `asn1:"optional,ext"`
}
