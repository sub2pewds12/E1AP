package e1ap_ies

import "github.com/lvdund/ngap/aper"

// BearerContextSetupRequest From: 9_4_4_PDU_Definitions.txt:672
// ASN.1 Data Type: SEQUENCE
type BearerContextSetupRequest struct {
	GNBCUCPUEE1APID                       aper.Integer                    `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SecurityInformation                   SecurityInformation             `aper:"mandatory,reject,ext"`
	UEDLAggregateMaximumBitRate           aper.Integer                    `aper:"mandatory,reject,ext"`
	UEDLMaximumIntegrityProtectedDataRate *aper.Integer                   `aper:"optional,reject,ext"`
	ServingPLMN                           aper.OctetString                `aper:"lb:3,ub:3,mandatory,ignore,ext"`
	ActivityNotificationLevel             ActivityNotificationLevel       `aper:"mandatory,reject,ext"`
	UEInactivityTimer                     *aper.Integer                   `aper:"optional,reject,ext"`
	BearerContextStatusChange             *BearerContextStatusChange      `aper:"optional,reject,ext"`
	SystemBearerContextSetupRequest       SystemBearerContextSetupRequest `aper:"mandatory,reject,ext"`
	RANUEID                               *aper.OctetString               `aper:"lb:8,ub:8,optional,ignore,ext"`
	GNBDUID                               *aper.Integer                   `aper:"lb:0,ub:68719476735,optional,ignore,ext"`
	TraceActivation                       *TraceActivation                `aper:"optional,ignore,ext"`
	NPNContextInfo                        *NPNContextInfo                 `aper:"optional,reject,ext"`
	ManagementBasedMDTPLMNList            []PLMNIdentity                  `aper:"lb:1,ub:MaxnoofMDTPLMNs,optional,ignore,ext"`
	CHOInitiation                         *CHOInitiation                  `aper:"optional,reject,ext"`
}
