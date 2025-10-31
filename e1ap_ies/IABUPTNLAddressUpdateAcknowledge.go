package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// IABUPTNLAddressUpdateAcknowledge is a generated SEQUENCE type.
type IABUPTNLAddressUpdateAcknowledge struct {
	TransactionID              TransactionID                `aper:"lb:0,ub:255,mandatory,ext"`
	CriticalityDiagnostics     *CriticalityDiagnostics      `aper:"optional,ext"`
	ULUPTNLAddressToUpdateList []ULUPTNLAddressToUpdateItem `aper:"ub:MaxnoofTNLAddresses,optional,ext"`
}

func (msg *IABUPTNLAddressUpdateAcknowledge) toIes() ([]E1APMessageIE, error) {
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
	if len(msg.ULUPTNLAddressToUpdateList) > 0 {

		{

			tmp_ULUPTNLAddressToUpdateList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAddresses},
				ext: false,
			}
			for i := 0; i < len(msg.ULUPTNLAddressToUpdateList); i++ {
				tmp_ULUPTNLAddressToUpdateList.Value = append(tmp_ULUPTNLAddressToUpdateList.Value, &msg.ULUPTNLAddressToUpdateList[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDULUPTNLAddressToUpdateList,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_ULUPTNLAddressToUpdateList,
			})
		}
	}
	return ies, nil
}

// Encode function for IABUPTNLAddressUpdateAcknowledge to be generated here.

// Decode function for IABUPTNLAddressUpdateAcknowledge to be generated here.

// Decoder helper for IABUPTNLAddressUpdateAcknowledge to be generated here.
