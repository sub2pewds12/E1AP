package e1ap_ies

// GNBCUCPConfigurationUpdateAcknowledge represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:584
type GNBCUCPConfigurationUpdateAcknowledge struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
