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

// toIes transforms the DRBBStatusTransfer struct into a slice of E1APMessageIEs.
func (msg *DRBBStatusTransfer) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	if msg.ReceiveStatusofPDCPSDU != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEID},
				Criticality: Criticality{Value: Criticality},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 1, Ub: 131072},
					ext:   false,
					Value: msg.ReceiveStatusofPDCPSDU.Value,
				},
			})
		}
	}
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEID},
				Criticality: Criticality{Value: Criticality},
				Value:       &msg.CountValue,
			})
		}
	}
	var err error
	return ies, err
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
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID]; !ok {
		err = fmt.Errorf("mandatory field CountValue is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID,
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
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *DRBBStatusTransferDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return nil, err
	}
	msgIe = new(E1APMessageIE)
	msgIe.Id = ProtocolIEID{Value: aper.Integer(id)}
	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return nil, err
	}
	msgIe.Criticality = Criticality{Value: aper.Enumerated(c)}

	if buf, err = r.ReadOpenType(); err != nil {
		return nil, err
	}

	ieId := msgIe.Id
	if _, ok := decoder.list[ieId]; ok {
		return nil, fmt.Errorf("duplicated protocol IE ID %%d", ieId)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id {
	case ProtocolIEID:
		msg.ReceiveStatusofPDCPSDU = new(DRBBStatusTransferReceiveStatusofPDCPSDU)
		if err = msg.ReceiveStatusofPDCPSDU.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode ReceiveStatusofPDCPSDU failed: %w", err)
		}
	case ProtocolIEID:
		if err = msg.CountValue.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode CountValue failed: %w", err)
		}
	case ProtocolIEID:
		msg.IEExtension = new(ProtocolExtensionContainer)
		if err = msg.IEExtension.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode IEExtension failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
