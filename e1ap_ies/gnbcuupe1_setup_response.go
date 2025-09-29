package e1ap_ies

// GNBCUUPE1SetupResponse represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:384
type GNBCUUPE1SetupResponse struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
