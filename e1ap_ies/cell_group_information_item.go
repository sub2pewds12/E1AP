package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CellGroupInformationItem is a generated SEQUENCE type.
type CellGroupInformationItem struct {
	CellGroupID     aper.Integer                `aper:"lb:0,ub:3,mandatory,ext"`
	ULConfiguration *ULConfiguration            `aper:"optional,ext"`
	DLTXStop        *DLTXStop                   `aper:"optional,ext"`
	RATType         *RATType                    `aper:"optional,ext"`
	IEExtensions    *ProtocolExtensionContainer `aper:"optional,ext"`
}
