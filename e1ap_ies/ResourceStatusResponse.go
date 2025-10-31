package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ResourceStatusResponse is a generated SEQUENCE type.
type ResourceStatusResponse struct {
	TransactionID          TransactionID                                   `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPMeasurementID   ResourceStatusResponseIEsIDGNBCUCPMeasurementID `aper:"lb:1,ub:4095,mandatory,ext"`
	GNBCUUPMeasurementID   ResourceStatusResponseIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,mandatory,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics                         `aper:"optional,ext"`
}

func (msg *ResourceStatusResponse) toIes() ([]E1APMessageIE, error) {
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

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDGNBCUUPMeasurementID,
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 4095},
				ext:   true,
				Value: aper.Integer(msg.GNBCUUPMeasurementID),
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
	return ies, nil
}

// Encode function for ResourceStatusResponse to be generated here.

// Decode function for ResourceStatusResponse to be generated here.

// Decoder helper for ResourceStatusResponse to be generated here.
