package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUUPStatusIndication is a generated SEQUENCE type.
type GNBCUUPStatusIndication struct {
	TransactionID              TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUUPOverloadInformation GNBCUUPOverloadInformation `aper:"mandatory,ext"`
}

func (msg *GNBCUUPStatusIndication) toIes() ([]E1APMessageIE, error) {
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
			Id:          ProtocolIEIDGNBCUUPOverloadInformation,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 1},
				ext:   false,
				Value: msg.GNBCUUPOverloadInformation.Value,
			},
		})
	}
	return ies, nil
}

// Encode function for GNBCUUPStatusIndication to be generated here.

// Decode function for GNBCUUPStatusIndication to be generated here.

// Decoder helper for GNBCUUPStatusIndication to be generated here.
