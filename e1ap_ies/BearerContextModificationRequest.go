package e1ap_ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextModificationRequest is a generated SEQUENCE type.
type BearerContextModificationRequest struct {
	GNBCUCPUEE1APID                        GNBCUCPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                        GNBCUUPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	SecurityInformation                    *SecurityInformation                    `aper:"optional,ext"`
	UEDLAggregateMaximumBitRate            *BitRate                                `aper:"lb:0,ub:4000000000000,optional,ext"`
	UEDLMaximumIntegrityProtectedDataRate  *BitRate                                `aper:"lb:0,ub:4000000000000,optional,ext"`
	BearerContextStatusChange              *BearerContextStatusChange              `aper:"optional,ext"`
	NewULTNLInformationRequired            *NewULTNLInformationRequired            `aper:"optional,ext"`
	UEInactivityTimer                      *InactivityTimer                        `aper:"lb:1,ub:7200,optional,ext"`
	DataDiscardRequired                    *DataDiscardRequired                    `aper:"optional,ext"`
	SystemBearerContextModificationRequest *SystemBearerContextModificationRequest `aper:"optional,ext"`
	RANUEID                                *RANUEID                                `aper:"lb:8,ub:8,optional,ext"`
	GNBDUID                                *GNBDUID                                `aper:"lb:0,ub:68719476735,optional,ext"`
	ActivityNotificationLevel              *ActivityNotificationLevel              `aper:"optional,ext"`
}

func (msg *BearerContextModificationRequest) toIes() ([]E1APMessageIE, error) {
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
			Id:          ProtocolIEIDGNBCUUPUEE1APID,
			Criticality: Criticality{Value: CriticalityReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(msg.GNBCUUPUEE1APID),
			},
		})
	}
	if msg.SecurityInformation != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDSecurityInformation,
				Criticality: Criticality{Value: CriticalityReject},
				Value:       msg.SecurityInformation,
			})
		}
	}
	if msg.UEDLAggregateMaximumBitRate != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDUEDLAggregateMaximumBitRate,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
					ext:   true,
					Value: aper.Integer((*msg.UEDLAggregateMaximumBitRate)),
				},
			})
		}
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
	if msg.NewULTNLInformationRequired != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDNewULTNLInformationRequired,
				Criticality: Criticality{Value: CriticalityReject},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   true,
					Value: (*msg.NewULTNLInformationRequired).Value,
				},
			})
		}
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
	if msg.DataDiscardRequired != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDDataDiscardRequired,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   true,
					Value: (*msg.DataDiscardRequired).Value,
				},
			})
		}
	}
	if msg.SystemBearerContextModificationRequest != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDSystemBearerContextModificationRequest,
				Criticality: Criticality{Value: CriticalityReject},
				Value:       msg.SystemBearerContextModificationRequest,
			})
		}
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
	if msg.ActivityNotificationLevel != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEIDActivityNotificationLevel,
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 2},
					ext:   true,
					Value: (*msg.ActivityNotificationLevel).Value,
				},
			})
		}
	}
	return ies, nil
}

// Encode implements the aper.AperMarshaller interface for BearerContextModificationRequest.
func (msg *BearerContextModificationRequest) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert BearerContextModificationRequest to IEs: %w", err)
	}

	return EncodeInitiatingMessage(w, ProcedureCodeBearerContextModification, Criticality{Value: CriticalityReject}, ies)
}

// Decode implements the aper.AperUnmarshaller interface for BearerContextModificationRequest.
func (msg *BearerContextModificationRequest) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("BearerContextModificationRequest: %w", err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(buf))

	decoder := BearerContextModificationRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEIDGNBCUUPUEE1APID]; !ok {
		err = fmt.Errorf("mandatory field GNBCUUPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject}, // Or from IE spec
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
	list     map[aper.Integer]*E1APMessageIE
}

func (decoder *BearerContextModificationRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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

	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
			}
			s.GNBCUUPUEE1APID = GNBCUUPUEE1APID(val)
		}

	case ProtocolIEIDSecurityInformation:
		s.SecurityInformation = new(SecurityInformation)
		if err = s.SecurityInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode SecurityInformation failed: %w", err)
		}

	case ProtocolIEIDUEDLAggregateMaximumBitRate:

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode UEDLAggregateMaximumBitRate failed: %w", err)
			}
			tmp := BitRate(val)
			s.UEDLAggregateMaximumBitRate = &tmp
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

	case ProtocolIEIDBearerContextStatusChange:
		s.BearerContextStatusChange = new(BearerContextStatusChange)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return fmt.Errorf("Decode BearerContextStatusChange failed: %w", err)
			}
			s.BearerContextStatusChange.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDNewULTNLInformationRequired:
		s.NewULTNLInformationRequired = new(NewULTNLInformationRequired)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode NewULTNLInformationRequired failed: %w", err)
			}
			s.NewULTNLInformationRequired.Value = aper.Enumerated(val)
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

	case ProtocolIEIDDataDiscardRequired:
		s.DataDiscardRequired = new(DataDiscardRequired)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode DataDiscardRequired failed: %w", err)
			}
			s.DataDiscardRequired.Value = aper.Enumerated(val)
		}

	case ProtocolIEIDSystemBearerContextModificationRequest:
		s.SystemBearerContextModificationRequest = new(SystemBearerContextModificationRequest)
		if err = s.SystemBearerContextModificationRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode SystemBearerContextModificationRequest failed: %w", err)
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

	case ProtocolIEIDActivityNotificationLevel:
		s.ActivityNotificationLevel = new(ActivityNotificationLevel)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return fmt.Errorf("Decode ActivityNotificationLevel failed: %w", err)
			}
			s.ActivityNotificationLevel.Value = aper.Enumerated(val)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
	}
	return
}
