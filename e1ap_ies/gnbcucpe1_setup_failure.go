package e1ap_ies

// GNBCUCPE1SetupFailure represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:467
type GNBCUCPE1SetupFailure struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
