package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// BearerContextSetupResponse is a generated SEQUENCE type.
type BearerContextSetupResponse struct {
	GNBCUCPUEE1APID                  GNBCUCPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                  GNBCUUPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,ext"`
	SystemBearerContextSetupResponse SystemBearerContextSetupResponse `aper:"mandatory,ext"`
}

func (msg *BearerContextSetupResponse) toIes() ([]E1APMessageIE, error) {
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
			Id:          ProtocolIEIDSystemBearerContextSetupResponse,
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &msg.SystemBearerContextSetupResponse,
		})
	}
	return ies, nil
}

// Encode function for BearerContextSetupResponse to be generated here.

// Decode function for BearerContextSetupResponse to be generated here.

// Decoder helper for BearerContextSetupResponse to be generated here.
