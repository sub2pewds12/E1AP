package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PrivateIEField is a generated SEQUENCE type.
type PrivateIEField struct {
	ID          PrivateIEID      `aper:"mandatory"`
	Criticality Criticality      `aper:"mandatory"`
	Value       aper.OctetString `aper:"mandatory"`
}
