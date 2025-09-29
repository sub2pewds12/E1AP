package e1ap_ies

// TraceStart represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1267
type TraceStart struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
