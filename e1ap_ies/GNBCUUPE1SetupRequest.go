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

func (msg *GNBCUUPE1SetupRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDTransactionID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 255},
				ext:   true,
				Value: aper.Integer(msg.TransactionID),
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDGNBCUUPID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 68719476735},
				ext:   false,
				Value: aper.Integer(msg.GNBCUUPID),
			},
		})
	}
	if msg.GNBCUUPName != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUUPName,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   false,
					Value: aper.OctetString((*msg.GNBCUUPName)),
				},
			})
		}
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDCNSupport,
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
			Id:          ProtocolIEIDSupportedPLMNs,
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &tmp_SupportedPLMNs,
		})
	}
	if msg.GNBCUUPCapacity != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUUPCapacity,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 255},
					ext:   false,
					Value: aper.Integer((*msg.GNBCUUPCapacity)),
				},
			})
		}
	}
	if msg.TransportLayerAddressInfo != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDTransportLayerAddressInfo,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.TransportLayerAddressInfo,
			})
		}
	}
	if msg.ExtendedGNBCUUPName != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDExtendedGNBCUUPName,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.ExtendedGNBCUUPName,
			})
		}
	}
	return ies, nil
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
		list: make(map[aper.Integer]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEIDTransactionID]; !ok {
		err = fmt.Errorf("mandatory field TransactionID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDTransactionID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDGNBCUUPID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUUPID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDCNSupport]; !ok {
		err = fmt.Errorf("mandatory field CNSupport is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDCNSupport},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDSupportedPLMNs]; !ok {
		err = fmt.Errorf("mandatory field SupportedPLMNs is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDSupportedPLMNs},
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
	list     map[aper.Integer]*E1APMessageIE
}

func (decoder *GNBCUUPE1SetupRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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

	case ProtocolIEIDTransactionID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
				return fmt.Errorf("Decode TransactionID failed: %w", err)
			}
			s.TransactionID = TransactionID(val)
		}

	case ProtocolIEIDGNBCUUPID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPID failed: %w", err)
			}
			s.GNBCUUPID = GNBCUUPID(val)
		}

	case ProtocolIEIDGNBCUUPName:

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPName failed: %w", err)
			}
			tmp := aper.OctetString(val)
			s.GNBCUUPName = &tmp
		}

	case ProtocolIEIDCNSupport:

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return fmt.Errorf("Decode CNSupport failed: %w", err)
			}
			s.CNSupport.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDSupportedPLMNs:
		if err = s.SupportedPLMNs.Decode(r); err != nil {
			return fmt.Errorf("Decode SupportedPLMNs failed: %w", err)
		}

	case ProtocolIEIDGNBCUUPCapacity:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPCapacity failed: %w", err)
			}
			tmp := GNBCUUPCapacity(val)
			s.GNBCUUPCapacity = &tmp
		}

	case ProtocolIEIDTransportLayerAddressInfo:
		s.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = s.TransportLayerAddressInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode TransportLayerAddressInfo failed: %w", err)
		}

	case ProtocolIEIDExtendedGNBCUUPName:
		s.ExtendedGNBCUUPName = new(ExtendedGNBCUUPName)
		if err = s.ExtendedGNBCUUPName.Decode(r); err != nil {
			return fmt.Errorf("Decode ExtendedGNBCUUPName failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
	}
	return
}
