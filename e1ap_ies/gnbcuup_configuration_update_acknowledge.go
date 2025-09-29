package e1ap_ies

// GNBCUUPConfigurationUpdateAcknowledge represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:516
type GNBCUUPConfigurationUpdateAcknowledge struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
