package e1ap_ies

// ResourceStatusFailure represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1381
type ResourceStatusFailure struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
