package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// BearerContextReleaseRequest is a generated SEQUENCE type.
type BearerContextReleaseRequest struct {
	GNBCUCPUEE1APID GNBCUCPUEE1APID `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID GNBCUUPUEE1APID `aper:"lb:0,ub:4294967295,mandatory,ext"`
	DRBStatusList   []DRBStatusItem `aper:"lb:1,ub:MaxnoofDRBs,optional,ext"`
	Cause           Cause           `aper:"mandatory,ext"`
}

func (msg *BearerContextReleaseRequest) toIes() ([]E1APMessageIE, error) {
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
	if len(msg.DRBStatusList) > 0 {

		{

			tmp_DRBStatusList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 1, Ub: MaxnoofDRBs},
				ext: false,
			}
			for i := 0; i < len(msg.DRBStatusList); i++ {
				tmp_DRBStatusList.Value = append(tmp_DRBStatusList.Value, &msg.DRBStatusList[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDDRBStatusList,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_DRBStatusList,
			})
		}
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDCause,
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &msg.Cause,
		})
	}
	return ies, nil
}

// Encode function for BearerContextReleaseRequest to be generated here.

// Decode function for BearerContextReleaseRequest to be generated here.

// Decoder helper for BearerContextReleaseRequest to be generated here.
