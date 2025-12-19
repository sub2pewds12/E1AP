package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextSetupRequest is a generated SEQUENCE type.
type BearerContextSetupRequest struct {
	GNBCUCPUEE1APID                       GNBCUCPUEE1APID                   `aper:"lb:0,ub:4294967295,mandatory,ext"`
	SecurityInformation                   SecurityInformation               `aper:"mandatory,ext"`
	UEDLAggregateMaximumBitRate           BitRate                           `aper:"lb:0,ub:4000000000000,mandatory,ext"`
	UEDLMaximumIntegrityProtectedDataRate *BitRate                          `aper:"lb:0,ub:4000000000000,optional,ext"`
	ServingPLMN                           PLMNIdentity                      `aper:"lb:3,ub:3,mandatory,ext"`
	ActivityNotificationLevel             ActivityNotificationLevel         `aper:"mandatory,ext"`
	UEInactivityTimer                     *InactivityTimer                  `aper:"lb:1,ub:7200,optional,ext"`
	BearerContextStatusChange             *BearerContextStatusChange        `aper:"optional,ext"`
	SystemBearerContextSetupRequest       SystemBearerContextSetupRequest   `aper:"mandatory,ext"`
	RANUEID                               *RANUEID                          `aper:"lb:8,ub:8,optional,ext"`
	GNBDUID                               *GNBDUID                          `aper:"lb:0,ub:68719476735,optional,ext"`
	TraceActivation                       *TraceActivation                  `aper:"optional,ext"`
	NPNContextInfo                        *NPNContextInfo                   `aper:"optional,ext"`
	ManagementBasedMDTPLMNList            *MDTPLMNList                      `aper:"lb:1,ub:MaxnoofMDTPLMNs,optional,ext"`
	CHOInitiation                         *CHOInitiation                    `aper:"optional,ext"`
	AdditionalHandoverInfo                *AdditionalHandoverInfo           `aper:"optional,ext"`
	DirectForwardingPathAvailability      *DirectForwardingPathAvailability `aper:"optional,ext"`
	GNBCUUPUEE1APID                       *GNBCUUPUEE1APID                  `aper:"lb:0,ub:4294967295,optional,ext"`
}

