package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ProtocolIEField is a generated SEQUENCE type.
type ProtocolIEField struct {
	ID          ProtocolIEID     `aper:"lb:0,ub:MaxProtocolIEs,mandatory"`
	Criticality Criticality      `aper:"mandatory"`
	Value       aper.OctetString `aper:"mandatory"`
}
