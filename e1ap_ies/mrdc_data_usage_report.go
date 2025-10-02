package e1ap_ies

// MRDCDataUsageReport From: 9_4_4_PDU_Definitions.txt:1242
type MRDCDataUsageReport struct {
	GNBCUCPUEE1APID                 int64                             `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                 int64                             `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	PDUSessionResourceDataUsageList []PDUSessionResourceDataUsageItem `asn1:"mandatory,ignore,ext"`
}
