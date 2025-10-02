package e1ap_ies

// BearerContextSetupRequest From: 9_4_4_PDU_Definitions.txt:672
type BearerContextSetupRequest struct {
	GNBCUCPUEE1APID                       int64                           `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SecurityInformation                   SecurityInformation             `asn1:"mandatory,reject,ext"`
	UEDLAggregateMaximumBitRate           int64                           `asn1:"mandatory,reject,ext"`
	UEDLMaximumIntegrityProtectedDataRate *int64                          `asn1:"optional,reject,ext"`
	ServingPLMN                           []byte                          `asn1:"lb:3,ub:3,mandatory,ignore,ext"`
	ActivityNotificationLevel             ActivityNotificationLevel       `asn1:"mandatory,reject,ext"`
	UEInactivityTimer                     *int64                          `asn1:"optional,reject,ext"`
	BearerContextStatusChange             *BearerContextStatusChange      `asn1:"optional,reject,ext"`
	SystemBearerContextSetupRequest       SystemBearerContextSetupRequest `asn1:"mandatory,reject,ext"`
	RANUEID                               *[]byte                         `asn1:"lb:8,ub:8,optional,ignore,ext"`
	GNBDUID                               *int64                          `asn1:"lb:0,ub:68719476735,optional,ignore,ext"`
	TraceActivation                       *TraceActivation                `asn1:"optional,ignore,ext"`
	NPNContextInfo                        *NPNContextInfo                 `asn1:"optional,reject,ext"`
	ManagementBasedMDTPLMNList            []PLMNIdentity                  `asn1:"lb:1,ub:MaxnoofMDTPLMNs,optional,ignore,ext"`
	CHOInitiation                         *CHOInitiation                  `asn1:"optional,reject,ext"`
}
