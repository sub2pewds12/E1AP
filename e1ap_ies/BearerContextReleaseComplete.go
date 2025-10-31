package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// BearerContextReleaseComplete is a generated SEQUENCE type.
type BearerContextReleaseComplete struct {
	GNBCUCPUEE1APID               GNBCUCPUEE1APID         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID               GNBCUUPUEE1APID         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	CriticalityDiagnostics        *CriticalityDiagnostics `aper:"optional,ext"`
	RetainabilityMeasurementsInfo []DRBRemovedItem        `aper:"optional,ext"`
}

func (msg *BearerContextReleaseComplete) toIes() ([]E1APMessageIE, error) {
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
	if msg.CriticalityDiagnostics != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDCriticalityDiagnostics,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.CriticalityDiagnostics,
			})
		}
	}
	if len(msg.RetainabilityMeasurementsInfo) > 0 {

		{

			tmp_RetainabilityMeasurementsInfo := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: 0},
				ext: false,
			}
			for i := 0; i < len(msg.RetainabilityMeasurementsInfo); i++ {
				tmp_RetainabilityMeasurementsInfo.Value = append(tmp_RetainabilityMeasurementsInfo.Value, &msg.RetainabilityMeasurementsInfo[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDRetainabilityMeasurementsInfo,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_RetainabilityMeasurementsInfo,
			})
		}
	}
	return ies, nil
}

// Encode function for BearerContextReleaseComplete to be generated here.

// Decode function for BearerContextReleaseComplete to be generated here.

// Decoder helper for BearerContextReleaseComplete to be generated here.
