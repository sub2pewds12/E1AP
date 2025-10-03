package e1ap_ies

import "github.com/lvdund/ngap/aper"

// BearerContextModificationRequest From: 9_4_4_PDU_Definitions.txt:793
// ASN.1 Data Type: SEQUENCE
type BearerContextModificationRequest struct {
	GNBCUCPUEE1APID                        aper.Integer                            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                        aper.Integer                            `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SecurityInformation                    *SecurityInformation                    `aper:"optional,reject,ext"`
	UEDLAggregateMaximumBitRate            *aper.Integer                           `aper:"optional,reject,ext"`
	UEDLMaximumIntegrityProtectedDataRate  *aper.Integer                           `aper:"optional,reject,ext"`
	BearerContextStatusChange              *BearerContextStatusChange              `aper:"optional,reject,ext"`
	NewULTNLInformationRequired            *NewULTNLInformationRequired            `aper:"optional,reject,ext"`
	UEInactivityTimer                      *aper.Integer                           `aper:"optional,reject,ext"`
	DataDiscardRequired                    *DataDiscardRequired                    `aper:"optional,ignore,ext"`
	SystemBearerContextModificationRequest *SystemBearerContextModificationRequest `aper:"optional,reject,ext"`
	RANUEID                                *aper.OctetString                       `aper:"lb:8,ub:8,optional,ignore,ext"`
	GNBDUID                                *aper.Integer                           `aper:"lb:0,ub:68719476735,optional,ignore,ext"`
	ActivityNotificationLevel              *ActivityNotificationLevel              `aper:"optional,ignore,ext"`
}
