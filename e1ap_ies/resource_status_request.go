package e1ap_ies

// ResourceStatusRequest From: 9_4_4_PDU_Definitions.txt:1340
type ResourceStatusRequest struct {
	TransactionID         int64                 `asn1:"mandatory,reject,ext"`
	RegistrationRequest   RegistrationRequest   `asn1:"mandatory,reject,ext"`
	ReportCharacteristics *[]byte               `asn1:"lb:36,ub:36,conditional,reject,ext"`
	ReportingPeriodicity  *ReportingPeriodicity `asn1:"optional,reject,ext"`
}
