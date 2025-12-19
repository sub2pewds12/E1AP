package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// EarlyForwardingSNTransfer is a generated SEQUENCE type.
type EarlyForwardingSNTransfer struct {
	GNBCUCPUEE1APID                  GNBCUCPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                  GNBCUUPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,ext"`
	DRBsSubjectToEarlyForwardingList DRBsSubjectToEarlyForwardingList `aper:"lb:1,ub:MaxnoofDRBs,mandatory,ext"`
}

// toIes transforms the EarlyForwardingSNTransfer struct into a slice of E1APMessageIEs.
func (msg *EarlyForwardingSNTransfer) toIes() ([]E1APMessageIE, error) {
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
	{

		tmpDRBsSubjectToEarlyForwardingList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 1, Ub: MaxnoofDRBs},
			ext: false,
		}

		for i := 0; i < len(msg.DRBsSubjectToEarlyForwardingList.Value); i++ {
			tmpDRBsSubjectToEarlyForwardingList.Value = append(tmpDRBsSubjectToEarlyForwardingList.Value, &msg.DRBsSubjectToEarlyForwardingList.Value[i])
		}

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDDRBsSubjectToEarlyForwardingList},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &tmpDRBsSubjectToEarlyForwardingList,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for EarlyForwardingSNTransfer.
func (msg *EarlyForwardingSNTransfer) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert EarlyForwardingSNTransfer to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodeEarlyForwardingSNTransfer}, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for EarlyForwardingSNTransfer.
func (msg *EarlyForwardingSNTransfer) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode EarlyForwardingSNTransfer failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := EarlyForwardingSNTransferDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDDRBsSubjectToEarlyForwardingList}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field DRBsSubjectToEarlyForwardingList is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDDRBsSubjectToEarlyForwardingList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type EarlyForwardingSNTransferDecoder struct {
	msg      *EarlyForwardingSNTransfer
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *EarlyForwardingSNTransferDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDDRBsSubjectToEarlyForwardingList:

		{
			itemDecoder := func(r *aper.AperReader) (*DRBsSubjectToEarlyForwardingItem, error) {

				item := new(DRBsSubjectToEarlyForwardingItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 1, Ub: MaxnoofDRBs}, false)
			if err != nil {
				return nil, fmt.Errorf("decode DRBsSubjectToEarlyForwardingList failed: %w", err)
			}
			msg.DRBsSubjectToEarlyForwardingList.Value = decodedItems
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
