package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
)

// IABUPTNLAddressUpdate is a generated SEQUENCE type.
type IABUPTNLAddressUpdate struct {
	TransactionID              TransactionID
	DLUPTNLAddressToUpdateList *DLUPTNLAddressToUpdateList
}

// toIes transforms the IABUPTNLAddressUpdate struct into a slice of E1APMessageIEs.
func (msg *IABUPTNLAddressUpdate) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TransactionID,
	})
	if msg.DLUPTNLAddressToUpdateList != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDDLUPTNLAddressToUpdateList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DLUPTNLAddressToUpdateList,
		})
	}
	return ies, nil
}

func (msg *IABUPTNLAddressUpdate) EncodeWithEncoder(e *per.Encoder) (err error) {
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

func (msg *IABUPTNLAddressUpdate) Encode(w io.Writer) error {
	e := per.NewEncoder(per.APER)
	if err := msg.EncodeWithEncoder(e); err != nil {
		return err
	}
	_, err := w.Write(e.Bytes())
	return err
}

// Decode implements the MessageUnmarshaller interface for IABUPTNLAddressUpdate.
func (msg *IABUPTNLAddressUpdate) Decode(data []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	r := per.NewDecoder(data, per.APER)
	return msg.DecodeFromDecoder(r)
}

func (msg *IABUPTNLAddressUpdate) DecodeFromDecoder(r *per.Decoder) (diagList []CriticalityDiagnosticsIEItem, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("decode IABUPTNLAddressUpdate failed: %w", err)
		}
	}()

	decoder := IABUPTNLAddressUpdateDecoder{
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

type IABUPTNLAddressUpdateDecoder struct {
	msg      *IABUPTNLAddressUpdate
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *IABUPTNLAddressUpdateDecoder) decodeIE(r *per.Decoder) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDTransactionID:

		{
			val, err := ieR.DecodeInteger(per.ConstrainedExtensible(0, 255))
			if err != nil {
				return nil, fmt.Errorf("decode TransactionID failed: %w", err)
			}
			msg.TransactionID.Value = val
		}
	case ProtocolIEIDDLUPTNLAddressToUpdateList:
		msg.DLUPTNLAddressToUpdateList = new(DLUPTNLAddressToUpdateList)

		{
			itemDecoder := func(r *per.Decoder) (*DLUPTNLAddressToUpdateItem, error) {
				item := new(DLUPTNLAddressToUpdateItem)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}

			c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxnoofTNLAddresses)}
			length, err := ieR.DecodeLengthDeterminant(c)
			if err != nil {
				return nil, fmt.Errorf("decode struct list length failed: %w", err)
			}
			for i := int64(0); i < length; i++ {
				item, err := itemDecoder(ieR)
				if err != nil {
					return nil, fmt.Errorf("decode item failed: %w", err)
				}
				msg.DLUPTNLAddressToUpdateList.Value = append(msg.DLUPTNLAddressToUpdateList.Value, *item)
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
