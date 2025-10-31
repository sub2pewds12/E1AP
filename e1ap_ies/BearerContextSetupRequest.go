package e1ap_ies

import (
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

// Encode function for BearerContextSetupRequest to be generated here.

// Decode function for BearerContextSetupRequest to be generated here.

// Decoder helper for BearerContextSetupRequest to be generated here.
