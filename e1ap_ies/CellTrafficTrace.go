package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
)

// CellTrafficTrace is a generated SEQUENCE type.
type CellTrafficTrace struct {
	GNBCUCPUEE1APID                GNBCUCPUEE1APID
	GNBCUUPUEE1APID                GNBCUUPUEE1APID
	TraceID                        TraceID
	TraceCollectionEntityIPAddress TransportLayerAddress
	PrivacyIndicator               *PrivacyIndicator
	URIaddress                     *URIaddress
}

// toIes transforms the CellTrafficTrace struct into a slice of E1APMessageIEs.
func (msg *CellTrafficTrace) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GNBCUCPUEE1APID,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GNBCUUPUEE1APID,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDTraceID},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.TraceID,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDTraceCollectionEntityIPAddress},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.TraceCollectionEntityIPAddress,
	})
	if msg.PrivacyIndicator != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDPrivacyIndicator},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PrivacyIndicator,
		})
	}
	if msg.URIaddress != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDURIaddress},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.URIaddress,
		})
	}
	return ies, nil
}

func (msg *CellTrafficTrace) EncodeWithEncoder(e *per.Encoder) (err error) {
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

func (msg *CellTrafficTrace) Encode(w io.Writer) error {
	e := per.NewEncoder(per.APER)
	if err := msg.EncodeWithEncoder(e); err != nil {
		return err
	}
	_, err := w.Write(e.Bytes())
	return err
}

// Decode implements the MessageUnmarshaller interface for CellTrafficTrace.
func (msg *CellTrafficTrace) Decode(data []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	r := per.NewDecoder(data, per.APER)
	return msg.DecodeFromDecoder(r)
}

func (msg *CellTrafficTrace) DecodeFromDecoder(r *per.Decoder) (diagList []CriticalityDiagnosticsIEItem, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("decode CellTrafficTrace failed: %w", err)
		}
	}()

	decoder := CellTrafficTraceDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field GNBCUCPUEE1APID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field GNBCUUPUEE1APID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDTraceID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field TraceID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDTraceID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDTraceCollectionEntityIPAddress}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field TraceCollectionEntityIPAddress is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDTraceCollectionEntityIPAddress},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type CellTrafficTraceDecoder struct {
	msg      *CellTrafficTrace
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *CellTrafficTraceDecoder) decodeIE(r *per.Decoder) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			val, err := ieR.DecodeInteger(per.Unconstrained())
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPUEE1APID failed: %w", err)
			}
			msg.GNBCUCPUEE1APID.Value = val
		}
	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			val, err := ieR.DecodeInteger(per.Unconstrained())
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPUEE1APID failed: %w", err)
			}
			msg.GNBCUUPUEE1APID.Value = val
		}
	case ProtocolIEIDTraceID:

		if err = msg.TraceID.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode TraceID failed: %w", err)
		}
	case ProtocolIEIDTraceCollectionEntityIPAddress:

		if err = msg.TraceCollectionEntityIPAddress.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode TraceCollectionEntityIPAddress failed: %w", err)
		}
	case ProtocolIEIDPrivacyIndicator:
		msg.PrivacyIndicator = new(PrivacyIndicator)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode PrivacyIndicator failed: %w", err)
			}
			msg.PrivacyIndicator.Value = val
		}
	case ProtocolIEIDURIaddress:

		{
			val, err := ieR.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(0)})
			if err != nil {
				return nil, fmt.Errorf("decode URIaddress failed: %w", err)
			}
			msg.URIaddress = new(URIaddress)
			msg.URIaddress.Value = val
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
