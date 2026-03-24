package e1ap_ies

import (
	"fmt"
	"io"

	"asn1go/per"
)

// BearerContextSetupRequest is a generated SEQUENCE type.
type BearerContextSetupRequest struct {
	GNBCUCPUEE1APID                       GNBCUCPUEE1APID
	SecurityInformation                   SecurityInformation
	UEDLAggregateMaximumBitRate           BitRate
	UEDLMaximumIntegrityProtectedDataRate *BitRate
	ServingPLMN                           PLMNIdentity
	ActivityNotificationLevel             ActivityNotificationLevel
	UEInactivityTimer                     *InactivityTimer
	BearerContextStatusChange             *BearerContextStatusChange
	SystemBearerContextSetupRequest       SystemBearerContextSetupRequest
	RANUEID                               *RANUEID
	GNBDUID                               *GNBDUID
	TraceActivation                       *TraceActivation
	NPNContextInfo                        *NPNContextInfo
	ManagementBasedMDTPLMNList            *MDTPLMNList
	CHOInitiation                         *CHOInitiation
	AdditionalHandoverInfo                *AdditionalHandoverInfo
	DirectForwardingPathAvailability      *DirectForwardingPathAvailability
	GNBCUUPUEE1APID                       *GNBCUUPUEE1APID
}

// toIes transforms the BearerContextSetupRequest struct into a slice of E1APMessageIEs.
func (msg *BearerContextSetupRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GNBCUCPUEE1APID,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDSecurityInformation},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SecurityInformation,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDUEDLAggregateMaximumBitRate},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.UEDLAggregateMaximumBitRate,
	})
	if msg.UEDLMaximumIntegrityProtectedDataRate != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.UEDLMaximumIntegrityProtectedDataRate,
		})
	}

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDServingPLMN},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.ServingPLMN,
	})

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDActivityNotificationLevel},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.ActivityNotificationLevel,
	})
	if msg.UEInactivityTimer != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDUEInactivityTimer},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.UEInactivityTimer,
		})
	}
	if msg.BearerContextStatusChange != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDBearerContextStatusChange},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.BearerContextStatusChange,
		})
	}

	ies = append(ies, E1APMessageIE{
		ID:          ProtocolIEID{Value: ProtocolIEIDSystemBearerContextSetupRequest},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SystemBearerContextSetupRequest,
	})
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
	if msg.TraceActivation != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDTraceActivation},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TraceActivation,
		})
	}
	if msg.NPNContextInfo != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDNPNContextInfo},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.NPNContextInfo,
		})
	}
	if msg.ManagementBasedMDTPLMNList != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDManagementBasedMDTPLMNList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ManagementBasedMDTPLMNList,
		})
	}
	if msg.CHOInitiation != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDCHOInitiation},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CHOInitiation,
		})
	}
	if msg.AdditionalHandoverInfo != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDAdditionalHandoverInfo},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.AdditionalHandoverInfo,
		})
	}
	if msg.DirectForwardingPathAvailability != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDDirectForwardingPathAvailability},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DirectForwardingPathAvailability,
		})
	}
	if msg.GNBCUUPUEE1APID != nil {

		ies = append(ies, E1APMessageIE{
			ID:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.GNBCUUPUEE1APID,
		})
	}
	return ies, nil
}

func (msg *BearerContextSetupRequest) EncodeWithEncoder(e *per.Encoder) (err error) {
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

func (msg *BearerContextSetupRequest) Encode(w io.Writer) error {
	e := per.NewEncoder(per.APER)
	if err := msg.EncodeWithEncoder(e); err != nil {
		return err
	}
	_, err := w.Write(e.Bytes())
	return err
}

// Decode implements the MessageUnmarshaller interface for BearerContextSetupRequest.
func (msg *BearerContextSetupRequest) Decode(data []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	r := per.NewDecoder(data, per.APER)
	return msg.DecodeFromDecoder(r)
}

func (msg *BearerContextSetupRequest) DecodeFromDecoder(r *per.Decoder) (diagList []CriticalityDiagnosticsIEItem, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("decode BearerContextSetupRequest failed: %w", err)
		}
	}()

	decoder := BearerContextSetupRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDSecurityInformation}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field SecurityInformation is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDSecurityInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDUEDLAggregateMaximumBitRate}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field UEDLAggregateMaximumBitRate is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDUEDLAggregateMaximumBitRate},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDServingPLMN}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field ServingPLMN is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDServingPLMN},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDActivityNotificationLevel}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field ActivityNotificationLevel is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDActivityNotificationLevel},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDSystemBearerContextSetupRequest}]; !ok {
		if err == nil {
			err = fmt.Errorf("mandatory field SystemBearerContextSetupRequest is missing")
		}
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDSystemBearerContextSetupRequest},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}
	if err != nil {
		return
	}

	return
}

