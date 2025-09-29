package e1ap_ies

// ResourceStatusResponse represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1362
type ResourceStatusResponse struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
