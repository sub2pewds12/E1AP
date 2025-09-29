package e1ap_ies

// GNBCUCPE1SetupRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:427
type GNBCUCPE1SetupRequest struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
