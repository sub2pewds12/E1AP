package e1ap_ies

import "github.com/lvdund/ngap/aper"

// NGRANAllocationAndRetentionPriority is a generated SEQUENCE type.
type NGRANAllocationAndRetentionPriority struct {
	PriorityLevel           aper.Integer                `aper:"lb:0,ub:0,mandatory"`
	PreEmptionCapability    PreEmptionCapability        `aper:"mandatory"`
	PreEmptionVulnerability PreEmptionVulnerability     `aper:"mandatory"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional"`
}
