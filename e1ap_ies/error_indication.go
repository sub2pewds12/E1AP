package e1ap_ies

// ErrorIndication represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:319
type ErrorIndication struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
