package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DataForwardingInformationRequest is a generated SEQUENCE type.
type DataForwardingInformationRequest struct {
	DataForwardingRequest         DataForwardingRequest       `aper:"mandatory,ext"`
	QOSFlowsForwardedOnFwdTunnels []QOSFlowMappingItem        `aper:"optional,ext"`
	IEExtensions                  *ProtocolExtensionContainer `aper:"optional,ext"`
}

func (msg *DataForwardingInformationRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID,
			Criticality: Criticality{Value: Criticality},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 2},
				ext:   true,
				Value: msg.DataForwardingRequest.Value,
			},
		})
	}
	if len(msg.QOSFlowsForwardedOnFwdTunnels) > 0 {

		{

			tmp_QOSFlowsForwardedOnFwdTunnels := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: 0},
				ext: false,
			}
			for i := 0; i < len(msg.QOSFlowsForwardedOnFwdTunnels); i++ {
				tmp_QOSFlowsForwardedOnFwdTunnels.Value = append(tmp_QOSFlowsForwardedOnFwdTunnels.Value, &msg.QOSFlowsForwardedOnFwdTunnels[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID,
				Criticality: Criticality{Value: Criticality},
				Value:       &tmp_QOSFlowsForwardedOnFwdTunnels,
			})
		}
	}
	return ies, nil
}

// Encode function for DataForwardingInformationRequest to be generated here.

// Decode function for DataForwardingInformationRequest to be generated here.

// Decoder helper for DataForwardingInformationRequest to be generated here.
