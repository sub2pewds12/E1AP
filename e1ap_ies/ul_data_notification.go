package e1ap_ies

// ULDataNotification From: 9_4_4_PDU_Definitions.txt:1129
type ULDataNotification struct {
	GNBCUCPUEE1APID        int64                    `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID        int64                    `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	PDUSessionToNotifyList []PDUSessionToNotifyItem `asn1:"mandatory,reject,ext"`
}
