package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ResourceStatusResponse is a generated SEQUENCE type.
type ResourceStatusResponse struct {
	TransactionID          TransactionID                                   `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPMeasurementID   ResourceStatusResponseIEsIDGNBCUCPMeasurementID `aper:"lb:1,ub:4095,mandatory,ext"`
	GNBCUUPMeasurementID   ResourceStatusResponseIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,mandatory,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics                         `aper:"optional,ext"`
}

// toIes transforms the ResourceStatusResponse struct into a slice of E1APMessageIEs.
func (msg *ResourceStatusResponse) toIes() ([]E1APMessageIE, error) {
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
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPMeasurementID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 1, Ub: 4095},
					ext:   true,
					Value: msg.GNBCUCPMeasurementID.Value,
				},
			})
		}
	}
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPMeasurementID},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 1, Ub: 4095},
					ext:   true,
					Value: msg.GNBCUUPMeasurementID.Value,
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

// Encode implements the aper.AperMarshaller interface for ResourceStatusResponse.
func (msg *ResourceStatusResponse) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert ResourceStatusResponse to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduSuccessfulOutcome, ProcedureCode{Value: ProcedureCodeResourceStatusReportingInitiation}, Criticality{Value: CriticalityIgnore}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for ResourceStatusResponse.
func (msg *ResourceStatusResponse) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("ResourceStatusResponse: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := ResourceStatusResponseDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDTransactionID}]; !ok {
		err = fmt.Errorf("mandatory field TransactionID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUCPMeasurementID}]; !ok {
		err = fmt.Errorf("mandatory field GNBCUCPMeasurementID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUUPMeasurementID}]; !ok {
		err = fmt.Errorf("mandatory field GNBCUUPMeasurementID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type ResourceStatusResponseDecoder struct {
	msg      *ResourceStatusResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *ResourceStatusResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
		return nil, fmt.Errorf("duplicated protocol IE ID %d", ieId.Value)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id.Value {
	case ProtocolIEIDTransactionID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
				return nil, fmt.Errorf("Decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUCPMeasurementID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true); err != nil {
				return nil, fmt.Errorf("Decode GNBCUCPMeasurementID failed: %w", err)
			}
			msg.GNBCUCPMeasurementID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPMeasurementID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true); err != nil {
				return nil, fmt.Errorf("Decode GNBCUUPMeasurementID failed: %w", err)
			}
			msg.GNBCUUPMeasurementID.Value = aper.Integer(val)
		}
	case ProtocolIEIDCriticalityDiagnostics:
		msg.CriticalityDiagnostics = new(CriticalityDiagnostics)
		if err = msg.CriticalityDiagnostics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode CriticalityDiagnostics failed: %w", err)
		}
	default:
		switch msgIe.Criticality.Value {
		case CriticalityReject:
			// If an unknown IE is critical, the PDU cannot be processed.
			err = fmt.Errorf("not comprehended IE ID %d (criticality: reject)", msgIe.Id.Value)
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.Id,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotUnderstood},
			})
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
