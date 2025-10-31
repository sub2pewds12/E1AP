package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ULDataNotification is a generated SEQUENCE type.
type ULDataNotification struct {
	GNBCUCPUEE1APID        GNBCUCPUEE1APID          `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID        GNBCUUPUEE1APID          `aper:"lb:0,ub:4294967295,mandatory,ext"`
	PDUSessionToNotifyList []PDUSessionToNotifyItem `aper:"mandatory,ext"`
}

func (msg *ULDataNotification) toIes() ([]E1APMessageIE, error) {
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

		tmp_PDUSessionToNotifyList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for i := 0; i < len(msg.PDUSessionToNotifyList); i++ {
			tmp_PDUSessionToNotifyList.Value = append(tmp_PDUSessionToNotifyList.Value, &msg.PDUSessionToNotifyList[i])
		}
		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDPDUSessionToNotifyList,
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &tmp_PDUSessionToNotifyList,
		})
	}
	return ies, nil
}

// Encode function for ULDataNotification to be generated here.

// Decode function for ULDataNotification to be generated here.

// Decoder helper for ULDataNotification to be generated here.
