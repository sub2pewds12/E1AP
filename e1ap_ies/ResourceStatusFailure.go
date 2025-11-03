package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ResourceStatusFailure is a generated SEQUENCE type.
type ResourceStatusFailure struct {
	TransactionID          TransactionID                                   `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPMeasurementID   ResourceStatusFailureIEsIDGNBCUCPMeasurementID  `aper:"lb:1,ub:4095,mandatory,ext"`
	GNBCUUPMeasurementID   *ResourceStatusFailureIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,optional,ext"`
	Cause                  Cause                                           `aper:"mandatory,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics                         `aper:"optional,ext"`
}

// toIes transforms the ResourceStatusFailure struct into a slice of E1APMessageIEs.
func (msg *ResourceStatusFailure) toIes() ([]E1APMessageIE, error) {
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
	if msg.GNBCUUPMeasurementID != nil {

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
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDCause},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &msg.Cause,
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

// Encode implements the aper.AperMarshaller interface for ResourceStatusFailure.
func (msg *ResourceStatusFailure) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert ResourceStatusFailure to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduUnsuccessfulOutcome, ProcedureCodeResourceStatusReportingInitiation, Criticality{Value: CriticalityIgnore}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for ResourceStatusFailure.
func (msg *ResourceStatusFailure) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("ResourceStatusFailure: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := ResourceStatusFailureDecoder{
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

	if _, ok := decoder.list[ProtocolIEIDGNBCUCPMeasurementID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUCPMeasurementID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDGNBCUCPMeasurementID,
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

type ResourceStatusFailureDecoder struct {
	msg      *ResourceStatusFailure
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *ResourceStatusFailureDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
			msg.GNBCUUPMeasurementID = new(ResourceStatusFailureIEsIDGNBCUUPMeasurementID)
			msg.GNBCUUPMeasurementID.Value = aper.Integer(val)
		}
	case ProtocolIEIDCause:
		if err = msg.Cause.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode Cause failed: %w", err)
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
