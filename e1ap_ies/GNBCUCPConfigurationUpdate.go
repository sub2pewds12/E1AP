package e1ap_ies

import (
	"bytes"
	"fmt"

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
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 255},
					ext:   true,
					Value: msg.TransactionID.Value,
				},
			})
		}
	}
	if msg.GNBCUCPName != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPName},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   false,
					Value: msg.GNBCUCPName.Value,
				},
			})
		}
	}
	if msg.GNBCUCPTNLAToAddList != nil {

		tmp_GNBCUCPTNLAToAddList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
			ext: false,
		}

		for i := 0; i < len(msg.GNBCUCPTNLAToAddList.Value); i++ {
			tmp_GNBCUCPTNLAToAddList.Value = append(tmp_GNBCUCPTNLAToAddList.Value, &msg.GNBCUCPTNLAToAddList.Value[i])
		}

		{

			tmp_GNBCUCPTNLAToAddList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
				ext: false,
			}

			for i := 0; i < len(msg.GNBCUCPTNLAToAddList.Value); i++ {
				tmp_GNBCUCPTNLAToAddList.Value = append(tmp_GNBCUCPTNLAToAddList.Value, &msg.GNBCUCPTNLAToAddList.Value[i])
			}

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPTNLAToAddList},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_GNBCUCPTNLAToAddList,
			})
		}
	}
	if msg.GNBCUCPTNLAToRemoveList != nil {

		tmp_GNBCUCPTNLAToRemoveList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
			ext: false,
		}

		for i := 0; i < len(msg.GNBCUCPTNLAToRemoveList.Value); i++ {
			tmp_GNBCUCPTNLAToRemoveList.Value = append(tmp_GNBCUCPTNLAToRemoveList.Value, &msg.GNBCUCPTNLAToRemoveList.Value[i])
		}

		{

			tmp_GNBCUCPTNLAToRemoveList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
				ext: false,
			}

			for i := 0; i < len(msg.GNBCUCPTNLAToRemoveList.Value); i++ {
				tmp_GNBCUCPTNLAToRemoveList.Value = append(tmp_GNBCUCPTNLAToRemoveList.Value, &msg.GNBCUCPTNLAToRemoveList.Value[i])
			}

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPTNLAToRemoveList},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_GNBCUCPTNLAToRemoveList,
			})
		}
	}
	if msg.GNBCUCPTNLAToUpdateList != nil {

		tmp_GNBCUCPTNLAToUpdateList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
			ext: false,
		}

		for i := 0; i < len(msg.GNBCUCPTNLAToUpdateList.Value); i++ {
			tmp_GNBCUCPTNLAToUpdateList.Value = append(tmp_GNBCUCPTNLAToUpdateList.Value, &msg.GNBCUCPTNLAToUpdateList.Value[i])
		}

		{

			tmp_GNBCUCPTNLAToUpdateList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
				ext: false,
			}

			for i := 0; i < len(msg.GNBCUCPTNLAToUpdateList.Value); i++ {
				tmp_GNBCUCPTNLAToUpdateList.Value = append(tmp_GNBCUCPTNLAToUpdateList.Value, &msg.GNBCUCPTNLAToUpdateList.Value[i])
			}

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPTNLAToUpdateList},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_GNBCUCPTNLAToUpdateList,
			})
		}
	}
	if msg.TransportLayerAddressInfo != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDTransportLayerAddressInfo},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.TransportLayerAddressInfo,
			})
		}
	}
	if msg.ExtendedGNBCUCPName != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDExtendedGNBCUCPName},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.ExtendedGNBCUCPName,
			})
		}
	}
	var err error
	return ies, err
}

// Encode for GNBCUCPConfigurationUpdate: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for GNBCUCPConfigurationUpdate.
func (msg *GNBCUCPConfigurationUpdate) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GNBCUCPConfigurationUpdate: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUCPConfigurationUpdateDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDTransactionID}]; !ok {
		err = fmt.Errorf("mandatory field TransactionID is missing")
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
		return nil, fmt.Errorf("duplicated protocol IE ID %d", ieId.Value)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id.Value {
	case ProtocolIEIDTransactionID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
				return nil, fmt.Errorf("Decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUCPName:

		{
			var val []byte
			if val, err = ieR.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUCPName failed: %w", err)
			}
			msg.GNBCUCPName = new(GNBCUCPName)
			msg.GNBCUCPName.Value = aper.OctetString(val)
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
			var decodedItems []GNBCUCPTNLAToAddItem
			if decodedItems, err = aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUCPTNLAToAddList failed: %w", err)
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
			var decodedItems []GNBCUCPTNLAToRemoveItem
			if decodedItems, err = aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUCPTNLAToRemoveList failed: %w", err)
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
			var decodedItems []GNBCUCPTNLAToUpdateItem
			if decodedItems, err = aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUCPTNLAToUpdateList failed: %w", err)
			}

			msg.GNBCUCPTNLAToUpdateList = new(GNBCUCPTNLAToUpdateList)
			msg.GNBCUCPTNLAToUpdateList.Value = decodedItems
		}
	case ProtocolIEIDTransportLayerAddressInfo:
		msg.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = msg.TransportLayerAddressInfo.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode TransportLayerAddressInfo failed: %w", err)
		}
	case ProtocolIEIDExtendedGNBCUCPName:
		msg.ExtendedGNBCUCPName = new(ExtendedGNBCUCPName)
		if err = msg.ExtendedGNBCUCPName.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode ExtendedGNBCUCPName failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
