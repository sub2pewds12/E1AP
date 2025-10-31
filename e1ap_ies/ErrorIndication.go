package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ErrorIndication is a generated SEQUENCE type.
type ErrorIndication struct {
	TransactionID          TransactionID           `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPUEE1APID        *GNBCUCPUEE1APID        `aper:"lb:0,ub:4294967295,optional,ext"`
	GNBCUUPUEE1APID        *GNBCUUPUEE1APID        `aper:"lb:0,ub:4294967295,optional,ext"`
	Cause                  *Cause                  `aper:"optional,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ext"`
}

func (msg *ErrorIndication) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDTransactionID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 255},
				ext:   true,
				Value: aper.Integer(msg.TransactionID),
			},
		})
	}
	if msg.GNBCUCPUEE1APID != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUCPUEE1APID,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: aper.Integer((*msg.GNBCUCPUEE1APID)),
				},
			})
		}
	}
	if msg.GNBCUUPUEE1APID != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUUPUEE1APID,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: aper.Integer((*msg.GNBCUUPUEE1APID)),
				},
			})
		}
	}
	if msg.Cause != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDCause,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.Cause,
			})
		}
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
	return ies, nil
}

// Encode function for ErrorIndication to be generated here.

// Decode function for ErrorIndication to be generated here.

// Decoder helper for ErrorIndication to be generated here.
