package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
)

// MRDCDataUsageReport is a generated SEQUENCE type.
type MRDCDataUsageReport struct {
	GNBCUCPUEE1APID                 GNBCUCPUEE1APID
	GNBCUUPUEE1APID                 GNBCUUPUEE1APID
	PDUSessionResourceDataUsageList PDUSessionResourceDataUsageList
}

// toIes transforms the MRDCDataUsageReport struct into a slice of E1APMessageIEs.
func (msg *MRDCDataUsageReport) toIes() ([]E1APMessageIE, error) {
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
		ID:          ProtocolIEID{Value: ProtocolIEIDPDUSessionResourceDataUsageList},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.PDUSessionResourceDataUsageList,
	})
	return ies, nil
}

func (msg *MRDCDataUsageReport) EncodeWithEncoder(e *per.Encoder) (err error) {
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

func (msg *MRDCDataUsageReport) Encode(w io.Writer) error {
	e := per.NewEncoder(per.APER)
	if err := msg.EncodeWithEncoder(e); err != nil {
		return err
	}
	_, err := w.Write(e.Bytes())
	return err
}

// Decode implements the MessageUnmarshaller interface for MRDCDataUsageReport.
func (msg *MRDCDataUsageReport) Decode(data []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	r := per.NewDecoder(data, per.APER)
	return msg.DecodeFromDecoder(r)
}

func (msg *MRDCDataUsageReport) DecodeFromDecoder(r *per.Decoder) (diagList []CriticalityDiagnosticsIEItem, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("decode MRDCDataUsageReport failed: %w", err)
		}
	}()

	decoder := MRDCDataUsageReportDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDPDUSessionResourceDataUsageList}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field PDUSessionResourceDataUsageList is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDPDUSessionResourceDataUsageList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type MRDCDataUsageReportDecoder struct {
	msg      *MRDCDataUsageReport
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *MRDCDataUsageReportDecoder) decodeIE(r *per.Decoder) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDPDUSessionResourceDataUsageList:

		{
			itemDecoder := func(r *per.Decoder) (*PDUSessionResourceDataUsageItem, error) {
				item := new(PDUSessionResourceDataUsageItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}

			c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxnoofPDUSessionResource)}
			length, err := ieR.DecodeLengthDeterminant(c)
			if err != nil {
				return nil, fmt.Errorf("decode struct list length failed: %w", err)
			}
			for i := int64(0); i < length; i++ {
				item, err := itemDecoder(ieR)
				if err != nil {
					return nil, fmt.Errorf("decode item failed: %w", err)
				}
				msg.PDUSessionResourceDataUsageList.Value = append(msg.PDUSessionResourceDataUsageList.Value, *item)
			}
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
