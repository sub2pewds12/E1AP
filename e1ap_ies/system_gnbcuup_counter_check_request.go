package e1ap_ies

// SystemGNBCUUPCounterCheckRequest represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:1190
// SystemGNBCUUPCounterCheckRequest represents a CHOICE type.
type SystemGNBCUUPCounterCheckRequest struct {
	EUTRANGNBCUUPCounterCheckRequest ProtocolIEContainer `asn1:"mandatory"`
	NGRANGNBCUUPCounterCheckRequest  ProtocolIEContainer `asn1:"mandatory"`
}
