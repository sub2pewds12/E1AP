package e1ap_ies

// NGRANAllocationAndRetentionPriority represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1415
type NGRANAllocationAndRetentionPriority struct {
	PriorityLevel           int64                   `asn1:"lb:0,ub:15,mandatory"`
	PreEmptionCapability    PreEmptionCapability    `asn1:"mandatory"`
	PreEmptionVulnerability PreEmptionVulnerability `asn1:"mandatory"`
}
