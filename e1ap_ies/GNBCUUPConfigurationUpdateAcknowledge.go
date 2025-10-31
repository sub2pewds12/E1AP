package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUUPConfigurationUpdateAcknowledge is a generated SEQUENCE type.
type GNBCUUPConfigurationUpdateAcknowledge struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	CriticalityDiagnostics    *CriticalityDiagnostics    `aper:"optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
}

func (msg *GNBCUUPConfigurationUpdateAcknowledge) toIes() ([]E1APMessageIE, error) {
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
	if msg.CriticalityDiagnostics != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDCriticalityDiagnostics,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.CriticalityDiagnostics,
			})
		}
	}
	if msg.TransportLayerAddressInfo != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDTransportLayerAddressInfo,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.TransportLayerAddressInfo,
			})
		}
	}
	return ies, nil
}

// Encode function for GNBCUUPConfigurationUpdateAcknowledge to be generated here.

// Decode function for GNBCUUPConfigurationUpdateAcknowledge to be generated here.

// Decoder helper for GNBCUUPConfigurationUpdateAcknowledge to be generated here.
