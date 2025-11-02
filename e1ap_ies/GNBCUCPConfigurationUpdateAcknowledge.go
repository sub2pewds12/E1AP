package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPConfigurationUpdateAcknowledge is a generated SEQUENCE type.
type GNBCUCPConfigurationUpdateAcknowledge struct {
	TransactionID                TransactionID                  `aper:"lb:0,ub:255,mandatory,ext"`
	CriticalityDiagnostics       *CriticalityDiagnostics        `aper:"optional,ext"`
	GNBCUCPTNLASetupList         []GNBCUCPTNLASetupItem         `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	GNBCUCPTNLAFailedToSetupList []GNBCUCPTNLAFailedToSetupItem `aper:"ub:MaxnoofTNLAssociations,optional,ext"`
	TransportLayerAddressInfo    *TransportLayerAddressInfo     `aper:"optional,ext"`
}

func (msg *GNBCUCPConfigurationUpdateAcknowledge) toIes() ([]E1APMessageIE, error) {
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
	if msg.CriticalityDiagnostics != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDCriticalityDiagnostics,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.CriticalityDiagnostics,
			})
		}
	}
	if len(msg.GNBCUCPTNLASetupList) > 0 {

		{

			tmp_GNBCUCPTNLASetupList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
				ext: false,
			}
			for i := 0; i < len(msg.GNBCUCPTNLASetupList); i++ {
				tmp_GNBCUCPTNLASetupList.Value = append(tmp_GNBCUCPTNLASetupList.Value, &msg.GNBCUCPTNLASetupList[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUCPTNLASetupList,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_GNBCUCPTNLASetupList,
			})
		}
	}
	if len(msg.GNBCUCPTNLAFailedToSetupList) > 0 {

		{

			tmp_GNBCUCPTNLAFailedToSetupList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
				ext: false,
			}
			for i := 0; i < len(msg.GNBCUCPTNLAFailedToSetupList); i++ {
				tmp_GNBCUCPTNLAFailedToSetupList.Value = append(tmp_GNBCUCPTNLAFailedToSetupList.Value, &msg.GNBCUCPTNLAFailedToSetupList[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUCPTNLAFailedToSetupList,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_GNBCUCPTNLAFailedToSetupList,
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
	return ies, nil
}

// Encode for GNBCUCPConfigurationUpdateAcknowledge: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for GNBCUCPConfigurationUpdateAcknowledge.
func (msg *GNBCUCPConfigurationUpdateAcknowledge) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GNBCUCPConfigurationUpdateAcknowledge: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := GNBCUCPConfigurationUpdateAcknowledgeDecoder{
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
	if err != nil {
		return
	}

	return
}

type GNBCUCPConfigurationUpdateAcknowledgeDecoder struct {
	msg      *GNBCUCPConfigurationUpdateAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*E1APMessageIE
}

func (decoder *GNBCUCPConfigurationUpdateAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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

	case ProtocolIEIDCriticalityDiagnostics:
		s.CriticalityDiagnostics = new(CriticalityDiagnostics)
		if err = s.CriticalityDiagnostics.Decode(r); err != nil {
			return fmt.Errorf("Decode CriticalityDiagnostics failed: %w", err)
		}

	case ProtocolIEIDGNBCUCPTNLASetupList:
		s.GNBCUCPTNLASetupList = new(GNBCUCPTNLASetupList)
		if err = s.GNBCUCPTNLASetupList.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUCPTNLASetupList failed: %w", err)
		}

	case ProtocolIEIDGNBCUCPTNLAFailedToSetupList:
		s.GNBCUCPTNLAFailedToSetupList = new(GNBCUCPTNLAFailedToSetupList)
		if err = s.GNBCUCPTNLAFailedToSetupList.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUCPTNLAFailedToSetupList failed: %w", err)
		}

	case ProtocolIEIDTransportLayerAddressInfo:
		s.TransportLayerAddressInfo = new(TransportLayerAddressInfo)
		if err = s.TransportLayerAddressInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode TransportLayerAddressInfo failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
	}
	return
}
