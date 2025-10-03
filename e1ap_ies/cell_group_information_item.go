package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CellGroupInformationItem From: 9_4_5_Information_Element_Definitions.txt:216
// ASN.1 Data Type: SEQUENCE
type CellGroupInformationItem struct {
	CellGroupID     aper.Integer     `aper:"mandatory,ext"`
	ULConfiguration *ULConfiguration `aper:"optional,ext"`
	DLTXStop        *DLTXStop        `aper:"optional,ext"`
	RATType         *RATType         `aper:"optional,ext"`
}
