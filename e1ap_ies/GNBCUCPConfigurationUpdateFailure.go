package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUCPConfigurationUpdateFailure is a generated SEQUENCE type.
type GNBCUCPConfigurationUpdateFailure struct {
	TransactionID          TransactionID           `aper:"lb:0,ub:255,mandatory,ext"`
	Cause                  Cause                   `aper:"mandatory,ext"`
	TimeToWait             *TimeToWait             `aper:"optional,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ext"`
}

func (msg *GNBCUCPConfigurationUpdateFailure) toIes() ([]E1APMessageIE, error) {
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
	if msg.TimeToWait != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDTimeToWait,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 5},
					ext:   true,
					Value: (*msg.TimeToWait).Value,
				},
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

// Encode function for GNBCUCPConfigurationUpdateFailure to be generated here.

// Decode function for GNBCUCPConfigurationUpdateFailure to be generated here.

// Decoder helper for GNBCUCPConfigurationUpdateFailure to be generated here.
