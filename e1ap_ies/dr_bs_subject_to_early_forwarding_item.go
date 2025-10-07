package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBsSubjectToEarlyForwardingItem is a generated SEQUENCE type.
type DRBsSubjectToEarlyForwardingItem struct {
	DRBID        aper.Integer                `aper:"lb:1,ub:32,mandatory,ext"`
	DLCountValue PDCPCount                   `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
