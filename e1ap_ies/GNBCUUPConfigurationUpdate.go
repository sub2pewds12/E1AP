package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPConfigurationUpdate is a generated SEQUENCE type.
type GNBCUUPConfigurationUpdate struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUUPID                 GNBCUUPID                  `aper:"lb:0,ub:68719476735,mandatory,ext"`
	GNBCUUPName               *GNBCUUPName               `aper:"optional,ext"`
	SupportedPLMNs            *SupportedPLMNsList        `aper:"lb:1,ub:MaxnoofSPLMNs,optional,ext"`
	GNBCUUPCapacity           *GNBCUUPCapacity           `aper:"lb:0,ub:255,optional,ext"`
	GNBCUUPTNLAToRemoveList   *GNBCUUPTNLAToRemoveList   `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
	ExtendedGNBCUUPName       *ExtendedGNBCUUPName       `aper:"optional,ext"`
}

// toIes transforms the GNBCUUPConfigurationUpdate struct into a slice of E1APMessageIEs.
func (msg *GNBCUUPConfigurationUpdate) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
		Criticality: Criticality{Value: CriticalityReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   true,
			Value: msg.TransactionID.Value,
		},
	})

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPID},
		Criticality: Criticality{Value: CriticalityReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 68719476735},
			ext:   false,
			Value: msg.GNBCUUPID.Value,
		},
	})
	if msg.GNBCUUPName != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPName},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.GNBCUUPName.Value,
			},
		})
	}
	if msg.SupportedPLMNs != nil {

		tmpSupportedPLMNs := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 1, Ub: MaxnoofSPLMNs},
			ext: false,
		}

		for i := 0; i < len(msg.SupportedPLMNs.Value); i++ {
			tmpSupportedPLMNs.Value = append(tmpSupportedPLMNs.Value, &msg.SupportedPLMNs.Value[i])
		}

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDSupportedPLMNs},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &tmpSupportedPLMNs,
		})
	}
	if msg.GNBCUUPCapacity != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPCapacity},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 255},
				ext:   false,
				Value: msg.GNBCUUPCapacity.Value,
			},
		})
	}
	if msg.GNBCUUPTNLAToRemoveList != nil {

		tmpGNBCUUPTNLAToRemoveList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
			ext: false,
		}

		for i := 0; i < len(msg.GNBCUUPTNLAToRemoveList.Value); i++ {
			tmpGNBCUUPTNLAToRemoveList.Value = append(tmpGNBCUUPTNLAToRemoveList.Value, &msg.GNBCUUPTNLAToRemoveList.Value[i])
		}

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPTNLAToRemoveList},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &tmpGNBCUUPTNLAToRemoveList,
		})
	}
	if msg.TransportLayerAddressInfo != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDTransportLayerAddressInfo},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TransportLayerAddressInfo,
		})
	}
	if msg.ExtendedGNBCUUPName != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDExtendedGNBCUUPName},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ExtendedGNBCUUPName,
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for GNBCUUPConfigurationUpdate.
func (msg *GNBCUUPConfigurationUpdate) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert GNBCUUPConfigurationUpdate to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodeGNBCUUPConfigurationUpdate}, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for GNBCUUPConfigurationUpdate.
func (msg *GNBCUUPConfigurationUpdate) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode GNBCUUPConfigurationUpdate failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUUPConfigurationUpdateDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUUPID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field GNBCUUPID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type GNBCUUPConfigurationUpdateDecoder struct {
	msg      *GNBCUUPConfigurationUpdate
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUUPConfigurationUpdateDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false)
	if err != nil {
		return nil, err
	}
	msgIe = new(E1APMessageIE)
	msgIe.Id = ProtocolIEID{Value: aper.Integer(id)}
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return nil, err
	}
	msgIe.Criticality = Criticality{Value: aper.Enumerated(c)}

	buf, err := r.ReadOpenType()
	if err != nil {
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
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true)
			if err != nil {
				return nil, fmt.Errorf("decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPID failed: %w", err)
			}
			msg.GNBCUUPID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPName:
		msg.GNBCUUPName = new(GNBCUUPName)
		if err = msg.GNBCUUPName.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode GNBCUUPName failed: %w", err)
		}
	case ProtocolIEIDSupportedPLMNs:

		{
			itemDecoder := func(r *aper.AperReader) (*SupportedPLMNsItem, error) {

				item := new(SupportedPLMNsItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 1, Ub: MaxnoofSPLMNs}, false)
			if err != nil {
				return nil, fmt.Errorf("decode SupportedPLMNs failed: %w", err)
			}

			msg.SupportedPLMNs = new(SupportedPLMNsList)
			msg.SupportedPLMNs.Value = decodedItems
		}
	case ProtocolIEIDGNBCUUPCapacity:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPCapacity failed: %w", err)
			}
			msg.GNBCUUPCapacity = new(GNBCUUPCapacity)
			msg.GNBCUUPCapacity.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPTNLAToRemoveList:

		{
			itemDecoder := func(r *aper.AperReader) (*GNBCUUPTNLAToRemoveItem, error) {

				item := new(GNBCUUPTNLAToRemoveItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPTNLAToRemoveList failed: %w", err)
			}

			msg.GNBCUUPTNLAToRemoveList = new(GNBCUUPTNLAToRemoveList)
			msg.GNBCUUPTNLAToRemoveList.Value = decodedItems
		}
	case ProtocolIEIDTransportLayerAddressInfo:
		msg.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = msg.TransportLayerAddressInfo.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode TransportLayerAddressInfo failed: %w", err)
		}
	case ProtocolIEIDExtendedGNBCUUPName:
		msg.ExtendedGNBCUUPName = new(ExtendedGNBCUUPName)
		if err = msg.ExtendedGNBCUUPName.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode ExtendedGNBCUUPName failed: %w", err)
		}
	default:
		switch msgIe.Criticality.Value {
		case CriticalityReject:
			// If an unknown IE is critical, the PDU cannot be processed.
			return nil, fmt.Errorf("not comprehended IE ID %d (criticality: reject)", msgIe.Id.Value)
		case CriticalityNotify:
			// Per 3GPP TS 38.463 Section 10.3, report and proceed.
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.Id,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotUnderstood},
			})
		case CriticalityIgnore:
			// Ignore and proceed.
		}
	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
