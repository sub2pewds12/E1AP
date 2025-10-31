package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SecurityIndication is a generated SEQUENCE type.
type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication       `aper:"mandatory,ext"`
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication `aper:"mandatory,ext"`
	MaximumIPdatarate                   *MaximumIPdatarate                  `aper:"optional,ext"`
	IEExtensions                        *ProtocolExtensionContainer         `aper:"optional,ext"`
}

func (msg *SecurityIndication) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID,
			Criticality: Criticality{Value: Criticality},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 2},
				ext:   true,
				Value: msg.IntegrityProtectionIndication.Value,
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID,
			Criticality: Criticality{Value: Criticality},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 2},
				ext:   true,
				Value: msg.ConfidentialityProtectionIndication.Value,
			},
		})
	}
	if msg.MaximumIPdatarate != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID,
				Criticality: Criticality{Value: Criticality},
				Value:       msg.MaximumIPdatarate,
			})
		}
	}
	return ies, nil
}

// Encode function for SecurityIndication to be generated here.

// Decode function for SecurityIndication to be generated here.

// Decoder helper for SecurityIndication to be generated here.
