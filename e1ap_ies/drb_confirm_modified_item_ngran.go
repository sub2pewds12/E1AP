package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBConfirmModifiedItemNGRAN is a generated SEQUENCE type.
type DRBConfirmModifiedItemNGRAN struct {
	DRBID                aper.Integer                `aper:"lb:1,ub:32,mandatory,ext"`
	CellGroupInformation []CellGroupInformationItem  `aper:"optional,ext"`
	IEExtensions         *ProtocolExtensionContainer `aper:"optional,ext"`
}
