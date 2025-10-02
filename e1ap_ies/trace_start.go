package e1ap_ies

// TraceStart From: 9_4_4_PDU_Definitions.txt:1267
type TraceStart struct {
	GNBCUCPUEE1APID int64           `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID int64           `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	TraceActivation TraceActivation `asn1:"mandatory,ignore,ext"`
}
