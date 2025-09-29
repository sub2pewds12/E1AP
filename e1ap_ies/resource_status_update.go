package e1ap_ies

// ResourceStatusUpdate represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1401
type ResourceStatusUpdate struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
