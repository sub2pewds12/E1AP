package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ResourceStatusRequest is a generated SEQUENCE type.
type ResourceStatusRequest struct {
	TransactionID         TransactionID                                   `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPMeasurementID  ResourceStatusRequestIEsIDGNBCUCPMeasurementID  `aper:"lb:1,ub:4095,mandatory,ext"`
	GNBCUUPMeasurementID  *ResourceStatusRequestIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,optional,ext"`
	RegistrationRequest   RegistrationRequest                             `aper:"mandatory,ext"`
	ReportCharacteristics *ReportCharacteristics                          `aper:"lb:36,ub:36,conditional,ext"`
	ReportingPeriodicity  *ReportingPeriodicity                           `aper:"optional,ext"`
}

// toIes transforms the ResourceStatusRequest struct into a slice of E1APMessageIEs.
func (msg *ResourceStatusRequest) toIes() ([]E1APMessageIE, error) {
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

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDRegistrationRequest},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.RegistrationRequest,
	})
	if msg.ReportCharacteristics != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDReportCharacteristics},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.ReportCharacteristics,
		})
	}
	if msg.ReportingPeriodicity != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDReportingPeriodicity},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.ReportingPeriodicity,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for ResourceStatusRequest.
func (msg *ResourceStatusRequest) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert ResourceStatusRequest to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodeResourceStatusReportingInitiation}, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for ResourceStatusRequest.
func (msg *ResourceStatusRequest) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode ResourceStatusRequest failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := ResourceStatusRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDRegistrationRequest}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field RegistrationRequest is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDRegistrationRequest},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type ResourceStatusRequestDecoder struct {
	msg      *ResourceStatusRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *ResourceStatusRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
			msg.GNBCUUPMeasurementID = new(ResourceStatusRequestIEsIDGNBCUUPMeasurementID)
			msg.GNBCUUPMeasurementID.Value = aper.Integer(val)
		}
	case ProtocolIEIDRegistrationRequest:

		{
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
			if err != nil {
				return nil, fmt.Errorf("decode RegistrationRequest failed: %w", err)
			}
			msg.RegistrationRequest.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDReportCharacteristics:
		msg.ReportCharacteristics = new(ReportCharacteristics)
		if err = msg.ReportCharacteristics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode ReportCharacteristics failed: %w", err)
		}
	case ProtocolIEIDReportingPeriodicity:
		msg.ReportingPeriodicity = new(ReportingPeriodicity)

		{
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, true)
			if err != nil {
				return nil, fmt.Errorf("decode ReportingPeriodicity failed: %w", err)
			}
			msg.ReportingPeriodicity.Value = aper.Enumerated(val)
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
