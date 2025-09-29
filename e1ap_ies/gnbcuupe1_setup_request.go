package e1ap_ies

// GNBCUUPE1SetupRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:345
type GNBCUUPE1SetupRequest struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
