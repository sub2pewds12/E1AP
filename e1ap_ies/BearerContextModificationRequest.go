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

// toIes transforms the BearerContextModificationRequest struct into a slice of E1APMessageIEs.
func (msg *BearerContextModificationRequest) toIes() ([]E1APMessageIE, error) {
	ies := make([]E1APMessageIE, 0)
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: msg.GNBCUCPUEE1APID.Value,
				},
			})
		}
	}
	{

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4294967295},
					ext:   false,
					Value: msg.GNBCUUPUEE1APID.Value,
				},
			})
		}
	}
	if msg.SecurityInformation != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDSecurityInformation},
				Criticality: Criticality{Value: CriticalityReject},
				Value:       msg.SecurityInformation,
			})
		}
	}
	if msg.UEDLAggregateMaximumBitRate != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDUEDLAggregateMaximumBitRate},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &INTEGER{
					c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
					ext:   true,
					Value: msg.UEDLAggregateMaximumBitRate.Value,
				},
			})
		}
	}
	if msg.UEDLMaximumIntegrityProtectedDataRate != nil {

		{

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
	}
	if msg.BearerContextStatusChange != nil {

		{

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
	}
	if msg.NewULTNLInformationRequired != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDNewULTNLInformationRequired},
				Criticality: Criticality{Value: CriticalityReject},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   true,
					Value: msg.NewULTNLInformationRequired.Value,
				},
			})
		}
	}
	if msg.UEInactivityTimer != nil {

		{

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
	}
	if msg.DataDiscardRequired != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDDataDiscardRequired},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 0},
					ext:   true,
					Value: msg.DataDiscardRequired.Value,
				},
			})
		}
	}
	if msg.SystemBearerContextModificationRequest != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDSystemBearerContextModificationRequest},
				Criticality: Criticality{Value: CriticalityReject},
				Value:       msg.SystemBearerContextModificationRequest,
			})
		}
	}
	if msg.RANUEID != nil {

		{

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
	}
	if msg.GNBDUID != nil {

		{

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
	}
	if msg.ActivityNotificationLevel != nil {

		{

			ies = append(ies, E1APMessageIE{
				Id:          ProtocolIEID{Value: ProtocolIEIDActivityNotificationLevel},
				Criticality: Criticality{Value: CriticalityIgnore},
				Value: &ENUMERATED{
					c:     aper.Constraint{Lb: 0, Ub: 2},
					ext:   true,
					Value: msg.ActivityNotificationLevel.Value,
				},
			})
		}
	}
	var err error
	return ies, err
}

// Encode implements the aper.AperMarshaller interface for BearerContextModificationRequest.
func (msg *BearerContextModificationRequest) Encode(w io.Writer) error {
	ies, err := msg.toIes()
	if err != nil {
		return fmt.Errorf("could not convert BearerContextModificationRequest to IEs: %w", err)
	}

	return encodeMessage(w, E1apPduInitiatingMessage, ProcedureCode{Value: ProcedureCodeBearerContextModification}, Criticality{Value: CriticalityReject}, ies)
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
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}

	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}

	// After decoding all present IEs, validate that mandatory ones were found.

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID}]; !ok {
		err = fmt.Errorf("mandatory field GNBCUCPUEE1APID is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: CriticalityReject},
			IEID:          ProtocolIEID{Value: ProtocolIEIDGNBCUCPUEE1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
	}

	if _, ok := decoder.list[ProtocolIEID{Value: ProtocolIEIDGNBCUUPUEE1APID}]; !ok {
		err = fmt.Errorf("mandatory field GNBCUUPUEE1APID is missing")
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

func (decoder *BearerContextModificationRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {
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
		return nil, fmt.Errorf("duplicated protocol IE ID %d", ieId.Value)
	}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id.Value {
	case ProtocolIEIDGNBCUCPUEE1APID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUCPUEE1APID failed: %w", err)
			}
			msg.GNBCUCPUEE1APID.Value = aper.Integer(val)
		}
	case ProtocolIEIDGNBCUUPUEE1APID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
			}
			msg.GNBCUUPUEE1APID.Value = aper.Integer(val)
		}
	case ProtocolIEIDSecurityInformation:
		msg.SecurityInformation = new(SecurityInformation)
		if err = msg.SecurityInformation.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode SecurityInformation failed: %w", err)
		}
	case ProtocolIEIDUEDLAggregateMaximumBitRate:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return nil, fmt.Errorf("Decode UEDLAggregateMaximumBitRate failed: %w", err)
			}
			msg.UEDLAggregateMaximumBitRate = new(BitRate)
			msg.UEDLAggregateMaximumBitRate.Value = aper.Integer(val)
		}
	case ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return nil, fmt.Errorf("Decode UEDLMaximumIntegrityProtectedDataRate failed: %w", err)
			}
			msg.UEDLMaximumIntegrityProtectedDataRate = new(BitRate)
			msg.UEDLMaximumIntegrityProtectedDataRate.Value = aper.Integer(val)
		}
	case ProtocolIEIDBearerContextStatusChange:
		msg.BearerContextStatusChange = new(BearerContextStatusChange)

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return nil, fmt.Errorf("Decode BearerContextStatusChange failed: %w", err)
			}
			msg.BearerContextStatusChange.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDNewULTNLInformationRequired:
		msg.NewULTNLInformationRequired = new(NewULTNLInformationRequired)

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return nil, fmt.Errorf("Decode NewULTNLInformationRequired failed: %w", err)
			}
			msg.NewULTNLInformationRequired.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDUEInactivityTimer:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
				return nil, fmt.Errorf("Decode UEInactivityTimer failed: %w", err)
			}
			msg.UEInactivityTimer = new(InactivityTimer)
			msg.UEInactivityTimer.Value = aper.Integer(val)
		}
	case ProtocolIEIDDataDiscardRequired:
		msg.DataDiscardRequired = new(DataDiscardRequired)

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return nil, fmt.Errorf("Decode DataDiscardRequired failed: %w", err)
			}
			msg.DataDiscardRequired.Value = aper.Enumerated(val)
		}
	case ProtocolIEIDSystemBearerContextModificationRequest:
		msg.SystemBearerContextModificationRequest = new(SystemBearerContextModificationRequest)
		if err = msg.SystemBearerContextModificationRequest.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode SystemBearerContextModificationRequest failed: %w", err)
		}
	case ProtocolIEIDRANUEID:
		msg.RANUEID = new(RANUEID)
		if err = msg.RANUEID.Decode(ieR); err != nil {
			return nil, fmt.Errorf("Decode RANUEID failed: %w", err)
		}
	case ProtocolIEIDGNBDUID:

		{
			var val int64
			if val, err = ieR.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false); err != nil {
				return nil, fmt.Errorf("Decode GNBDUID failed: %w", err)
			}
			msg.GNBDUID = new(GNBDUID)
			msg.GNBDUID.Value = aper.Integer(val)
		}
	case ProtocolIEIDActivityNotificationLevel:
		msg.ActivityNotificationLevel = new(ActivityNotificationLevel)

		{
			var val uint64
			if val, err = ieR.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
				return nil, fmt.Errorf("Decode ActivityNotificationLevel failed: %w", err)
			}
			msg.ActivityNotificationLevel.Value = aper.Enumerated(val)
		}
	default:
		// Handle unknown IEs based on criticality here, if needed.
		// For now, we'll just ignore them.

	}
	return msgIe, nil // Return the populated msgIe and a nil error
}
