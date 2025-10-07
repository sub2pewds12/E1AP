package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TSCTrafficInformation is a generated SEQUENCE type.
type TSCTrafficInformation struct {
	Periodicity      aper.Integer                `aper:"lb:1,ub:640000,mandatory"`
	BurstArrivalTime *aper.OctetString           `aper:"optional"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional"`
}
