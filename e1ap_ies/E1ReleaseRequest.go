package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// E1ReleaseRequest is a generated SEQUENCE type.
type E1ReleaseRequest struct {
	TransactionID TransactionID `aper:"lb:0,ub:255,mandatory,ext"`
	Cause         Cause         `aper:"mandatory,ext"`
}

func (msg *E1ReleaseRequest) toIes() ([]E1APMessageIE, error) {
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

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDCause,
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &msg.Cause,
		})
	}
	return ies, nil
}

// Encode function for E1ReleaseRequest to be generated here.

// Decode function for E1ReleaseRequest to be generated here.

// Decoder helper for E1ReleaseRequest to be generated here.
