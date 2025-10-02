package e1ap_ies

// EarlyForwardingSNTransfer From: 9_4_4_PDU_Definitions.txt:1495
type EarlyForwardingSNTransfer struct {
	GNBCUCPUEE1APID                  int64                              `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                  int64                              `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	DRBsSubjectToEarlyForwardingList []DRBsSubjectToEarlyForwardingItem `asn1:"mandatory,reject,ext"`
}
