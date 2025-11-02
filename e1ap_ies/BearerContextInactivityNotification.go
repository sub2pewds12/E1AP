package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextInactivityNotification is a generated SEQUENCE type.
type BearerContextInactivityNotification struct {
	GNBCUCPUEE1APID     GNBCUCPUEE1APID     `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID     GNBCUUPUEE1APID     `aper:"lb:0,ub:4294967295,mandatory,ext"`
	ActivityInformation ActivityInformation `aper:"mandatory,ext"`
}

func (msg *BearerContextInactivityNotification) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDGNBCUCPUEE1APID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(msg.GNBCUCPUEE1APID),
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDGNBCUUPUEE1APID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(msg.GNBCUUPUEE1APID),
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDActivityInformation,
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &msg.ActivityInformation,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for BearerContextInactivityNotification.
func (msg *BearerContextInactivityNotification) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert BearerContextInactivityNotification to IEs: %w", err)
	}

	return EncodeInitiatingMessage(w, ProcedureCodeBearerContextInactivityNotification, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for BearerContextInactivityNotification.
func (msg *BearerContextInactivityNotification) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("BearerContextInactivityNotification: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := BearerContextInactivityNotificationDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEIDGNBCUCPUEE1APID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUCPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDGNBCUUPUEE1APID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUUPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDActivityInformation]; !ok {
		err = fmt.Errorf("mandatory field ActivityInformation is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDActivityInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type BearerContextInactivityNotificationDecoder struct {
	msg      *BearerContextInactivityNotification
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*E1APMessageIE
}

func (decoder *BearerContextInactivityNotificationDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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

	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return fmt.Errorf("Decode GNBCUCPUEE1APID failed: %w", err)
			}
			s.GNBCUCPUEE1APID = GNBCUCPUEE1APID(val)
		}

	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
			}
			s.GNBCUUPUEE1APID = GNBCUUPUEE1APID(val)
		}

	case ProtocolIEIDActivityInformation:
		if err = s.ActivityInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode ActivityInformation failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
	}
	return
}
