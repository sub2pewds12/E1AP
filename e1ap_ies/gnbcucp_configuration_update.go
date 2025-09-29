package e1ap_ies

// GNBCUCPConfigurationUpdate represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:559
type GNBCUCPConfigurationUpdate struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
