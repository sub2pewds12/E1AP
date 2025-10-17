package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ExtendedGNBCUCPName is a generated SEQUENCE type.
type ExtendedGNBCUCPName struct {
	GNBCUCPNameVisibleString *aper.OctetString           `aper:"optional,ext"`
	GNBCUCPNameUTF8String    *aper.OctetString           `aper:"optional,ext"`
	IEExtensions             *ProtocolExtensionContainer `aper:"optional,ext"`
}
