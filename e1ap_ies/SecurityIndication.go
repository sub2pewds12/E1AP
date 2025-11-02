package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SecurityIndication is a generated SEQUENCE type.
type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication       `aper:"mandatory,ext"`
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication `aper:"mandatory,ext"`
	MaximumIPdatarate                   *MaximumIPdatarate                  `aper:"optional,ext"`
	IEExtensions                        *ProtocolExtensionContainer         `aper:"optional,ext"`
}

func (msg *SecurityIndication) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID,
			Criticality: Criticality{Value: Criticality},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 2},
				ext:   true,
				Value: msg.IntegrityProtectionIndication.Value,
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID,
			Criticality: Criticality{Value: Criticality},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 2},
				ext:   true,
				Value: msg.ConfidentialityProtectionIndication.Value,
			},
		})
	}
	if msg.MaximumIPdatarate != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID,
				Criticality: Criticality{Value: Criticality},
				Value:       msg.MaximumIPdatarate,
			})
		}
	}
	return ies, nil
}

// Encode for SecurityIndication: Could not find associated procedure.

// Decode implements the aper.AperUnmarshaller interface for SecurityIndication.
func (msg *SecurityIndication) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("SecurityIndication: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := SecurityIndicationDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID]; !ok {
		err = fmt.Errorf("mandatory field IntegrityProtectionIndication is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID]; !ok {
		err = fmt.Errorf("mandatory field ConfidentialityProtectionIndication is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type SecurityIndicationDecoder struct {
	msg      *SecurityIndication
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*E1APMessageIE
}

func (decoder *SecurityIndicationDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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

	case ProtocolIEID:

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return fmt.Errorf("Decode IntegrityProtectionIndication failed: %w", err)
			}
			s.IntegrityProtectionIndication.Value = aper.Enumerated(val)
		}

	case ProtocolIEID:

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return fmt.Errorf("Decode ConfidentialityProtectionIndication failed: %w", err)
			}
			s.ConfidentialityProtectionIndication.Value = aper.Enumerated(val)
		}

	case ProtocolIEID:
		s.MaximumIPdatarate = new(MaximumIPdatarate)
		if err = s.MaximumIPdatarate.Decode(r); err != nil {
			return fmt.Errorf("Decode MaximumIPdatarate failed: %w", err)
		}

	case ProtocolIEID:
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
	}
	return
}
