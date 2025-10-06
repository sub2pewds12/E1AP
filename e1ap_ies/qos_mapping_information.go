package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSMappingInformation is a generated SEQUENCE type.
type QOSMappingInformation struct {
	Dscp      *aper.BitString `aper:"lb:6,ub:6,optional,ext"`
	FlowLabel *aper.BitString `aper:"lb:20,ub:20,optional,ext"`
}
