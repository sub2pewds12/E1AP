package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// BearerContextModificationResponse is a generated SEQUENCE type.
type BearerContextModificationResponse struct {
	GNBCUCPUEE1APID                         GNBCUCPUEE1APID                          `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                         GNBCUUPUEE1APID                          `aper:"lb:0,ub:4294967295,mandatory,ext"`
	SystemBearerContextModificationResponse *SystemBearerContextModificationResponse `aper:"optional,ext"`
}

func (msg *BearerContextModificationResponse) toIes() ([]E1APMessageIE, error) {
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
	if msg.SystemBearerContextModificationResponse != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDSystemBearerContextModificationResponse,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.SystemBearerContextModificationResponse,
			})
		}
	}
	return ies, nil
}

// Encode function for BearerContextModificationResponse to be generated here.

// Decode function for BearerContextModificationResponse to be generated here.

// Decoder helper for BearerContextModificationResponse to be generated here.
