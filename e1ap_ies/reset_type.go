package e1ap_ies

// ResetType represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:261
// ResetType represents a CHOICE type.
type ResetType struct {
	E1Interface       ResetAll                    `asn1:"mandatory"`
	PartOfE1Interface []ProtocolIESingleContainer `asn1:"lb:1,ub:SingleContainer,mandatory"`
}
