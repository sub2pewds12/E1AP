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
	ManagementBasedMDTPLMNList            []PLMNIdentity                    `aper:"lb:1,ub:MaxnoofMDTPLMNs,optional,ext"`
	CHOInitiation                         *CHOInitiation                    `aper:"optional,ext"`
	AdditionalHandoverInfo                *AdditionalHandoverInfo           `aper:"optional,ext"`
	DirectForwardingPathAvailability      *DirectForwardingPathAvailability `aper:"optional,ext"`
	GNBCUUPUEE1APID                       *GNBCUUPUEE1APID                  `aper:"lb:0,ub:4294967295,optional,ext"`
}

func (msg *BearerContextSetupRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDGNBCUCPUEE1APID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(msg.GNBCUCPUEE1APID),
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDSecurityInformation,
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &msg.SecurityInformation,
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDUEDLAggregateMaximumBitRate,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
				ext:   true,
				Value: aper.Integer(msg.UEDLAggregateMaximumBitRate),
			},
		})
	}
	if msg.UEDLMaximumIntegrityProtectedDataRate != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
					ext:   true,
					Value: aper.Integer((*msg.UEDLMaximumIntegrityProtectedDataRate)),
				},
			})
		}
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDServingPLMN,
			Criticality: Criticality{Value: CriticalityIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 3, Ub: 3},
				ext:   false,
				Value: aper.OctetString(msg.ServingPLMN),
			},
		})
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDActivityNotificationLevel,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &ENUMERATED{
				c:     aper.Constraint{Lb: 0, Ub: 2},
				ext:   true,
				Value: msg.ActivityNotificationLevel.Value,
			},
		})
	}
	if msg.UEInactivityTimer != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDUEInactivityTimer,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 1, Ub: 7200},
					ext:   true,
					Value: aper.Integer((*msg.UEInactivityTimer)),
				},
			})
		}
	}
	if msg.BearerContextStatusChange != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDBearerContextStatusChange,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 1},
					ext:   true,
					Value: (*msg.BearerContextStatusChange).Value,
				},
			})
		}
	}

	{

		ies = append(ies, E1APMessageIE{
			Id:          ProtocolIEIDSystemBearerContextSetupRequest,
			Criticality: Criticality{Value: CriticalityReject},
			Value:       &msg.SystemBearerContextSetupRequest,
		})
	}
	if msg.RANUEID != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDRANUEID,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &OCTETSTRING{
					c:     aper.Constraint{Lb: 8, Ub: 8},
					ext:   false,
					Value: aper.OctetString((*msg.RANUEID)),
				},
			})
		}
	}
	if msg.GNBDUID != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBDUID,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 68719476735},
					ext:   false,
					Value: aper.Integer((*msg.GNBDUID)),
				},
			})
		}
	}
	if msg.TraceActivation != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDTraceActivation,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       msg.TraceActivation,
			})
		}
	}
	if msg.NPNContextInfo != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDNPNContextInfo,
				Criticality: Criticality{Value: CriticalityReject},
				Value:       msg.NPNContextInfo,
			})
		}
	}
	if len(msg.ManagementBasedMDTPLMNList) > 0 {

		{

			tmp_ManagementBasedMDTPLMNList := Sequence[aper.IE]{
				c:   aper.Constraint{Lb: 1, Ub: MaxnoofMDTPLMNs},
				ext: false,
			}
			for i := 0; i < len(msg.ManagementBasedMDTPLMNList); i++ {
				tmp_ManagementBasedMDTPLMNList.Value = append(tmp_ManagementBasedMDTPLMNList.Value, &msg.ManagementBasedMDTPLMNList[i])
			}
			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDManagementBasedMDTPLMNList,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value:       &tmp_ManagementBasedMDTPLMNList,
			})
		}
	}
	if msg.CHOInitiation != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDCHOInitiation,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   true,
					Value: (*msg.CHOInitiation).Value,
				},
			})
		}
	}
	if msg.AdditionalHandoverInfo != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDAdditionalHandoverInfo,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   true,
					Value: (*msg.AdditionalHandoverInfo).Value,
				},
			})
		}
	}
	if msg.DirectForwardingPathAvailability != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDDirectForwardingPathAvailability,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 1},
					ext:   true,
					Value: (*msg.DirectForwardingPathAvailability).Value,
				},
			})
		}
	}
	if msg.GNBCUUPUEE1APID != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDGNBCUUPUEE1APID,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: aper.Integer((*msg.GNBCUUPUEE1APID)),
				},
			})
		}
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for BearerContextSetupRequest.
func (msg *BearerContextSetupRequest) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert BearerContextSetupRequest to IEs: %w", err)
	}

	return EncodeInitiatingMessage(w, ProcedureCodeBearerContextSetup, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for BearerContextSetupRequest.
func (msg *BearerContextSetupRequest) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("BearerContextSetupRequest: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := BearerContextSetupRequestDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEIDGNBCUCPUEE1APID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUCPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDSecurityInformation]; !ok {
		err = fmt.Errorf("mandatory field SecurityInformation is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDSecurityInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDUEDLAggregateMaximumBitRate]; !ok {
		err = fmt.Errorf("mandatory field UEDLAggregateMaximumBitRate is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDUEDLAggregateMaximumBitRate},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDServingPLMN]; !ok {
		err = fmt.Errorf("mandatory field ServingPLMN is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDServingPLMN},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDActivityNotificationLevel]; !ok {
		err = fmt.Errorf("mandatory field ActivityNotificationLevel is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
			IEID:          ProtocolIEID{Value: ProtocolIEIDActivityNotificationLevel},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEIDSystemBearerContextSetupRequest]; !ok {
		err = fmt.Errorf("mandatory field SystemBearerContextSetupRequest is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
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
	list     map[aper.Integer]*E1APMessageIE
}

func (decoder *BearerContextSetupRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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

	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return fmt.Errorf("Decode GNBCUCPUEE1APID failed: %w", err)
			}
			s.GNBCUCPUEE1APID = GNBCUCPUEE1APID(val)
		}

	case ProtocolIEIDSecurityInformation:
		if err = s.SecurityInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode SecurityInformation failed: %w", err)
		}

	case ProtocolIEIDUEDLAggregateMaximumBitRate:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode UEDLAggregateMaximumBitRate failed: %w", err)
			}
			s.UEDLAggregateMaximumBitRate = BitRate(val)
		}

	case ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode UEDLMaximumIntegrityProtectedDataRate failed: %w", err)
			}
			tmp := BitRate(val)
			s.UEDLMaximumIntegrityProtectedDataRate = &tmp
		}

	case ProtocolIEIDServingPLMN:

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
				return fmt.Errorf("Decode ServingPLMN failed: %w", err)
			}
			s.ServingPLMN = PLMNIdentity(val)
		}

	case ProtocolIEIDActivityNotificationLevel:

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return fmt.Errorf("Decode ActivityNotificationLevel failed: %w", err)
			}
			s.ActivityNotificationLevel.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDUEInactivityTimer:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
				return fmt.Errorf("Decode UEInactivityTimer failed: %w", err)
			}
			tmp := InactivityTimer(val)
			s.UEInactivityTimer = &tmp
		}

	case ProtocolIEIDBearerContextStatusChange:
		s.BearerContextStatusChange = new(BearerContextStatusChange)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return fmt.Errorf("Decode BearerContextStatusChange failed: %w", err)
			}
			s.BearerContextStatusChange.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDSystemBearerContextSetupRequest:
		if err = s.SystemBearerContextSetupRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode SystemBearerContextSetupRequest failed: %w", err)
		}

	case ProtocolIEIDRANUEID:

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
				return fmt.Errorf("Decode RANUEID failed: %w", err)
			}
			tmp := RANUEID(val)
			s.RANUEID = &tmp
		}

	case ProtocolIEIDGNBDUID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false); err != nil {
				return fmt.Errorf("Decode GNBDUID failed: %w", err)
			}
			tmp := GNBDUID(val)
			s.GNBDUID = &tmp
		}

	case ProtocolIEIDTraceActivation:
		s.TraceActivation = new(TraceActivation)
		if err = s.TraceActivation.Decode(r); err != nil {
			return fmt.Errorf("Decode TraceActivation failed: %w", err)
		}

	case ProtocolIEIDNPNContextInfo:
		s.NPNContextInfo = new(NPNContextInfo)
		if err = s.NPNContextInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode NPNContextInfo failed: %w", err)
		}

	case ProtocolIEIDManagementBasedMDTPLMNList:
		s.ManagementBasedMDTPLMNList = new(MDTPLMNList)
		if err = s.ManagementBasedMDTPLMNList.Decode(r); err != nil {
			return fmt.Errorf("Decode ManagementBasedMDTPLMNList failed: %w", err)
		}

	case ProtocolIEIDCHOInitiation:
		s.CHOInitiation = new(CHOInitiation)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode CHOInitiation failed: %w", err)
			}
			s.CHOInitiation.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDAdditionalHandoverInfo:
		s.AdditionalHandoverInfo = new(AdditionalHandoverInfo)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode AdditionalHandoverInfo failed: %w", err)
			}
			s.AdditionalHandoverInfo.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDDirectForwardingPathAvailability:
		s.DirectForwardingPathAvailability = new(DirectForwardingPathAvailability)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return fmt.Errorf("Decode DirectForwardingPathAvailability failed: %w", err)
			}
			s.DirectForwardingPathAvailability.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
			}
			tmp := GNBCUUPUEE1APID(val)
			s.GNBCUUPUEE1APID = &tmp
		}

	default:
		// Handle unknown IEs based on criticality here, if needed.
	}
	return
}
