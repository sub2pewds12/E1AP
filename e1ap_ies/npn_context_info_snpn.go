package e1ap_ies

import "github.com/lvdund/ngap/aper"

// NPNContextInfoSNPN is a generated SEQUENCE type.
type NPNContextInfoSNPN struct {
	NID          aper.BitString              `aper:"lb:44,ub:44,mandatory"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional"`
}
