package e1ap_ies

// CellTrafficTrace From: 9_4_4_PDU_Definitions.txt:1303
type CellTrafficTrace struct {
	GNBCUCPUEE1APID                int64             `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                int64             `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	TraceID                        []byte            `asn1:"lb:8,ub:8,mandatory,ignore,ext"`
	TraceCollectionEntityIPAddress []byte            `asn1:"mandatory,ignore,ext"`
	PrivacyIndicator               *PrivacyIndicator `asn1:"optional,ignore,ext"`
	URIaddress                     *[]byte           `asn1:"optional,ignore,ext"`
}
