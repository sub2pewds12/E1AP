package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// IABUPTNLAddressUpdateAcknowledge is a generated SEQUENCE type.
type IABUPTNLAddressUpdateAcknowledge struct {
	TransactionID              TransactionID               `aper:"lb:0,ub:255,mandatory,ext"`
	CriticalityDiagnostics     *CriticalityDiagnostics     `aper:"optional,ext"`
	ULUPTNLAddressToUpdateList *ULUPTNLAddressToUpdateList `aper:"ub:MaxnoofTNLAddresses,optional,ext"`
}

// toIes transforms the IABUPTNLAddressUpdateAcknowledge struct into a slice of E1APMessageIEs.
func (msg *IABUPTNLAddressUpdateAcknowledge) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
		Criticality: Criticality{Value: CriticalityReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   true,
			Value: msg.TransactionID.Value,
		},
	})
	if msg.CriticalityDiagnostics != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDCriticalityDiagnostics},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.ULUPTNLAddressToUpdateList != nil {

		tmpULUPTNLAddressToUpdateList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAddresses},
			ext: false,
		}

		for i := 0; i < len(msg.ULUPTNLAddressToUpdateList.Value); i++ {
			tmpULUPTNLAddressToUpdateList.Value = append(tmpULUPTNLAddressToUpdateList.Value, &msg.ULUPTNLAddressToUpdateList.Value[i])
		}

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDULUPTNLAddressToUpdateList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &tmpULUPTNLAddressToUpdateList,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for IABUPTNLAddressUpdateAcknowledge.
func (msg *IABUPTNLAddressUpdateAcknowledge) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert IABUPTNLAddressUpdateAcknowledge to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduSuccessfulOutcome, ProcedureCode{Value: ProcedureCodeIABUPTNLAddressUpdate}, Criticality{Value: CriticalityIgnore}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for IABUPTNLAddressUpdateAcknowledge.
func (msg *IABUPTNLAddressUpdateAcknowledge) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode IABUPTNLAddressUpdateAcknowledge failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := IABUPTNLAddressUpdateAcknowledgeDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf(decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDTransactionID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field TransactionID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type IABUPTNLAddressUpdateAcknowledgeDecoder struct {
	msg      *IABUPTNLAddressUpdateAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *IABUPTNLAddressUpdateAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false)
	if err != nil {
		return nil, err
	}
	msgIe = new(E1APMessageIE)
	msgIe.Id = ProtocolIEID{Value: aper.Integer(id)}
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return nil, err
	}
	msgIe.Criticality = Criticality{Value: aper.Enumerated(c)}

	buf, err := r.ReadOpenType()
	if err != nil {
		return nil, err
	}

	ieId := msgIe.Id
	if _, ok := decoder.list[ieId]; ok {
		return nil, fmt.Errorf("duplicated protocol IE ID %d", ieId.Value)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id.Value {
	case ProtocolIEIDTransactionID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true)
			if err != nil {
				return nil, fmt.Errorf("decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = aper.Integer(val)
		}
	case ProtocolIEIDCriticalityDiagnostics:
		msg.CriticalityDiagnostics = new(CriticalityDiagnostics)
		if err = msg.CriticalityDiagnostics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode CriticalityDiagnostics failed: %w", err)
		}
	case ProtocolIEIDULUPTNLAddressToUpdateList:

		{
			itemDecoder := func(r *aper.AperReader) (*ULUPTNLAddressToUpdateItem, error) {

				item := new(ULUPTNLAddressToUpdateItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAddresses}, false)
			if err != nil {
				return nil, fmt.Errorf("decode ULUPTNLAddressToUpdateList failed: %w", err)
			}

			msg.ULUPTNLAddressToUpdateList = new(ULUPTNLAddressToUpdateList)
			msg.ULUPTNLAddressToUpdateList.Value = decodedItems
		}
	default:
		switch msgIe.Criticality.Value {
		case CriticalityReject:
			// If an unknown IE is critical, the PDU cannot be processed.
			return nil, fmt.Errorf("not comprehended IE ID %d (criticality: reject)", msgIe.Id.Value)
		case CriticalityNotify:
			// Per 3GPP TS 38.463 Section 10.3, report and proceed.
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.Id,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotUnderstood},
			})
		case CriticalityIgnore:
			// Ignore and proceed.
		}
	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
