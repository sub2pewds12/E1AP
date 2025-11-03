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

// toIes transforms the SecurityIndication struct into a slice of E1APMessageIEs.
func (msg *SecurityIndication) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEID},
				Criticality: Criticality{Value: Criticality},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 2},
					ext:   true,
					Value: msg.IntegrityProtectionIndication.Value,
				},
			})
		}
	}
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEID},
				Criticality: Criticality{Value: Criticality},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 2},
					ext:   true,
					Value: msg.ConfidentialityProtectionIndication.Value,
				},
			})
		}
	}
	if msg.MaximumIPdatarate != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEID},
				Criticality: Criticality{Value: Criticality},
				Value:       msg.MaximumIPdatarate,
			})
		}
	}
	var err error
	return ies, err
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
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID]; !ok {
		err = fmt.Errorf("mandatory field IntegrityProtectionIndication is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID,
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID]; !ok {
		err = fmt.Errorf("mandatory field ConfidentialityProtectionIndication is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID,
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
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *SecurityIndicationDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
		return nil, fmt.Errorf("duplicated protocol IE ID %%d", ieId)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id {
	case ProtocolIEID:

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return nil, fmt.Errorf("Decode IntegrityProtectionIndication failed: %w", err)
			}
			msg.IntegrityProtectionIndication.Value = aper.Enumerated(val)
		}
	case ProtocolIEID:

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return nil, fmt.Errorf("Decode ConfidentialityProtectionIndication failed: %w", err)
			}
			msg.ConfidentialityProtectionIndication.Value = aper.Enumerated(val)
		}
	case ProtocolIEID:
		msg.MaximumIPdatarate = new(MaximumIPdatarate)
		if err = msg.MaximumIPdatarate.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode MaximumIPdatarate failed: %w", err)
		}
	case ProtocolIEID:
		msg.IEExtensions = new(ProtocolExtensionContainer)
		if err = msg.IEExtensions.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
