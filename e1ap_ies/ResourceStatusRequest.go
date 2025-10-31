package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ResourceStatusRequest is a generated SEQUENCE type.
type ResourceStatusRequest struct {
	TransactionID         TransactionID                                   `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPMeasurementID  ResourceStatusRequestIEsIDGNBCUCPMeasurementID  `aper:"lb:1,ub:4095,mandatory,ext"`
	GNBCUUPMeasurementID  *ResourceStatusRequestIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,optional,ext"`
	RegistrationRequest   RegistrationRequest                             `aper:"mandatory,ext"`
	ReportCharacteristics *ReportCharacteristics                          `aper:"lb:36,ub:36,conditional,ext"`
	ReportingPeriodicity  *ReportingPeriodicity                           `aper:"optional,ext"`
}

func (msg *ResourceStatusRequest) toIes() ([]E1APMessageIE, error) {
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
			Id:          ProtocolIEIDRegistrationRequest,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 1},
				ext:   true,
				Value: msg.RegistrationRequest.Value,
			},
		})
	}
	if msg.ReportCharacteristics != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDReportCharacteristics,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 36, Ub: 36},
					ext:   false,
					Value: aper.OctetString((*msg.ReportCharacteristics)),
				},
			})
		}
	}
	if msg.ReportingPeriodicity != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDReportingPeriodicity,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 15},
					ext:   true,
					Value: (*msg.ReportingPeriodicity).Value,
				},
			})
		}
	}
	return ies, nil
}

// Encode function for ResourceStatusRequest to be generated here.

// Decode function for ResourceStatusRequest to be generated here.

// Decoder helper for ResourceStatusRequest to be generated here.
