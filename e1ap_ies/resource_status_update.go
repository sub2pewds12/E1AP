package e1ap_ies

// ResourceStatusUpdate From: 9_4_4_PDU_Definitions.txt:1401
type ResourceStatusUpdate struct {
	TransactionID                 int64                          `asn1:"mandatory,reject,ext"`
	TNLAvailableCapacityIndicator *TNLAvailableCapacityIndicator `asn1:"optional,ignore,ext"`
	HWCapacityIndicator           HWCapacityIndicator            `asn1:"mandatory,ignore,ext"`
}
