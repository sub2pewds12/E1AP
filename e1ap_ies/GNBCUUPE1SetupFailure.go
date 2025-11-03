package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPE1SetupFailure is a generated SEQUENCE type.
type GNBCUUPE1SetupFailure struct {
	TransactionID          TransactionID           `aper:"lb:0,ub:255,mandatory,ext"`
	Cause                  Cause                   `aper:"mandatory,ext"`
	TimeToWait             *TimeToWait             `aper:"optional,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ext"`
}

// toIes transforms the GNBCUUPE1SetupFailure struct into a slice of E1APMessageIEs.
func (msg *GNBCUUPE1SetupFailure) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 255},
					ext:   true,
					Value: msg.TransactionID.Value,
				},
			})
		}
	}
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDCause},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &msg.Cause,
			})
		}
	}
	if msg.TimeToWait != nil {

		{

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
	}
	if msg.CriticalityDiagnostics != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDCriticalityDiagnostics},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.CriticalityDiagnostics,
			})
		}
	}
	var err error
	return ies, err
}

// Encode for GNBCUUPE1SetupFailure: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for GNBCUUPE1SetupFailure.
func (msg *GNBCUUPE1SetupFailure) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GNBCUUPE1SetupFailure: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUUPE1SetupFailureDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEIDTransactionID]; !ok {
		err = fmt.Errorf("mandatory field TransactionID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDTransactionID,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDCause]; !ok {
		err = fmt.Errorf("mandatory field Cause is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDCause,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type GNBCUUPE1SetupFailureDecoder struct {
	msg      *GNBCUUPE1SetupFailure
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUUPE1SetupFailureDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDTransactionID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
				return nil, fmt.Errorf("Decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = aper.Integer(val)
		}
	case ProtocolIEIDCause:
		if err = msg.Cause.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode Cause failed: %w", err)
		}
	case ProtocolIEIDTimeToWait:
		msg.TimeToWait = new(TimeToWait)

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true); err != nil {
				return nil, fmt.Errorf("Decode TimeToWait failed: %w", err)
			}
			msg.TimeToWait.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDCriticalityDiagnostics:
		msg.CriticalityDiagnostics = new(CriticalityDiagnostics)
		if err = msg.CriticalityDiagnostics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode CriticalityDiagnostics failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
