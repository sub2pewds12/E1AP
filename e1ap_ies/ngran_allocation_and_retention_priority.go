package e1ap_ies

// NGRANAllocationAndRetentionPriority is a generated SEQUENCE type.
type NGRANAllocationAndRetentionPriority struct {
	PriorityLevel           PriorityLevel               `aper:"lb:0,ub:0,mandatory"`
	PreEmptionCapability    PreEmptionCapability        `aper:"mandatory"`
	PreEmptionVulnerability PreEmptionVulnerability     `aper:"mandatory"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional"`
}
