package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUCPE1SetupResponse is a generated SEQUENCE type.
type GNBCUCPE1SetupResponse struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUUPID                 GNBCUUPID                  `aper:"lb:0,ub:68719476735,mandatory,ext"`
	GNBCUUPName               *aper.OctetString          `aper:"optional,ext"`
	CNSupport                 CNSupport                  `aper:"mandatory,ext"`
	SupportedPLMNs            []SupportedPLMNsItem       `aper:"lb:1,ub:MaxnoofSPLMNs,mandatory,ext"`
	GNBCUUPCapacity           *GNBCUUPCapacity           `aper:"lb:0,ub:255,optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
	ExtendedGNBCUUPName       *ExtendedGNBCUUPName       `aper:"optional,ext"`
}

func (msg *GNBCUCPE1SetupResponse) toIes() ([]E1APMessageIE, error) {
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
			Id:          ProtocolIEIDGNBCUUPID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 68719476735},
				ext:   false,
				Value: aper.Integer(msg.GNBCUUPID),
			},
		})
	}
	if msg.GNBCUUPName != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUUPName,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   false,
					Value: aper.OctetString((*msg.GNBCUUPName)),
				},
			})
		}
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDCNSupport,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 2},
				ext:   true,
				Value: msg.CNSupport.Value,
			},
		})
	}

	{

		tmp_SupportedPLMNs := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 1, Ub: MaxnoofSPLMNs},
			ext: false,
		}
		for i := 0; i < len(msg.SupportedPLMNs); i++ {
			tmp_SupportedPLMNs.Value = append(tmp_SupportedPLMNs.Value, &msg.SupportedPLMNs[i])
		}
		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDSupportedPLMNs,
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &tmp_SupportedPLMNs,
		})
	}
	if msg.GNBCUUPCapacity != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUUPCapacity,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 255},
					ext:   false,
					Value: aper.Integer((*msg.GNBCUUPCapacity)),
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
	if msg.ExtendedGNBCUUPName != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDExtendedGNBCUUPName,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.ExtendedGNBCUUPName,
			})
		}
	}
	return ies, nil
}

// Encode function for GNBCUCPE1SetupResponse to be generated here.

// Decode function for GNBCUCPE1SetupResponse to be generated here.

// Decoder helper for GNBCUCPE1SetupResponse to be generated here.
