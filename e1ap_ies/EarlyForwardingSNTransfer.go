package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// EarlyForwardingSNTransfer is a generated SEQUENCE type.
type EarlyForwardingSNTransfer struct {
	GNBCUCPUEE1APID                  GNBCUCPUEE1APID                    `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                  GNBCUUPUEE1APID                    `aper:"lb:0,ub:4294967295,mandatory,ext"`
	DRBsSubjectToEarlyForwardingList []DRBsSubjectToEarlyForwardingItem `aper:"mandatory,ext"`
}

func (msg *EarlyForwardingSNTransfer) toIes() ([]E1APMessageIE, error) {
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

		tmp_DRBsSubjectToEarlyForwardingList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for i := 0; i < len(msg.DRBsSubjectToEarlyForwardingList); i++ {
			tmp_DRBsSubjectToEarlyForwardingList.Value = append(tmp_DRBsSubjectToEarlyForwardingList.Value, &msg.DRBsSubjectToEarlyForwardingList[i])
		}
		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDDRBsSubjectToEarlyForwardingList,
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &tmp_DRBsSubjectToEarlyForwardingList,
		})
	}
	return ies, nil
}

// Encode function for EarlyForwardingSNTransfer to be generated here.

// Decode function for EarlyForwardingSNTransfer to be generated here.

// Decoder helper for EarlyForwardingSNTransfer to be generated here.
