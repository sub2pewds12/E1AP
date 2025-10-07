package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBActivityItem is a generated SEQUENCE type.
type DRBActivityItem struct {
	DRBID        aper.Integer                `aper:"lb:1,ub:32,mandatory,ext"`
	DRBActivity  DRBActivity                 `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
