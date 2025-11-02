package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPConfigurationUpdateAcknowledge is a generated SEQUENCE type.
type GNBCUUPConfigurationUpdateAcknowledge struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	CriticalityDiagnostics    *CriticalityDiagnostics    `aper:"optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
}

// toIes transforms the GNBCUUPConfigurationUpdateAcknowledge struct into a slice of E1APMessageIEs.
func (msg *GNBCUUPConfigurationUpdateAcknowledge) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDTransactionID),
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 255},
				ext:   true,
				Value: msg.TransactionID.Value,
			},
		})
	}
	if msg.CriticalityDiagnostics != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDCriticalityDiagnostics),
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.TransportLayerAddressInfo != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDTransportLayerAddressInfo),
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TransportLayerAddressInfo,
		})
	}
	var err error
	return ies, err
}

// Encode for GNBCUUPConfigurationUpdateAcknowledge: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for GNBCUUPConfigurationUpdateAcknowledge.
func (msg *GNBCUUPConfigurationUpdateAcknowledge) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GNBCUUPConfigurationUpdateAcknowledge: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUUPConfigurationUpdateAcknowledgeDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEIDTransactionID]; !ok {
		err = fmt.Errorf("mandatory field TransactionID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDTransactionID,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type GNBCUUPConfigurationUpdateAcknowledgeDecoder struct {
	msg      *GNBCUUPConfigurationUpdateAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUUPConfigurationUpdateAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDTransactionID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
				return nil, fmt.Errorf("Decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = aper.Integer(val)
		}
	case ProtocolIEIDCriticalityDiagnostics:
		msg.CriticalityDiagnostics = new(CriticalityDiagnostics)
		if err = msg.CriticalityDiagnostics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode CriticalityDiagnostics failed: %w", err)
		}
	case ProtocolIEIDTransportLayerAddressInfo:
		msg.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = msg.TransportLayerAddressInfo.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode TransportLayerAddressInfo failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
