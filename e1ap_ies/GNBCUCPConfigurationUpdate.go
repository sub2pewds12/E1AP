package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPConfigurationUpdate is a generated SEQUENCE type.
type GNBCUCPConfigurationUpdate struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPName               *GNBCUCPName               `aper:"optional,ext"`
	GNBCUCPTNLAToAddList      *GNBCUCPTNLAToAddList      `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	GNBCUCPTNLAToRemoveList   *GNBCUCPTNLAToRemoveList   `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	GNBCUCPTNLAToUpdateList   *GNBCUCPTNLAToUpdateList   `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
	ExtendedGNBCUCPName       *ExtendedGNBCUCPName       `aper:"optional,ext"`
}

// toIes transforms the GNBCUCPConfigurationUpdate struct into a slice of E1APMessageIEs.
func (msg *GNBCUCPConfigurationUpdate) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TransactionID,
	})
	if msg.GNBCUCPName != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPName},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUCPName,
		})
	}
	if msg.GNBCUCPTNLAToAddList != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPTNLAToAddList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUCPTNLAToAddList,
		})
	}
	if msg.GNBCUCPTNLAToRemoveList != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPTNLAToRemoveList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUCPTNLAToRemoveList,
		})
	}
	if msg.GNBCUCPTNLAToUpdateList != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPTNLAToUpdateList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUCPTNLAToUpdateList,
		})
	}
	if msg.TransportLayerAddressInfo != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDTransportLayerAddressInfo},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TransportLayerAddressInfo,
		})
	}
	if msg.ExtendedGNBCUCPName != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDExtendedGNBCUCPName},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ExtendedGNBCUCPName,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for GNBCUCPConfigurationUpdate.
func (msg *GNBCUCPConfigurationUpdate) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert GNBCUCPConfigurationUpdate to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodeGNBCUCPConfigurationUpdate}, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for GNBCUCPConfigurationUpdate.
func (msg *GNBCUCPConfigurationUpdate) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode GNBCUCPConfigurationUpdate failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUCPConfigurationUpdateDecoder{
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

type GNBCUCPConfigurationUpdateDecoder struct {
	msg      *GNBCUCPConfigurationUpdate
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUCPConfigurationUpdateDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDGNBCUCPName:
		msg.GNBCUCPName = new(GNBCUCPName)
		if err = msg.GNBCUCPName.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode GNBCUCPName failed: %w", err)
		}
	case ProtocolIEIDGNBCUCPTNLAToAddList:

		{
			itemDecoder := func(r *aper.AperReader) (*GNBCUCPTNLAToAddItem, error) {

				item := new(GNBCUCPTNLAToAddItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPTNLAToAddList failed: %w", err)
			}

			msg.GNBCUCPTNLAToAddList = new(GNBCUCPTNLAToAddList)
			msg.GNBCUCPTNLAToAddList.Value = decodedItems
		}
	case ProtocolIEIDGNBCUCPTNLAToRemoveList:

		{
			itemDecoder := func(r *aper.AperReader) (*GNBCUCPTNLAToRemoveItem, error) {

				item := new(GNBCUCPTNLAToRemoveItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPTNLAToRemoveList failed: %w", err)
			}

			msg.GNBCUCPTNLAToRemoveList = new(GNBCUCPTNLAToRemoveList)
			msg.GNBCUCPTNLAToRemoveList.Value = decodedItems
		}
	case ProtocolIEIDGNBCUCPTNLAToUpdateList:

		{
			itemDecoder := func(r *aper.AperReader) (*GNBCUCPTNLAToUpdateItem, error) {

				item := new(GNBCUCPTNLAToUpdateItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPTNLAToUpdateList failed: %w", err)
			}

			msg.GNBCUCPTNLAToUpdateList = new(GNBCUCPTNLAToUpdateList)
			msg.GNBCUCPTNLAToUpdateList.Value = decodedItems
		}
	case ProtocolIEIDTransportLayerAddressInfo:
		msg.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = msg.TransportLayerAddressInfo.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode TransportLayerAddressInfo failed: %w", err)
		}
	case ProtocolIEIDExtendedGNBCUCPName:
		msg.ExtendedGNBCUCPName = new(ExtendedGNBCUCPName)
		if err = msg.ExtendedGNBCUCPName.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode ExtendedGNBCUCPName failed: %w", err)
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
