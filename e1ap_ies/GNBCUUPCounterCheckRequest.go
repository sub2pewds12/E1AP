package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPCounterCheckRequest is a generated SEQUENCE type.
type GNBCUUPCounterCheckRequest struct {
	GNBCUCPUEE1APID                  GNBCUCPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                  GNBCUUPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,ext"`
	SystemGNBCUUPCounterCheckRequest SystemGNBCUUPCounterCheckRequest `aper:"mandatory,ext"`
}

// toIes transforms the GNBCUUPCounterCheckRequest struct into a slice of E1APMessageIEs.
func (msg *GNBCUUPCounterCheckRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDGNBCUCPUEE1APID),
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: msg.GNBCUCPUEE1APID.Value,
			},
		})
	}
	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDGNBCUUPUEE1APID),
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: msg.GNBCUUPUEE1APID.Value,
			},
		})
	}
	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDSystemGNBCUUPCounterCheckRequest),
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &msg.SystemGNBCUUPCounterCheckRequest,
		})
	}
	var err error
	return ies, err
}

// Encode for GNBCUUPCounterCheckRequest: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for GNBCUUPCounterCheckRequest.
func (msg *GNBCUUPCounterCheckRequest) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GNBCUUPCounterCheckRequest: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUUPCounterCheckRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEIDSystemGNBCUUPCounterCheckRequest]; !ok {
		err = fmt.Errorf("mandatory field SystemGNBCUUPCounterCheckRequest is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDSystemGNBCUUPCounterCheckRequest,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type GNBCUUPCounterCheckRequestDecoder struct {
	msg      *GNBCUUPCounterCheckRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUUPCounterCheckRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return nil, err
	}
	msgIe = new(E1APMessageIE)
	msgIe.Id = ProtocolIEID(id)

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
	case ProtocolIEIDSystemGNBCUUPCounterCheckRequest:
		if err = msg.SystemGNBCUUPCounterCheckRequest.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode SystemGNBCUUPCounterCheckRequest failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
