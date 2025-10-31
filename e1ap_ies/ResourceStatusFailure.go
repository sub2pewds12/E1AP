package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ResourceStatusFailure is a generated SEQUENCE type.
type ResourceStatusFailure struct {
	TransactionID          TransactionID                                   `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPMeasurementID   ResourceStatusFailureIEsIDGNBCUCPMeasurementID  `aper:"lb:1,ub:4095,mandatory,ext"`
	GNBCUUPMeasurementID   *ResourceStatusFailureIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,optional,ext"`
	Cause                  Cause                                           `aper:"mandatory,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics                         `aper:"optional,ext"`
}

func (msg *ResourceStatusFailure) toIes() ([]E1APMessageIE, error) {
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
			Id:          ProtocolIEIDGNBCUCPMeasurementID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 4095},
				ext:   true,
				Value: aper.Integer(msg.GNBCUCPMeasurementID),
			},
		})
	}
	if msg.GNBCUUPMeasurementID != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUUPMeasurementID,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 1, Ub: 4095},
					ext:   true,
					Value: aper.Integer((*msg.GNBCUUPMeasurementID)),
				},
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

// Encode function for ResourceStatusFailure to be generated here.

// Decode function for ResourceStatusFailure to be generated here.

// Decoder helper for ResourceStatusFailure to be generated here.
