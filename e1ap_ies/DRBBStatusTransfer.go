package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DRBBStatusTransfer is a generated SEQUENCE type.
type DRBBStatusTransfer struct {
	ReceiveStatusofPDCPSDU *DRBBStatusTransferReceiveStatusofPDCPSDU `aper:"lb:1,ub:131072,optional,ext"`
	CountValue             PDCPCount                                 `aper:"mandatory,ext"`
	IEExtension            *ProtocolExtensionContainer               `aper:"optional,ext"`
}

func (msg *DRBBStatusTransfer) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	if msg.ReceiveStatusofPDCPSDU != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID,
				Criticality: Criticality{Value: Criticality},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 1, Ub: 131072},
					ext:   false,
					Value: aper.OctetString((*msg.ReceiveStatusofPDCPSDU)),
				},
			})
		}
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID,
			Criticality: Criticality{Value: Criticality},
			Value:       &msg.CountValue,
		})
	}
	return ies, nil
}

// Encode function for DRBBStatusTransfer to be generated here.

// Decode function for DRBBStatusTransfer to be generated here.

// Decoder helper for DRBBStatusTransfer to be generated here.
