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

func (msg *ResourceStatusRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDTransactionID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 255},
				ext:   true,
				Value: aper.Integer(msg.TransactionID),
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDGNBCUCPMeasurementID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 4095},
				ext:   true,
				Value: aper.Integer(msg.GNBCUCPMeasurementID),
			},
		})
	}
	if msg.GNBCUUPMeasurementID != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUUPMeasurementID,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 1, Ub: 4095},
					ext:   true,
					Value: aper.Integer((*msg.GNBCUUPMeasurementID)),
				},
			})
		}
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDRegistrationRequest,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 1},
				ext:   true,
				Value: msg.RegistrationRequest.Value,
			},
		})
	}
	if msg.ReportCharacteristics != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDReportCharacteristics,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 36, Ub: 36},
					ext:   false,
					Value: aper.OctetString((*msg.ReportCharacteristics)),
				},
			})
		}
	}
	if msg.ReportingPeriodicity != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDReportingPeriodicity,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 15},
					ext:   true,
					Value: (*msg.ReportingPeriodicity).Value,
				},
			})
		}
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for ResourceStatusRequest.
func (msg *ResourceStatusRequest) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert ResourceStatusRequest to IEs: %w", err)
	}

	return EncodeInitiatingMessage(w, ProcedureCodeResourceStatusReportingInitiation, Criticality{Value: CriticalityReject}, ies)
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
		list: make(map[aper.Integer]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEIDTransactionID]; !ok {
		err = fmt.Errorf("mandatory field TransactionID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDGNBCUCPMeasurementID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUCPMeasurementID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDRegistrationRequest]; !ok {
		err = fmt.Errorf("mandatory field RegistrationRequest is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
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
	list     map[aper.Integer]*E1APMessageIE
}

func (decoder *ResourceStatusRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}
	msgIe = new(E1APMessageIE)
	msgIe.Id.Value = aper.Integer(id)

	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)

	if buf, err = r.ReadOpenType(); err != nil {
		return
	}

	ieId := msgIe.Id.Value
	if _, ok := decoder.list[ieId]; ok {
		err = fmt.Errorf("duplicated protocol IE ID %%d", ieId)
		return
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id.Value {

	case ProtocolIEIDTransactionID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
				return fmt.Errorf("Decode TransactionID failed: %w", err)
			}
			s.TransactionID = TransactionID(val)
		}

	case ProtocolIEIDGNBCUCPMeasurementID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true); err != nil {
				return fmt.Errorf("Decode GNBCUCPMeasurementID failed: %w", err)
			}
			s.GNBCUCPMeasurementID = ResourceStatusRequestIEsIDGNBCUCPMeasurementID(val)
		}

	case ProtocolIEIDGNBCUUPMeasurementID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true); err != nil {
				return fmt.Errorf("Decode GNBCUUPMeasurementID failed: %w", err)
			}
			tmp := ResourceStatusRequestIEsIDGNBCUUPMeasurementID(val)
			s.GNBCUUPMeasurementID = &tmp
		}

	case ProtocolIEIDRegistrationRequest:

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return fmt.Errorf("Decode RegistrationRequest failed: %w", err)
			}
			s.RegistrationRequest.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDReportCharacteristics:

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 36, Ub: 36}, false); err != nil {
				return fmt.Errorf("Decode ReportCharacteristics failed: %w", err)
			}
			tmp := ReportCharacteristics(val)
			s.ReportCharacteristics = &tmp
		}

	case ProtocolIEIDReportingPeriodicity:
		s.ReportingPeriodicity = new(ReportingPeriodicity)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, true); err != nil {
				return fmt.Errorf("Decode ReportingPeriodicity failed: %w", err)
			}
			s.ReportingPeriodicity.Value = aper.Enumerated(val)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
	}
	return
}
