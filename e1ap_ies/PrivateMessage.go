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
	{

		tmp_PrivateIEs := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: MaxPrivateIEs},
			ext: false,
		}

		for i := 0; i < len(msg.PrivateIEs.Value); i++ {
			tmp_PrivateIEs.Value = append(tmp_PrivateIEs.Value, &msg.PrivateIEs.Value[i])
		}

		{

			tmp_PrivateIEs := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxPrivateIEs},
				ext: false,
			}

			for i := 0; i < len(msg.PrivateIEs.Value); i++ {
				tmp_PrivateIEs.Value = append(tmp_PrivateIEs.Value, &msg.PrivateIEs.Value[i])
			}

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEID},
				Criticality: Criticality{Value: Criticality},
				Value:       &tmp_PrivateIEs,
			})
		}
	}
	var err error
	return ies, err
}

// Encode implements the aper.AperMarshaller interface for PrivateMessage.
func (msg *PrivateMessage) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert PrivateMessage to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCodePrivateMessage, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for PrivateMessage.
func (msg *PrivateMessage) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("PrivateMessage: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := PrivateMessageDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID]; !ok {
		err = fmt.Errorf("mandatory field PrivateIEs is missing")
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

type PrivateMessageDecoder struct {
	msg      *PrivateMessage
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *PrivateMessageDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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

		{
			itemDecoder := func(r *aper.AperReader) (*PrivateIEField, error) {

				item := new(PrivateIEField)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			var decodedItems []PrivateIEField
			if decodedItems, err = aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxPrivateIEs}, false); err != nil {
				return nil, fmt.Errorf("Decode PrivateIEs failed: %w", err)
			}
			msg.PrivateIEs.Value = decodedItems
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
