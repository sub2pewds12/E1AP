package e1ap_ies

// BearerContextModificationRequest From: 9_4_4_PDU_Definitions.txt:793
type BearerContextModificationRequest struct {
	GNBCUCPUEE1APID                        int64                                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                        int64                                   `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SecurityInformation                    *SecurityInformation                    `asn1:"optional,reject,ext"`
	UEDLAggregateMaximumBitRate            *int64                                  `asn1:"optional,reject,ext"`
	UEDLMaximumIntegrityProtectedDataRate  *int64                                  `asn1:"optional,reject,ext"`
	BearerContextStatusChange              *BearerContextStatusChange              `asn1:"optional,reject,ext"`
	NewULTNLInformationRequired            *NewULTNLInformationRequired            `asn1:"optional,reject,ext"`
	UEInactivityTimer                      *int64                                  `asn1:"optional,reject,ext"`
	DataDiscardRequired                    *DataDiscardRequired                    `asn1:"optional,ignore,ext"`
	SystemBearerContextModificationRequest *SystemBearerContextModificationRequest `asn1:"optional,reject,ext"`
	RANUEID                                *[]byte                                 `asn1:"lb:8,ub:8,optional,ignore,ext"`
	GNBDUID                                *int64                                  `asn1:"lb:0,ub:68719476735,optional,ignore,ext"`
	ActivityNotificationLevel              *ActivityNotificationLevel              `asn1:"optional,ignore,ext"`
}
