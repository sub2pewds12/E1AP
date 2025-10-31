package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUUPCounterCheckRequest is a generated SEQUENCE type.
type GNBCUUPCounterCheckRequest struct {
	GNBCUCPUEE1APID                  GNBCUCPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                  GNBCUUPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,ext"`
	SystemGNBCUUPCounterCheckRequest SystemGNBCUUPCounterCheckRequest `aper:"mandatory,ext"`
}

func (msg *GNBCUUPCounterCheckRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDGNBCUCPUEE1APID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(msg.GNBCUCPUEE1APID),
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDGNBCUUPUEE1APID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(msg.GNBCUUPUEE1APID),
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDSystemGNBCUUPCounterCheckRequest,
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &msg.SystemGNBCUUPCounterCheckRequest,
		})
	}
	return ies, nil
}

// Encode function for GNBCUUPCounterCheckRequest to be generated here.

// Decode function for GNBCUUPCounterCheckRequest to be generated here.

// Decoder helper for GNBCUUPCounterCheckRequest to be generated here.
