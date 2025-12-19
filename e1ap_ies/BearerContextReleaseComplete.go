package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextReleaseComplete is a generated SEQUENCE type.
type BearerContextReleaseComplete struct {
	GNBCUCPUEE1APID               GNBCUCPUEE1APID                `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID               GNBCUUPUEE1APID                `aper:"lb:0,ub:4294967295,mandatory,ext"`
	CriticalityDiagnostics        *CriticalityDiagnostics        `aper:"optional,ext"`
	RetainabilityMeasurementsInfo *RetainabilityMeasurementsInfo `aper:"lb:1,ub:MaxnoofDRBs,optional,ext"`
}

// toIes transforms the BearerContextReleaseComplete struct into a slice of E1APMessageIEs.
func (msg *BearerContextReleaseComplete) toIes() ([]E1APMessageIE, error) {
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
	if msg.CriticalityDiagnostics != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDCriticalityDiagnostics},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.RetainabilityMeasurementsInfo != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDRetainabilityMeasurementsInfo},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RetainabilityMeasurementsInfo,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for BearerContextReleaseComplete.
func (msg *BearerContextReleaseComplete) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert BearerContextReleaseComplete to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduSuccessfulOutcome, ProcedureCode{Value: ProcedureCodeBearerContextRelease}, Criticality{Value: CriticalityIgnore}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for BearerContextReleaseComplete.
func (msg *BearerContextReleaseComplete) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode BearerContextReleaseComplete failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := BearerContextReleaseCompleteDecoder{
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

type BearerContextReleaseCompleteDecoder struct {
	msg      *BearerContextReleaseComplete
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *BearerContextReleaseCompleteDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDCriticalityDiagnostics:
		msg.CriticalityDiagnostics = new(CriticalityDiagnostics)
		if err = msg.CriticalityDiagnostics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode CriticalityDiagnostics failed: %w", err)
		}
	case ProtocolIEIDRetainabilityMeasurementsInfo:

		{
			itemDecoder := func(r *aper.AperReader) (*DRBRemovedItem, error) {

				item := new(DRBRemovedItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 1, Ub: MaxnoofDRBs}, false)
			if err != nil {
				return nil, fmt.Errorf("decode RetainabilityMeasurementsInfo failed: %w", err)
			}

			msg.RetainabilityMeasurementsInfo = new(RetainabilityMeasurementsInfo)
			msg.RetainabilityMeasurementsInfo.Value = decodedItems
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
