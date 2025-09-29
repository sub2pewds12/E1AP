package e1ap_ies

// DRBRequiredToModifyItemNGRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:607
type DRBRequiredToModifyItemNGRAN struct {
	DRBID                                int64                                      `asn1:"lb:1,ub:32,mandatory,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `asn1:"lb:1,ub:Item,optional,ext"`
	FlowToRemove                         []QOSFlowItem                              `asn1:"lb:1,ub:Item,optional,ext"`
	Cause                                *Cause                                     `asn1:"optional,ext"`
}
