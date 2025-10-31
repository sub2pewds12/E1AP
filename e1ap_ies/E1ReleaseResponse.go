package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// E1ReleaseResponse is a generated SEQUENCE type.
type E1ReleaseResponse struct {
	TransactionID TransactionID `aper:"lb:0,ub:255,mandatory,ext"`
}

func (msg *E1ReleaseResponse) toIes() ([]E1APMessageIE, error) {
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
	return ies, nil
}

// Encode function for E1ReleaseResponse to be generated here.

// Decode function for E1ReleaseResponse to be generated here.

// Decoder helper for E1ReleaseResponse to be generated here.
