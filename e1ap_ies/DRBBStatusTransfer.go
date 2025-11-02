package e1ap_ies

import (
	"bytes"
	"fmt"

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

// Encode for DRBBStatusTransfer: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for DRBBStatusTransfer.
func (msg *DRBBStatusTransfer) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("DRBBStatusTransfer: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := DRBBStatusTransferDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID]; !ok {
		err = fmt.Errorf("mandatory field CountValue is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type DRBBStatusTransferDecoder struct {
	msg      *DRBBStatusTransfer
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*E1APMessageIE
}

func (decoder *DRBBStatusTransferDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}
	msgIe = new(E1APMessageIE)
	msgIe.Id.Value = aper.Integer(id)

	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)

	if buf, err = r.ReadOpenType(); err != nil {
		return
	}

	ieId := msgIe.Id.Value
	if _, ok := decoder.list[ieId]; ok {
		err = fmt.Errorf("duplicated protocol IE ID %%d", ieId)
		return
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id.Value {

	case ProtocolIEID:

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 131072}, false); err != nil {
				return fmt.Errorf("Decode ReceiveStatusofPDCPSDU failed: %w", err)
			}
			tmp := DRBBStatusTransferReceiveStatusofPDCPSDU(val)
			s.ReceiveStatusofPDCPSDU = &tmp
		}

	case ProtocolIEID:
		if err = s.CountValue.Decode(r); err != nil {
			return fmt.Errorf("Decode CountValue failed: %w", err)
		}

	case ProtocolIEID:
		s.IEExtension = new(ProtocolExtensionContainer)
		if err = s.IEExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtension failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
	}
	return
}
