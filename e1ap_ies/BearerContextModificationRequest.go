package e1ap_ies

import (
	"fmt"
	"io"

	"asn1go/per"
)

// BearerContextModificationRequest is a generated SEQUENCE type.
type BearerContextModificationRequest struct {
	GNBCUCPUEE1APID                        GNBCUCPUEE1APID
	GNBCUUPUEE1APID                        GNBCUUPUEE1APID
	SecurityInformation                    *SecurityInformation
	UEDLAggregateMaximumBitRate            *BitRate
	UEDLMaximumIntegrityProtectedDataRate  *BitRate
	BearerContextStatusChange              *BearerContextStatusChange
	NewULTNLInformationRequired            *NewULTNLInformationRequired
	UEInactivityTimer                      *InactivityTimer
	DataDiscardRequired                    *DataDiscardRequired
	SystemBearerContextModificationRequest *SystemBearerContextModificationRequest
	RANUEID                                *RANUEID
	GNBDUID                                *GNBDUID
	ActivityNotificationLevel              *ActivityNotificationLevel
}

// toIes transforms the BearerContextModificationRequest struct into a slice of E1APMessageIEs.
func (msg *BearerContextModificationRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GNBCUCPUEE1APID,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GNBCUUPUEE1APID,
	})
	if msg.SecurityInformation != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDSecurityInformation},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SecurityInformation,
		})
	}
	if msg.UEDLAggregateMaximumBitRate != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDUEDLAggregateMaximumBitRate},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.UEDLAggregateMaximumBitRate,
		})
	}
	if msg.UEDLMaximumIntegrityProtectedDataRate != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.UEDLMaximumIntegrityProtectedDataRate,
		})
	}
	if msg.BearerContextStatusChange != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDBearerContextStatusChange},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.BearerContextStatusChange,
		})
	}
	if msg.NewULTNLInformationRequired != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDNewULTNLInformationRequired},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.NewULTNLInformationRequired,
		})
	}
	if msg.UEInactivityTimer != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDUEInactivityTimer},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.UEInactivityTimer,
		})
	}
	if msg.DataDiscardRequired != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDDataDiscardRequired},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DataDiscardRequired,
		})
	}
	if msg.SystemBearerContextModificationRequest != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDSystemBearerContextModificationRequest},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SystemBearerContextModificationRequest,
		})
	}
	if msg.RANUEID != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDRANUEID},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RANUEID,
		})
	}
	if msg.GNBDUID != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBDUID},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBDUID,
		})
	}
	if msg.ActivityNotificationLevel != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDActivityNotificationLevel},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ActivityNotificationLevel,
		})
	}
	return ies, nil
}

func (msg *BearerContextModificationRequest) EncodeWithEncoder(e *per.Encoder) (err error) {
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

func (msg *BearerContextModificationRequest) Encode(w io.Writer) error {
	e := per.NewEncoder(per.APER)
	if err := msg.EncodeWithEncoder(e); err != nil {
		return err
	}
	_, err := w.Write(e.Bytes())
	return err
}

// Decode implements the MessageUnmarshaller interface for BearerContextModificationRequest.
func (msg *BearerContextModificationRequest) Decode(data []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	r := per.NewDecoder(data, per.APER)
	return msg.DecodeFromDecoder(r)
}

func (msg *BearerContextModificationRequest) DecodeFromDecoder(r *per.Decoder) (diagList []CriticalityDiagnosticsIEItem, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("decode BearerContextModificationRequest failed: %w", err)
		}
	}()

	decoder := BearerContextModificationRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field GNBCUCPUEE1APID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field GNBCUUPUEE1APID is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type BearerContextModificationRequestDecoder struct {
	msg      *BearerContextModificationRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *BearerContextModificationRequestDecoder) decodeIE(r *per.Decoder) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			val, err := ieR.DecodeInteger(per.Unconstrained())
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPUEE1APID failed: %w", err)
			}
			msg.GNBCUCPUEE1APID.Value = val
		}
	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			val, err := ieR.DecodeInteger(per.Unconstrained())
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPUEE1APID failed: %w", err)
			}
			msg.GNBCUUPUEE1APID.Value = val
		}
	case ProtocolIEIDSecurityInformation:
		msg.SecurityInformation = new(SecurityInformation)

		if err = msg.SecurityInformation.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode SecurityInformation failed: %w", err)
		}
	case ProtocolIEIDUEDLAggregateMaximumBitRate:

		{
			val, err := ieR.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return nil, fmt.Errorf("decode UEDLAggregateMaximumBitRate failed: %w", err)
			}
			msg.UEDLAggregateMaximumBitRate = new(BitRate)
			msg.UEDLAggregateMaximumBitRate.Value = val
		}
	case ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate:

		{
			val, err := ieR.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return nil, fmt.Errorf("decode UEDLMaximumIntegrityProtectedDataRate failed: %w", err)
			}
			msg.UEDLMaximumIntegrityProtectedDataRate = new(BitRate)
			msg.UEDLMaximumIntegrityProtectedDataRate.Value = val
		}
	case ProtocolIEIDBearerContextStatusChange:
		msg.BearerContextStatusChange = new(BearerContextStatusChange)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode BearerContextStatusChange failed: %w", err)
			}
			msg.BearerContextStatusChange.Value = val
		}
	case ProtocolIEIDNewULTNLInformationRequired:
		msg.NewULTNLInformationRequired = new(NewULTNLInformationRequired)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode NewULTNLInformationRequired failed: %w", err)
			}
			msg.NewULTNLInformationRequired.Value = val
		}
	case ProtocolIEIDUEInactivityTimer:

		{
			val, err := ieR.DecodeInteger(per.ConstrainedExtensible(1, 7200))
			if err != nil {
				return nil, fmt.Errorf("decode UEInactivityTimer failed: %w", err)
			}
			msg.UEInactivityTimer = new(InactivityTimer)
			msg.UEInactivityTimer.Value = val
		}
	case ProtocolIEIDDataDiscardRequired:
		msg.DataDiscardRequired = new(DataDiscardRequired)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode DataDiscardRequired failed: %w", err)
			}
			msg.DataDiscardRequired.Value = val
		}
	case ProtocolIEIDSystemBearerContextModificationRequest:
		msg.SystemBearerContextModificationRequest = new(SystemBearerContextModificationRequest)

		if err = msg.SystemBearerContextModificationRequest.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode SystemBearerContextModificationRequest failed: %w", err)
		}
	case ProtocolIEIDRANUEID:
		msg.RANUEID = new(RANUEID)

		if err = msg.RANUEID.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode RANUEID failed: %w", err)
		}
	case ProtocolIEIDGNBDUID:

		{
			val, err := ieR.DecodeInteger(per.Constrained(0, 68719476735))
			if err != nil {
				return nil, fmt.Errorf("decode GNBDUID failed: %w", err)
			}
			msg.GNBDUID = new(GNBDUID)
			msg.GNBDUID.Value = val
		}
	case ProtocolIEIDActivityNotificationLevel:
		msg.ActivityNotificationLevel = new(ActivityNotificationLevel)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode ActivityNotificationLevel failed: %w", err)
			}
			msg.ActivityNotificationLevel.Value = val
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
