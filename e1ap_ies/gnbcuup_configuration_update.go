package e1ap_ies

// GNBCUUPConfigurationUpdate represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:492
type GNBCUUPConfigurationUpdate struct {
	ProtocolIEs ProtocolIEContainer `asn1:"mandatory,ext"`
}
