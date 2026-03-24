package e1ap_ies

import (
	"fmt"
	"io"

	"asn1go/per"
)

// ResourceStatusRequest is a generated SEQUENCE type.
type ResourceStatusRequest struct {
	TransactionID         TransactionID
	GNBCUCPMeasurementID  ResourceStatusRequestIEsIDGNBCUCPMeasurementID
	GNBCUUPMeasurementID  *ResourceStatusRequestIEsIDGNBCUUPMeasurementID
	RegistrationRequest   RegistrationRequest
	ReportCharacteristics *ReportCharacteristics
	ReportingPeriodicity  *ReportingPeriodicity
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

func (msg *ResourceStatusRequest) EncodeWithEncoder(e *per.Encoder) (err error) {
	ies, err := msg.toIes()
	if err != nil {
		return err
	}

	sizeC := per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(65535)}
	if err = e.EncodeLengthDeterminant(int64(len(ies)), sizeC); err != nil {
		return fmt.Errorf("encode IE count failed: %w", err)
	}
	for i := range ies {
		if err = ies[i].Encode(e); err != nil {
			return fmt.Errorf("encode IE %d failed: %w", i, err)
		}
	}
	return nil
}

func (msg *ResourceStatusRequest) Encode(w io.Writer) error {
	e := per.NewEncoder(per.APER)
	if err := msg.EncodeWithEncoder(e); err != nil {
		return err
	}
	_, err := w.Write(e.Bytes())
	return err
}

// Decode implements the MessageUnmarshaller interface for ResourceStatusRequest.
func (msg *ResourceStatusRequest) Decode(data []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	r := per.NewDecoder(data, per.APER)
	return msg.DecodeFromDecoder(r)
}

func (msg *ResourceStatusRequest) DecodeFromDecoder(r *per.Decoder) (diagList []CriticalityDiagnosticsIEItem, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("decode ResourceStatusRequest failed: %w", err)
		}
	}()

	decoder := ResourceStatusRequestDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	c := per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(65535)}
	length, err := r.DecodeLengthDeterminant(c)
	if err != nil {
		return
	}

	for i := int64(0); i < length; i++ {
		if _, err = decoder.decodeIE(r); err != nil {
			return
		}
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

func (decoder *ResourceStatusRequestDecoder) decodeIE(r *per.Decoder) (msgIe *E1APMessageIE, err error) {
	id, err := r.DecodeInteger(per.Constrained(0, 65535))
	if err != nil {
		return nil, err
	}
	msgIe = new(E1APMessageIE)
	msgIe.ID = ProtocolIEID{Value: id}

	enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3)}
	c, err := r.DecodeEnumerated(enumC)
	if err != nil {
		return nil, err
	}
	msgIe.Criticality = Criticality{Value: c}

	buf, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: nil, Max: nil})
	if err != nil {
		return nil, err
	}

	ieId := msgIe.ID
	if _, ok := decoder.list[ieId]; ok {
		return nil, fmt.Errorf("duplicated protocol IE ID %d", ieId.Value)
	}
	decoder.list[ieId] = msgIe

	ieR := per.NewDecoder(buf, per.APER)
	msg := decoder.msg

	switch msgIe.ID.Value {
	case ProtocolIEIDTransactionID:

		{
			val, err := ieR.DecodeInteger(per.ConstrainedExtensible(0, 255))
			if err != nil {
				return nil, fmt.Errorf("decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = val
		}
	case ProtocolIEIDGNBCUCPMeasurementID:

		{
			val, err := ieR.DecodeInteger(per.ConstrainedExtensible(1, 4095))
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPMeasurementID failed: %w", err)
			}
			msg.GNBCUCPMeasurementID.Value = val
		}
	case ProtocolIEIDGNBCUUPMeasurementID:

		{
			val, err := ieR.DecodeInteger(per.ConstrainedExtensible(1, 4095))
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPMeasurementID failed: %w", err)
			}
			msg.GNBCUUPMeasurementID = new(ResourceStatusRequestIEsIDGNBCUUPMeasurementID)
			msg.GNBCUUPMeasurementID.Value = val
		}
	case ProtocolIEIDRegistrationRequest:

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode RegistrationRequest failed: %w", err)
			}
			msg.RegistrationRequest.Value = val
		}
	case ProtocolIEIDReportCharacteristics:
		msg.ReportCharacteristics = new(ReportCharacteristics)

		if err = msg.ReportCharacteristics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode ReportCharacteristics failed: %w", err)
		}
	case ProtocolIEIDReportingPeriodicity:
		msg.ReportingPeriodicity = new(ReportingPeriodicity)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 16)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode ReportingPeriodicity failed: %w", err)
			}
			msg.ReportingPeriodicity.Value = val
		}
	default:
		switch msgIe.Criticality.Value {
		case CriticalityReject:
			return nil, fmt.Errorf("not comprehended IE ID %d (criticality: reject)", msgIe.ID.Value)
		case CriticalityNotify:
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.ID,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotUnderstood},
			})
		case CriticalityIgnore:
		}
	}
	return msgIe, nil
}
