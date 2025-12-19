package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// IABUPTNLAddressUpdateFailure is a generated SEQUENCE type.
type IABUPTNLAddressUpdateFailure struct {
	TransactionID          TransactionID           `aper:"lb:0,ub:255,mandatory,ext"`
	Cause                  Cause                   `aper:"mandatory,ext"`
	TimeToWait             *TimeToWait             `aper:"optional,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ext"`
}

// toIes transforms the IABUPTNLAddressUpdateFailure struct into a slice of E1APMessageIEs.
func (msg *IABUPTNLAddressUpdateFailure) toIes() ([]E1APMessageIE, error) {
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

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDCause},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if msg.TimeToWait != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDTimeToWait},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 5},
				ext:   true,
				Value: msg.TimeToWait.Value,
			},
		})
	}
	if msg.CriticalityDiagnostics != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDCriticalityDiagnostics},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for IABUPTNLAddressUpdateFailure.
func (msg *IABUPTNLAddressUpdateFailure) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert IABUPTNLAddressUpdateFailure to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduUnsuccessfulOutcome, ProcedureCode{Value: ProcedureCodeIABUPTNLAddressUpdate}, Criticality{Value: CriticalityIgnore}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for IABUPTNLAddressUpdateFailure.
func (msg *IABUPTNLAddressUpdateFailure) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode IABUPTNLAddressUpdateFailure failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := IABUPTNLAddressUpdateFailureDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDCause}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field Cause is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDCause},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type IABUPTNLAddressUpdateFailureDecoder struct {
	msg      *IABUPTNLAddressUpdateFailure
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *IABUPTNLAddressUpdateFailureDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDCause:
		if err = msg.Cause.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode Cause failed: %w", err)
		}
	case ProtocolIEIDTimeToWait:
		msg.TimeToWait = new(TimeToWait)

		{
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
			if err != nil {
				return nil, fmt.Errorf("decode TimeToWait failed: %w", err)
			}
			msg.TimeToWait.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDCriticalityDiagnostics:
		msg.CriticalityDiagnostics = new(CriticalityDiagnostics)
		if err = msg.CriticalityDiagnostics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode CriticalityDiagnostics failed: %w", err)
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
