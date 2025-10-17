package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ExtendedGNBCUUPName is a generated SEQUENCE type.
type ExtendedGNBCUUPName struct {
	GNBCUUPNameVisibleString *aper.OctetString           `aper:"optional,ext"`
	GNBCUUPNameUTF8String    *aper.OctetString           `aper:"optional,ext"`
	IEExtensions             *ProtocolExtensionContainer `aper:"optional,ext"`
}
