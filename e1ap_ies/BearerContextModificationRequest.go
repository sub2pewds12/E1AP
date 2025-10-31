package e1ap_ies

import (
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

// Encode function for BearerContextModificationRequest to be generated here.

// Decode function for BearerContextModificationRequest to be generated here.

// Decoder helper for BearerContextModificationRequest to be generated here.
