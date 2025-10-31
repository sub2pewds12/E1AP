package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUCPE1SetupRequest is a generated SEQUENCE type.
type GNBCUCPE1SetupRequest struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPName               *aper.OctetString          `aper:"optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
	ExtendedGNBCUCPName       *ExtendedGNBCUCPName       `aper:"optional,ext"`
}

func (msg *GNBCUCPE1SetupRequest) toIes() ([]E1APMessageIE, error) {
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
	if msg.GNBCUCPName != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUCPName,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   false,
					Value: aper.OctetString((*msg.GNBCUCPName)),
				},
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
	if msg.ExtendedGNBCUCPName != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDExtendedGNBCUCPName,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.ExtendedGNBCUCPName,
			})
		}
	}
	return ies, nil
}

// Encode function for GNBCUCPE1SetupRequest to be generated here.

// Decode function for GNBCUCPE1SetupRequest to be generated here.

// Decoder helper for GNBCUCPE1SetupRequest to be generated here.
