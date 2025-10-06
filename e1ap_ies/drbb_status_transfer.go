package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBBStatusTransfer is a generated SEQUENCE type.
type DRBBStatusTransfer struct {
	ReceiveStatusofPDCPSDU *aper.BitString             `aper:"lb:1,ub:131072,optional,ext"`
	CountValue             PDCPCount                   `aper:"mandatory,ext"`
	IEExtension            *ProtocolExtensionContainer `aper:"optional,ext"`
}
