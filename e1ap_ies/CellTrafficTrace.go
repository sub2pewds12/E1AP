package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// CellTrafficTrace is a generated SEQUENCE type.
type CellTrafficTrace struct {
	GNBCUCPUEE1APID                GNBCUCPUEE1APID       `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                GNBCUUPUEE1APID       `aper:"lb:0,ub:4294967295,mandatory,ext"`
	TraceID                        TraceID               `aper:"lb:8,ub:8,mandatory,ext"`
	TraceCollectionEntityIPAddress TransportLayerAddress `aper:"lb:1,ub:160,mandatory,ext"`
	PrivacyIndicator               *PrivacyIndicator     `aper:"optional,ext"`
	URIaddress                     *URIaddress           `aper:"optional,ext"`
}

// toIes transforms the CellTrafficTrace struct into a slice of E1APMessageIEs.
func (msg *CellTrafficTrace) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
		Criticality: Criticality{Value: CriticalityReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: msg.GNBCUCPUEE1APID.Value,
		},
	})

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
		Criticality: Criticality{Value: CriticalityReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: msg.GNBCUUPUEE1APID.Value,
		},
	})

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDTraceID},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 8, Ub: 8},
			ext:   false,
			Value: msg.TraceID.Value,
		},
	})

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDTraceCollectionEntityIPAddress},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value: &BITSTRING{
			c:     aper.Constraint{Lb: 1, Ub: 160},
			ext:   false,
			Value: msg.TraceCollectionEntityIPAddress.Value,
		},
	})
	if msg.PrivacyIndicator != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDPrivacyIndicator},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 1},
				ext:   true,
				Value: msg.PrivacyIndicator.Value,
			},
		})
	}
	if msg.URIaddress != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDURIaddress},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.URIaddress.Value,
			},
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for CellTrafficTrace.
func (msg *CellTrafficTrace) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert CellTrafficTrace to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodeCellTrafficTrace}, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for CellTrafficTrace.
func (msg *CellTrafficTrace) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode CellTrafficTrace failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := CellTrafficTraceDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf(decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
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

func (decoder *CellTrafficTraceDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false)
	if err != nil {
		return nil, err
	}
	msgIe = new(E1APMessageIE)
	msgIe.Id = ProtocolIEID{Value: aper.Integer(id)}
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return nil, err
	}
	msgIe.Criticality = Criticality{Value: aper.Enumerated(c)}

	buf, err := r.ReadOpenType()
	if err != nil {
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
	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPUEE1APID failed: %w", err)
			}
			msg.GNBCUCPUEE1APID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPUEE1APID failed: %w", err)
			}
			msg.GNBCUUPUEE1APID.Value = aper.Integer(val)
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
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
			if err != nil {
				return nil, fmt.Errorf("decode PrivacyIndicator failed: %w", err)
			}
			msg.PrivacyIndicator.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDURIaddress:
		msg.URIaddress = new(URIaddress)
		if err = msg.URIaddress.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode URIaddress failed: %w", err)
		}
	default:
		switch msgIe.Criticality.Value {
		case CriticalityReject:
			// If an unknown IE is critical, the PDU cannot be processed.
			return nil, fmt.Errorf("not comprehended IE ID %d (criticality: reject)", msgIe.Id.Value)
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