// toIes transforms the BearerContextSetupRequest struct into a slice of E1APMessageIEs.
func (msg *BearerContextSetupRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
		Criticality: Criticality{Value: CriticalityReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: msg.GNBCUCPUEE1APID.Value,
		},
	})

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDSecurityInformation},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SecurityInformation,
	})

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDUEDLAggregateMaximumBitRate},
		Criticality: Criticality{Value: CriticalityReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
			ext:   true,
			Value: msg.UEDLAggregateMaximumBitRate.Value,
		},
	})
	if msg.UEDLMaximumIntegrityProtectedDataRate != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate},
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
				ext:   true,
				Value: msg.UEDLMaximumIntegrityProtectedDataRate.Value,
			},
		})
	}

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDServingPLMN},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 3, Ub: 3},
			ext:   false,
			Value: msg.ServingPLMN.Value,
		},
	})

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDActivityNotificationLevel},
		Criticality: Criticality{Value: CriticalityReject},
		Value: &ENUMERATED{
			c:     aper.Constraint{Lb: 0, Ub: 2},
			ext:   true,
			Value: msg.ActivityNotificationLevel.Value,
		},
	})
	if msg.UEInactivityTimer != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDUEInactivityTimer},
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 7200},
				ext:   true,
				Value: msg.UEInactivityTimer.Value,
			},
		})
	}
	if msg.BearerContextStatusChange != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDBearerContextStatusChange},
			Criticality: Criticality{Value: CriticalityReject},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 1},
				ext:   true,
				Value: msg.BearerContextStatusChange.Value,
			},
		})
	}

	ies = append(ies, E1APMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEIDSystemBearerContextSetupRequest},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SystemBearerContextSetupRequest,
	})
	if msg.RANUEID != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDRANUEID},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 8, Ub: 8},
				ext:   false,
				Value: msg.RANUEID.Value,
			},
		})
	}
	if msg.GNBDUID != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDGNBDUID},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 68719476735},
				ext:   false,
				Value: msg.GNBDUID.Value,
			},
		})
	}
	if msg.TraceActivation != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDTraceActivation},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TraceActivation,
		})
	}
	if msg.NPNContextInfo != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDNPNContextInfo},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.NPNContextInfo,
		})
	}
	if msg.ManagementBasedMDTPLMNList != nil {

		tmpManagementBasedMDTPLMNList := Sequence[aper.IE]{
			c:   aper.Constraint{Lb: 1, Ub: MaxnoofMDTPLMNs},
			ext: false,
		}

		for _, item := range msg.ManagementBasedMDTPLMNList.Value {
			wrapped_item := &OCTETSTRING{
				c:     aper.Constraint{Lb: 3, Ub: 3},
				ext:   false,
				Value: item.Value,
			}
			tmpManagementBasedMDTPLMNList.Value = append(tmpManagementBasedMDTPLMNList.Value, wrapped_item)
		}

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDManagementBasedMDTPLMNList},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &tmpManagementBasedMDTPLMNList,
		})
	}
	if msg.CHOInitiation != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDCHOInitiation},
			Criticality: Criticality{Value: CriticalityReject},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   true,
				Value: msg.CHOInitiation.Value,
			},
		})
	}
	if msg.AdditionalHandoverInfo != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDAdditionalHandoverInfo},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   true,
				Value: msg.AdditionalHandoverInfo.Value,
			},
		})
	}
	if msg.DirectForwardingPathAvailability != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDDirectForwardingPathAvailability},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 1},
				ext:   true,
				Value: msg.DirectForwardingPathAvailability.Value,
			},
		})
	}
	if msg.GNBCUUPUEE1APID != nil {

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: msg.GNBCUUPUEE1APID.Value,
			},
		})
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for BearerContextSetupRequest.
func (msg *BearerContextSetupRequest) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert BearerContextSetupRequest to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodeBearerContextSetup}, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for BearerContextSetupRequest.
func (msg *BearerContextSetupRequest) Decode(buf []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("decode BearerContextSetupRequest failed: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := BearerContextSetupRequestDecoder{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf(decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
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

func (decoder *BearerContextSetupRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUCPUEE1APID failed: %w", err)
			}
			msg.GNBCUCPUEE1APID.Value = aper.Integer(val)
		}
	case ProtocolIEIDSecurityInformation:
		if err = msg.SecurityInformation.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode SecurityInformation failed: %w", err)
		}
	case ProtocolIEIDUEDLAggregateMaximumBitRate:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
			if err != nil {
				return nil, fmt.Errorf("decode UEDLAggregateMaximumBitRate failed: %w", err)
			}
			msg.UEDLAggregateMaximumBitRate.Value = aper.Integer(val)
		}
	case ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
			if err != nil {
				return nil, fmt.Errorf("decode UEDLMaximumIntegrityProtectedDataRate failed: %w", err)
			}
			msg.UEDLMaximumIntegrityProtectedDataRate = new(BitRate)
			msg.UEDLMaximumIntegrityProtectedDataRate.Value = aper.Integer(val)
		}
	case ProtocolIEIDServingPLMN:
		if err = msg.ServingPLMN.Decode(ieR); err != nil {
			return nil, fmt.Errorf("decode ServingPLMN failed: %w", err)
		}
	case ProtocolIEIDActivityNotificationLevel:

		{
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
			if err != nil {
				return nil, fmt.Errorf("decode ActivityNotificationLevel failed: %w", err)
			}
			msg.ActivityNotificationLevel.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDUEInactivityTimer:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7200}, true)
			if err != nil {
				return nil, fmt.Errorf("decode UEInactivityTimer failed: %w", err)
			}
			msg.UEInactivityTimer = new(InactivityTimer)
			msg.UEInactivityTimer.Value = aper.Integer(val)
		}
	case ProtocolIEIDBearerContextStatusChange:
		msg.BearerContextStatusChange = new(BearerContextStatusChange)

		{
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
			if err != nil {
				return nil, fmt.Errorf("decode BearerContextStatusChange failed: %w", err)
			}
			msg.BearerContextStatusChange.Value = aper.Enumerated(val)
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
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBDUID failed: %w", err)
			}
			msg.GNBDUID = new(GNBDUID)
			msg.GNBDUID.Value = aper.Integer(val)
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

		{
			itemDecoder := func(r *aper.AperReader) (*PLMNIdentity, error) {

				item := new(PLMNIdentity)
				if err := item.Decode(r); err != nil {
					return nil, err
				}
				return item, nil
			}
			decodedItems, err := aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{Lb: 1, Ub: MaxnoofMDTPLMNs}, false)
			if err != nil {
				return nil, fmt.Errorf("decode ManagementBasedMDTPLMNList failed: %w", err)
			}

			msg.ManagementBasedMDTPLMNList = new(MDTPLMNList)
			msg.ManagementBasedMDTPLMNList.Value = decodedItems
		}
	case ProtocolIEIDCHOInitiation:
		msg.CHOInitiation = new(CHOInitiation)

		{
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
			if err != nil {
				return nil, fmt.Errorf("decode CHOInitiation failed: %w", err)
			}
			msg.CHOInitiation.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDAdditionalHandoverInfo:
		msg.AdditionalHandoverInfo = new(AdditionalHandoverInfo)

		{
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
			if err != nil {
				return nil, fmt.Errorf("decode AdditionalHandoverInfo failed: %w", err)
			}
			msg.AdditionalHandoverInfo.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDDirectForwardingPathAvailability:
		msg.DirectForwardingPathAvailability = new(DirectForwardingPathAvailability)

		{
			val, err := ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
			if err != nil {
				return nil, fmt.Errorf("decode DirectForwardingPathAvailability failed: %w", err)
			}
			msg.DirectForwardingPathAvailability.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			val, err := ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
			if err != nil {
				return nil, fmt.Errorf("decode GNBCUUPUEE1APID failed: %w", err)
			}
			msg.GNBCUUPUEE1APID = new(GNBCUUPUEE1APID)
			msg.GNBCUUPUEE1APID.Value = aper.Integer(val)
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
