package e1ap_ies

// DeactivateTrace represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1285
type DeactivateTrace struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
