package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ResourceStatusUpdate is a generated SEQUENCE type.
type ResourceStatusUpdate struct {
	TransactionID                 TransactionID                                  `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPMeasurementID          ResourceStatusUpdateIEsIDGNBCUCPMeasurementID  `aper:"lb:1,ub:4095,mandatory,ext"`
	GNBCUUPMeasurementID          *ResourceStatusUpdateIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,optional,ext"`
	TNLAvailableCapacityIndicator *TNLAvailableCapacityIndicator                 `aper:"optional,ext"`
	HWCapacityIndicator           HWCapacityIndicator                            `aper:"mandatory,ext"`
}

// toIes transforms the ResourceStatusUpdate struct into a slice of E1APMessageIEs.
func (msg *ResourceStatusUpdate) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TransactionID,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPMeasurementID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GNBCUCPMeasurementID,
	})
	if msg.GNBCUUPMeasurementID != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPMeasurementID},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUUPMeasurementID,
		})
	}
	if msg.TNLAvailableCapacityIndicator != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDTNLAvailableCapacityIndicator},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TNLAvailableCapacityIndicator,
		})
	}

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDHWCapacityIndicator},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.HWCapacityIndicator,
	})
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for ResourceStatusUpdate.
func (msg *ResourceStatusUpdate) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert ResourceStatusUpdate to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodeResourceStatusReporting}, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for ResourceStatusUpdate.
func (msg *ResourceStatusUpdate) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode ResourceStatusUpdate failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := ResourceStatusUpdateDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUCPMeasurementID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field GNBCUCPMeasurementID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDHWCapacityIndicator}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field HWCapacityIndicator is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDHWCapacityIndicator},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type ResourceStatusUpdateDecoder struct {
	msg      *ResourceStatusUpdate
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *ResourceStatusUpdateDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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

	switch msgIe.ID.Value {
	case ProtocolIEIDTransactionID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true)
			if err != nil {
				return nil, fmt.Errorf("decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUCPMeasurementID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPMeasurementID failed: %w", err)
			}
			msg.GNBCUCPMeasurementID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPMeasurementID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPMeasurementID failed: %w", err)
			}
			msg.GNBCUUPMeasurementID = new(ResourceStatusUpdateIEsIDGNBCUUPMeasurementID)
			msg.GNBCUUPMeasurementID.Value = aper.Integer(val)
		}
	case ProtocolIEIDTNLAvailableCapacityIndicator:
		msg.TNLAvailableCapacityIndicator = new(TNLAvailableCapacityIndicator)
		if err = msg.TNLAvailableCapacityIndicator.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode TNLAvailableCapacityIndicator failed: %w", err)
		}
	case ProtocolIEIDHWCapacityIndicator:
		if err = msg.HWCapacityIndicator.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode HWCapacityIndicator failed: %w", err)
		}
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
