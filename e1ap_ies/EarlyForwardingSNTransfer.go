package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// EarlyForwardingSNTransfer is a generated SEQUENCE type.
type EarlyForwardingSNTransfer struct {
	GNBCUCPUEE1APID                  GNBCUCPUEE1APID                    `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                  GNBCUUPUEE1APID                    `aper:"lb:0,ub:4294967295,mandatory,ext"`
	DRBsSubjectToEarlyForwardingList []DRBsSubjectToEarlyForwardingItem `aper:"mandatory,ext"`
}

// toIes transforms the EarlyForwardingSNTransfer struct into a slice of E1APMessageIEs.
func (msg *EarlyForwardingSNTransfer) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: msg.GNBCUCPUEE1APID.Value,
				},
			})
		}
	}
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: msg.GNBCUUPUEE1APID.Value,
				},
			})
		}
	}
	{

		tmp_DRBsSubjectToEarlyForwardingList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}

		for i := 0; i < len(msg.DRBsSubjectToEarlyForwardingList); i++ {
			tmp_DRBsSubjectToEarlyForwardingList.Value = append(tmp_DRBsSubjectToEarlyForwardingList.Value, &msg.DRBsSubjectToEarlyForwardingList[i])
		}

		{

			tmp_DRBsSubjectToEarlyForwardingList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: 0},
				ext: false,
			}

			for i := 0; i < len(msg.DRBsSubjectToEarlyForwardingList); i++ {
				tmp_DRBsSubjectToEarlyForwardingList.Value = append(tmp_DRBsSubjectToEarlyForwardingList.Value, &msg.DRBsSubjectToEarlyForwardingList[i])
			}

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDDRBsSubjectToEarlyForwardingList},
				Criticality: Criticality{Value: CriticalityReject},
				Value:       &tmp_DRBsSubjectToEarlyForwardingList,
			})
		}
	}
	var err error
	return ies, err
}

// Encode implements the aper.AperMarshaller interface for EarlyForwardingSNTransfer.
func (msg *EarlyForwardingSNTransfer) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert EarlyForwardingSNTransfer to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCodeEarlyForwardingSNTransfer, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for EarlyForwardingSNTransfer.
func (msg *EarlyForwardingSNTransfer) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("EarlyForwardingSNTransfer: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := EarlyForwardingSNTransferDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEIDGNBCUCPUEE1APID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUCPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDGNBCUCPUEE1APID,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDGNBCUUPUEE1APID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUUPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDGNBCUUPUEE1APID,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDDRBsSubjectToEarlyForwardingList]; !ok {
		err = fmt.Errorf("mandatory field DRBsSubjectToEarlyForwardingList is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDDRBsSubjectToEarlyForwardingList,
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
		return nil, fmt.Errorf("duplicated protocol IE ID %%d", ieId)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id {
	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUCPUEE1APID failed: %w", err)
			}
			msg.GNBCUCPUEE1APID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
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
			var decodedItems []DRBsSubjectToEarlyForwardingItem
			if decodedItems, err = aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return nil, fmt.Errorf("Decode DRBsSubjectToEarlyForwardingList failed: %w", err)
			}
			msg.DRBsSubjectToEarlyForwardingList = decodedItems
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
