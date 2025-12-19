package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPConfigurationUpdateAcknowledge is a generated SEQUENCE type.
type GNBCUCPConfigurationUpdateAcknowledge struct {
	TransactionID                TransactionID                 `aper:"lb:0,ub:255,mandatory,ext"`
	CriticalityDiagnostics       *CriticalityDiagnostics       `aper:"optional,ext"`
	GNBCUCPTNLASetupList         *GNBCUCPTNLASetupList         `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	GNBCUCPTNLAFailedToSetupList *GNBCUCPTNLAFailedToSetupList `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	TransportLayerAddressInfo    *TransportLayerAddressInfo    `aper:"optional,ext"`
}

// toIes transforms the GNBCUCPConfigurationUpdateAcknowledge struct into a slice of E1APMessageIEs.
func (msg *GNBCUCPConfigurationUpdateAcknowledge) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TransactionID,
	})
	if msg.CriticalityDiagnostics != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDCriticalityDiagnostics},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.GNBCUCPTNLASetupList != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPTNLASetupList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUCPTNLASetupList,
		})
	}
	if msg.GNBCUCPTNLAFailedToSetupList != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPTNLAFailedToSetupList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUCPTNLAFailedToSetupList,
		})
	}
	if msg.TransportLayerAddressInfo != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDTransportLayerAddressInfo},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TransportLayerAddressInfo,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for GNBCUCPConfigurationUpdateAcknowledge.
func (msg *GNBCUCPConfigurationUpdateAcknowledge) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert GNBCUCPConfigurationUpdateAcknowledge to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduSuccessfulOutcome, ProcedureCode{Value: ProcedureCodeGNBCUCPConfigurationUpdate}, Criticality{Value: CriticalityIgnore}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for GNBCUCPConfigurationUpdateAcknowledge.
func (msg *GNBCUCPConfigurationUpdateAcknowledge) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode GNBCUCPConfigurationUpdateAcknowledge failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUCPConfigurationUpdateAcknowledgeDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf(decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDTransactionID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field TransactionID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type GNBCUCPConfigurationUpdateAcknowledgeDecoder struct {
	msg      *GNBCUCPConfigurationUpdateAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUCPConfigurationUpdateAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDTransactionID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true)
			if err != nil {
				return nil, fmt.Errorf("decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = aper.Integer(val)
		}
	case ProtocolIEIDCriticalityDiagnostics:
		msg.CriticalityDiagnostics = new(CriticalityDiagnostics)
		if err = msg.CriticalityDiagnostics.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode CriticalityDiagnostics failed: %w", err)
		}
	case ProtocolIEIDGNBCUCPTNLASetupList:

		{
			itemDecoder := func(r *aper.AperReader) (*GNBCUCPTNLASetupItem, error) {

				item := new(GNBCUCPTNLASetupItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPTNLASetupList failed: %w", err)
			}

			msg.GNBCUCPTNLASetupList = new(GNBCUCPTNLASetupList)
			msg.GNBCUCPTNLASetupList.Value = decodedItems
		}
	case ProtocolIEIDGNBCUCPTNLAFailedToSetupList:

		{
			itemDecoder := func(r *aper.AperReader) (*GNBCUCPTNLAFailedToSetupItem, error) {

				item := new(GNBCUCPTNLAFailedToSetupItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPTNLAFailedToSetupList failed: %w", err)
			}

			msg.GNBCUCPTNLAFailedToSetupList = new(GNBCUCPTNLAFailedToSetupList)
			msg.GNBCUCPTNLAFailedToSetupList.Value = decodedItems
		}
	case ProtocolIEIDTransportLayerAddressInfo:
		msg.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = msg.TransportLayerAddressInfo.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode TransportLayerAddressInfo failed: %w", err)
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
