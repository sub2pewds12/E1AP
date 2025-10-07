package e1ap_ies

import "github.com/lvdund/ngap/aper"

// SNSSAI is a generated SEQUENCE type.
type SNSSAI struct {
	SST          aper.OctetString            `aper:"lb:1,ub:1,mandatory,ext"`
	SD           *aper.OctetString           `aper:"lb:3,ub:3,optional,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
