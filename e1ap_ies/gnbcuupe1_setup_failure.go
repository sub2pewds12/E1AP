package e1ap_ies

// GNBCUUPE1SetupFailure represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:402
type GNBCUUPE1SetupFailure struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
