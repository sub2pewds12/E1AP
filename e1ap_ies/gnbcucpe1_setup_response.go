package e1ap_ies

// GNBCUCPE1SetupResponse represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:445
type GNBCUCPE1SetupResponse struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
