package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// PrivateMessage is a generated SEQUENCE type.
type PrivateMessage struct {
	PrivateIEs PrivateIEContainer `aper:"ub:MaxPrivateIEs,mandatory,ext"`
}

// toIes transforms the PrivateMessage struct into a slice of E1APMessageIEs.
func (msg *PrivateMessage) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for PrivateMessage.
func (msg *PrivateMessage) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert PrivateMessage to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodePrivateMessage}, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for PrivateMessage.
func (msg *PrivateMessage) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode PrivateMessage failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := PrivateMessageDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf(decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	return
}

type PrivateMessageDecoder struct {
	msg      *PrivateMessage
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *PrivateMessageDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false)
	if err != nil {
		return nil, err
	}
	msgIe = new(E1APMessageIE)
	msgIe.ID = ProtocolIEID{Value: aper.Integer(id)}
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return nil, err
	}
	msgIe.Criticality = Criticality{Value: aper.Enumerated(c)}

	buf, err := r.ReadOpenType()
	if err != nil {
		return nil, err
	}

	ieId := msgIe.ID
	if _, ok := decoder.list[ieId]; ok {
		return nil, fmt.Errorf("duplicated protocol IE ID %d", ieId.Value)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	_ = ieR
	_ = msg
	switch msgIe.ID.Value {
	default:
		switch msgIe.Criticality.Value {
		case CriticalityReject:
			// If an unknown IE is critical, the PDU cannot be processed.
			return nil, fmt.Errorf("not comprehended IE ID %d (criticality: reject)", msgIe.ID.Value)
		case CriticalityNotify:
			// Per 3GPP TS 38.463 Section 10.3, report and proceed.
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.ID,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotUnderstood},
			})
		case CriticalityIgnore:
			// Ignore and proceed.
		}
	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
