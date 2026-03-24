package e1ap_ies

import (
	"fmt"
	"io"

	"asn1go/per"
)

// GNBCUUPE1SetupRequest is a generated SEQUENCE type.
type GNBCUUPE1SetupRequest struct {
	TransactionID             TransactionID
	GNBCUUPID                 GNBCUUPID
	GNBCUUPName               *GNBCUUPName
	CNSupport                 CNSupport
	SupportedPLMNs            SupportedPLMNsList
	GNBCUUPCapacity           *GNBCUUPCapacity
	TransportLayerAddressInfo *TransportLayerAddressInfo
	ExtendedGNBCUUPName       *ExtendedGNBCUUPName
}

// toIes transforms the GNBCUUPE1SetupRequest struct into a slice of E1APMessageIEs.
func (msg *GNBCUUPE1SetupRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TransactionID,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GNBCUUPID,
	})
	if msg.GNBCUUPName != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPName},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUUPName,
		})
	}

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDCNSupport},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.CNSupport,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDSupportedPLMNs},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SupportedPLMNs,
	})
	if msg.GNBCUUPCapacity != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPCapacity},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUUPCapacity,
		})
	}
	if msg.TransportLayerAddressInfo != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDTransportLayerAddressInfo},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TransportLayerAddressInfo,
		})
	}
	if msg.ExtendedGNBCUUPName != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDExtendedGNBCUUPName},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ExtendedGNBCUUPName,
		})
	}
	return ies, nil
}

func (msg *GNBCUUPE1SetupRequest) EncodeWithEncoder(e *per.Encoder) (err error) {
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

func (msg *GNBCUUPE1SetupRequest) Encode(w io.Writer) error {
	e := per.NewEncoder(per.APER)
	if err := msg.EncodeWithEncoder(e); err != nil {
		return err
	}
	_, err := w.Write(e.Bytes())
	return err
}

// Decode implements the MessageUnmarshaller interface for GNBCUUPE1SetupRequest.
func (msg *GNBCUUPE1SetupRequest) Decode(data []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	r := per.NewDecoder(data, per.APER)
	return msg.DecodeFromDecoder(r)
}

func (msg *GNBCUUPE1SetupRequest) DecodeFromDecoder(r *per.Decoder) (diagList []CriticalityDiagnosticsIEItem, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("decode GNBCUUPE1SetupRequest failed: %w", err)
		}
	}()

	decoder := GNBCUUPE1SetupRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUUPID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field GNBCUUPID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDCNSupport}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field CNSupport is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDCNSupport},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDSupportedPLMNs}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field SupportedPLMNs is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDSupportedPLMNs},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type GNBCUUPE1SetupRequestDecoder struct {
	msg      *GNBCUUPE1SetupRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUUPE1SetupRequestDecoder) decodeIE(r *per.Decoder) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDGNBCUUPID:

		{
			val, err := ieR.DecodeInteger(per.Constrained(0, 68719476735))
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPID failed: %w", err)
			}
			msg.GNBCUUPID.Value = val
		}
	case ProtocolIEIDGNBCUUPName:

		{
			val, err := ieR.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPName failed: %w", err)
			}
			msg.GNBCUUPName = new(GNBCUUPName)
			msg.GNBCUUPName.Value = val
		}
	case ProtocolIEIDCNSupport:

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode CNSupport failed: %w", err)
			}
			msg.CNSupport.Value = val
		}
	case ProtocolIEIDSupportedPLMNs:

		{
			itemDecoder := func(r *per.Decoder) (*SupportedPLMNsItem, error) {
				item := new(SupportedPLMNsItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}

			c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxnoofSPLMNs)}
			length, err := ieR.DecodeLengthDeterminant(c)
			if err != nil {
				return nil, fmt.Errorf("decode struct list length failed: %w", err)
			}
			for i := int64(0); i < length; i++ {
				item, err := itemDecoder(ieR)
				if err != nil {
					return nil, fmt.Errorf("decode item failed: %w", err)
				}
				msg.SupportedPLMNs.Value = append(msg.SupportedPLMNs.Value, *item)
			}
		}
	case ProtocolIEIDGNBCUUPCapacity:

		{
			val, err := ieR.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPCapacity failed: %w", err)
			}
			msg.GNBCUUPCapacity = new(GNBCUUPCapacity)
			msg.GNBCUUPCapacity.Value = val
		}
	case ProtocolIEIDTransportLayerAddressInfo:
		msg.TransportLayerAddressInfo = new(TransportLayerAddressInfo)

		if err = msg.TransportLayerAddressInfo.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode TransportLayerAddressInfo failed: %w", err)
		}
	case ProtocolIEIDExtendedGNBCUUPName:
		msg.ExtendedGNBCUUPName = new(ExtendedGNBCUUPName)

		if err = msg.ExtendedGNBCUUPName.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode ExtendedGNBCUUPName failed: %w", err)
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
