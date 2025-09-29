package e1ap_ies

// GNBCUUPConfigurationUpdateFailure represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:534
type GNBCUUPConfigurationUpdateFailure struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
