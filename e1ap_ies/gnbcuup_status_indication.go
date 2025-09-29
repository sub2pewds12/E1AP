package e1ap_ies

// GNBCUUPStatusIndication represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1224
type GNBCUUPStatusIndication struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
