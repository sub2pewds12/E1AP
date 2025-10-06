package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ProtocolExtensionField is a generated SEQUENCE type.
type ProtocolExtensionField struct {
	ID             aper.Integer     `aper:"lb:0,ub:MaxProtocolIEs,mandatory"`
	Criticality    Criticality      `aper:"mandatory"`
	ExtensionValue aper.OctetString `aper:"mandatory"`
}
