package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DLDataNotification is a generated SEQUENCE type.
type DLDataNotification struct {
	GNBCUCPUEE1APID        GNBCUCPUEE1APID          `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID        GNBCUUPUEE1APID          `aper:"lb:0,ub:4294967295,mandatory,ext"`
	PPI                    *PPI                     `aper:"lb:0,ub:7,optional,ext"`
	PDUSessionToNotifyList []PDUSessionToNotifyItem `aper:"optional,ext"`
}

func (msg *DLDataNotification) toIes() ([]E1APMessageIE, error) {
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
	if msg.PPI != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDPPI,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 7},
					ext:   true,
					Value: aper.Integer((*msg.PPI)),
				},
			})
		}
	}
	if len(msg.PDUSessionToNotifyList) > 0 {

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
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_PDUSessionToNotifyList,
			})
		}
	}
	return ies, nil
}

// Encode function for DLDataNotification to be generated here.

// Decode function for DLDataNotification to be generated here.

// Decoder helper for DLDataNotification to be generated here.