type BearerContextSetupRequestDecoder struct {
	msg      *BearerContextSetupRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}

func (decoder *BearerContextSetupRequestDecoder) decodeIE(r *per.Decoder) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDSecurityInformation:

		if err = msg.SecurityInformation.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode SecurityInformation failed: %w", err)
		}
	case ProtocolIEIDUEDLAggregateMaximumBitRate:

		{
			val, err := ieR.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return nil, fmt.Errorf("decode UEDLAggregateMaximumBitRate failed: %w", err)
			}
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
	case ProtocolIEIDServingPLMN:

		if err = msg.ServingPLMN.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode ServingPLMN failed: %w", err)
		}
	case ProtocolIEIDActivityNotificationLevel:

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode ActivityNotificationLevel failed: %w", err)
			}
			msg.ActivityNotificationLevel.Value = val
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
	case ProtocolIEIDSystemBearerContextSetupRequest:

		if err = msg.SystemBearerContextSetupRequest.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode SystemBearerContextSetupRequest failed: %w", err)
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
	case ProtocolIEIDTraceActivation:
		msg.TraceActivation = new(TraceActivation)

		if err = msg.TraceActivation.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode TraceActivation failed: %w", err)
		}
	case ProtocolIEIDNPNContextInfo:
		msg.NPNContextInfo = new(NPNContextInfo)

		if err = msg.NPNContextInfo.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode NPNContextInfo failed: %w", err)
		}
	case ProtocolIEIDManagementBasedMDTPLMNList:
		msg.ManagementBasedMDTPLMNList = new(MDTPLMNList)

		{
			itemDecoder := func(r *per.Decoder) (*PLMNIdentity, error) {
				item := new(PLMNIdentity)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}

			c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxnoofMDTPLMNs)}
			length, err := ieR.DecodeLengthDeterminant(c)
			if err != nil {
				return nil, fmt.Errorf("decode struct list length failed: %w", err)
			}
			for i := int64(0); i < length; i++ {
				item, err := itemDecoder(ieR)
				if err != nil {
					return nil, fmt.Errorf("decode item failed: %w", err)
				}
				msg.ManagementBasedMDTPLMNList.Value = append(msg.ManagementBasedMDTPLMNList.Value, *item)
			}
		}
	case ProtocolIEIDCHOInitiation:
		msg.CHOInitiation = new(CHOInitiation)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode CHOInitiation failed: %w", err)
			}
			msg.CHOInitiation.Value = val
		}
	case ProtocolIEIDAdditionalHandoverInfo:
		msg.AdditionalHandoverInfo = new(AdditionalHandoverInfo)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode AdditionalHandoverInfo failed: %w", err)
			}
			msg.AdditionalHandoverInfo.Value = val
		}
	case ProtocolIEIDDirectForwardingPathAvailability:
		msg.DirectForwardingPathAvailability = new(DirectForwardingPathAvailability)

		{
			c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2)}
			val, err := ieR.DecodeEnumerated(c)
			if err != nil {
				return nil, fmt.Errorf("decode DirectForwardingPathAvailability failed: %w", err)
			}
			msg.DirectForwardingPathAvailability.Value = val
		}
	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			val, err := ieR.DecodeInteger(per.Unconstrained())
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPUEE1APID failed: %w", err)
			}
			msg.GNBCUUPUEE1APID = new(GNBCUUPUEE1APID)
			msg.GNBCUUPUEE1APID.Value = val
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
