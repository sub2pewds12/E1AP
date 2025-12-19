package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextModificationConfirm is a generated SEQUENCE type.
type BearerContextModificationConfirm struct {
	GNBCUCPUEE1APID                        GNBCUCPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                        GNBCUUPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	SystemBearerContextModificationConfirm *SystemBearerContextModificationConfirm `aper:"optional,ext"`
}

// toIes transforms the BearerContextModificationConfirm struct into a slice of E1APMessageIEs.
func (msg *BearerContextModificationConfirm) toIes() ([]E1APMessageIE, error) {
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
	if msg.SystemBearerContextModificationConfirm != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDSystemBearerContextModificationConfirm},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SystemBearerContextModificationConfirm,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for BearerContextModificationConfirm.
func (msg *BearerContextModificationConfirm) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert BearerContextModificationConfirm to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduSuccessfulOutcome, ProcedureCode{Value: ProcedureCodeBearerContextModificationRequired}, Criticality{Value: CriticalityIgnore}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for BearerContextModificationConfirm.
func (msg *BearerContextModificationConfirm) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode BearerContextModificationConfirm failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := BearerContextModificationConfirmDecoder{
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
	if err != nil {
		return
	}

	return
}

type BearerContextModificationConfirmDecoder struct {
	msg      *BearerContextModificationConfirm
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *BearerContextModificationConfirmDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDSystemBearerContextModificationConfirm:
		msg.SystemBearerContextModificationConfirm = new(SystemBearerContextModificationConfirm)
		if err = msg.SystemBearerContextModificationConfirm.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode SystemBearerContextModificationConfirm failed: %w", err)
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
