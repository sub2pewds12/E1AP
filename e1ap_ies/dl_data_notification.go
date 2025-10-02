package e1ap_ies

// DLDataNotification From: 9_4_4_PDU_Definitions.txt:1109
type DLDataNotification struct {
	GNBCUCPUEE1APID int64  `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID int64  `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	PPI             *int64 `asn1:"optional,ignore,ext"`
}
