package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPE1SetupRequest is a generated SEQUENCE type.
type GNBCUUPE1SetupRequest struct {
	TransactionID             TransactionID              `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUUPID                 GNBCUUPID                  `aper:"lb:0,ub:68719476735,mandatory,ext"`
	GNBCUUPName               *aper.OctetString          `aper:"optional,ext"`
	CNSupport                 CNSupport                  `aper:"mandatory,ext"`
	SupportedPLMNs            []SupportedPLMNsItem       `aper:"lb:1,ub:MaxnoofSPLMNs,mandatory,ext"`
	GNBCUUPCapacity           *GNBCUUPCapacity           `aper:"lb:0,ub:255,optional,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ext"`
	ExtendedGNBCUUPName       *ExtendedGNBCUUPName       `aper:"optional,ext"`
}

// toIes transforms the GNBCUUPE1SetupRequest struct into a slice of E1APMessageIEs.
func (msg *GNBCUUPE1SetupRequest) toIes() ([]E1APMessageIE, error) {
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
	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDGNBCUUPID),
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 68719476735},
				ext:   false,
				Value: msg.GNBCUUPID.Value,
			},
		})
	}
	if msg.GNBCUUPName != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDGNBCUUPName),
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.GNBCUUPName.Value,
			},
		})
	}
	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDCNSupport),
			Criticality: Criticality{Value: CriticalityReject},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 2},
				ext:   true,
				Value: msg.CNSupport.Value,
			},
		})
	}
	{

		tmp_SupportedPLMNs := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 1, Ub: MaxnoofSPLMNs},
			ext: false,
		}

		for i := 0; i < len(msg.SupportedPLMNs); i++ {
			tmp_SupportedPLMNs.Value = append(tmp_SupportedPLMNs.Value, &msg.SupportedPLMNs[i])
		}

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDSupportedPLMNs),
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &tmp_SupportedPLMNs,
		})
	}
	if msg.GNBCUUPCapacity != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDGNBCUUPCapacity),
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 255},
				ext:   false,
				Value: msg.GNBCUUPCapacity.Value,
			},
		})
	}
	if msg.TransportLayerAddressInfo != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDTransportLayerAddressInfo),
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TransportLayerAddressInfo,
		})
	}
	if msg.ExtendedGNBCUUPName != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID(ProtocolIEIDExtendedGNBCUUPName),
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ExtendedGNBCUUPName,
		})
	}
	var err error
	return ies, err
}

// Encode for GNBCUUPE1SetupRequest: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for GNBCUUPE1SetupRequest.
func (msg *GNBCUUPE1SetupRequest) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GNBCUUPE1SetupRequest: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUUPE1SetupRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEIDGNBCUUPID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUUPID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDGNBCUUPID,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDCNSupport]; !ok {
		err = fmt.Errorf("mandatory field CNSupport is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDCNSupport,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDSupportedPLMNs]; !ok {
		err = fmt.Errorf("mandatory field SupportedPLMNs is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEIDSupportedPLMNs,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type GNBCUUPE1SetupRequestDecoder struct {
	msg      *GNBCUUPE1SetupRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *GNBCUUPE1SetupRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDGNBCUUPID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUUPID failed: %w", err)
			}
			msg.GNBCUUPID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPName:

		{
			var val []byte
			if val, err = ieR.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUUPName failed: %w", err)
			}
			msg.GNBCUUPName = new(aper.OctetString)
			msg.GNBCUUPName.Value = aper.OctetString(val)
		}
	case ProtocolIEIDCNSupport:

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return nil, fmt.Errorf("Decode CNSupport failed: %w", err)
			}
			msg.CNSupport.Value = aper.Enumerated(val)
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
			var decodedItems []SupportedPLMNsItem
			if decodedItems, err = aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 1, Ub: MaxnoofSPLMNs}, false); err != nil {
				return nil, fmt.Errorf("Decode SupportedPLMNs failed: %w", err)
			}
			msg.SupportedPLMNs = decodedItems
		}
	case ProtocolIEIDGNBCUUPCapacity:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUUPCapacity failed: %w", err)
			}
			msg.GNBCUUPCapacity = new(GNBCUUPCapacity)
			msg.GNBCUUPCapacity.Value = aper.Integer(val)
		}
	case ProtocolIEIDTransportLayerAddressInfo:
		msg.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = msg.TransportLayerAddressInfo.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode TransportLayerAddressInfo failed: %w", err)
		}
	case ProtocolIEIDExtendedGNBCUUPName:
		msg.ExtendedGNBCUUPName = new(ExtendedGNBCUUPName)
		if err = msg.ExtendedGNBCUUPName.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode ExtendedGNBCUUPName failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
