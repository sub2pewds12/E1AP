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
				Id:          ProtocolIEID{Value: ProtocolIEIDRegistrationRequest},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 1},
					ext:   true,
					Value: msg.RegistrationRequest.Value,
				},
			})
		}
	}
	if msg.ReportCharacteristics != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDReportCharacteristics},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &BITSTRING{
					c:     aper.Constraint{Lb: 36, Ub: 36},
					ext:   false,
					Value: msg.ReportCharacteristics.Value,
				},
			})
		}
	}
	if msg.ReportingPeriodicity != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDReportingPeriodicity},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 15},
					ext:   true,
					Value: msg.ReportingPeriodicity.Value,
				},
			})
		}
	}
	var err error
	return ies, err
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
func (msg *ResourceStatusRequest) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("ResourceStatusRequest: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := ResourceStatusRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDRegistrationRequest}]; !ok {
		err = fmt.Errorf("mandatory field RegistrationRequest is missing")
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
			msg.GNBCUUPMeasurementID = new(ResourceStatusRequestIEsIDGNBCUUPMeasurementID)
			msg.GNBCUUPMeasurementID.Value = aper.Integer(val)
		}
	case ProtocolIEIDRegistrationRequest:

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return nil, fmt.Errorf("Decode RegistrationRequest failed: %w", err)
			}
			msg.RegistrationRequest.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDReportCharacteristics:
		msg.ReportCharacteristics = new(ReportCharacteristics)
		if err = msg.ReportCharacteristics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode ReportCharacteristics failed: %w", err)
		}
	case ProtocolIEIDReportingPeriodicity:
		msg.ReportingPeriodicity = new(ReportingPeriodicity)

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, true); err != nil {
				return nil, fmt.Errorf("Decode ReportingPeriodicity failed: %w", err)
			}
			msg.ReportingPeriodicity.Value = aper.Enumerated(val)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
