package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUCPConfigurationUpdateAcknowledge is a generated SEQUENCE type.
type GNBCUCPConfigurationUpdateAcknowledge struct {
	TransactionID                TransactionID                  `aper:"lb:0,ub:255,mandatory,ext"`
	CriticalityDiagnostics       *CriticalityDiagnostics        `aper:"optional,ext"`
	GNBCUCPTNLASetupList         []GNBCUCPTNLASetupItem         `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	GNBCUCPTNLAFailedToSetupList []GNBCUCPTNLAFailedToSetupItem `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	TransportLayerAddressInfo    *TransportLayerAddressInfo     `aper:"optional,ext"`
}

func (msg *GNBCUCPConfigurationUpdateAcknowledge) toIes() ([]E1APMessageIE, error) {
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
	if len(msg.GNBCUCPTNLASetupList) > 0 {

		{

			tmp_GNBCUCPTNLASetupList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
				ext: false,
			}
			for i := 0; i < len(msg.GNBCUCPTNLASetupList); i++ {
				tmp_GNBCUCPTNLASetupList.Value = append(tmp_GNBCUCPTNLASetupList.Value, &msg.GNBCUCPTNLASetupList[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUCPTNLASetupList,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_GNBCUCPTNLASetupList,
			})
		}
	}
	if len(msg.GNBCUCPTNLAFailedToSetupList) > 0 {

		{

			tmp_GNBCUCPTNLAFailedToSetupList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
				ext: false,
			}
			for i := 0; i < len(msg.GNBCUCPTNLAFailedToSetupList); i++ {
				tmp_GNBCUCPTNLAFailedToSetupList.Value = append(tmp_GNBCUCPTNLAFailedToSetupList.Value, &msg.GNBCUCPTNLAFailedToSetupList[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUCPTNLAFailedToSetupList,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_GNBCUCPTNLAFailedToSetupList,
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

// Encode function for GNBCUCPConfigurationUpdateAcknowledge to be generated here.

// Decode function for GNBCUCPConfigurationUpdateAcknowledge to be generated here.

// Decoder helper for GNBCUCPConfigurationUpdateAcknowledge to be generated here.
