package e1ap_ies

// CellTrafficTrace represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1303
type CellTrafficTrace struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
