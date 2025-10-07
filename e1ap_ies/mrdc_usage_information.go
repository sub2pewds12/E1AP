package e1ap_ies

// MRDCUsageInformation is a generated SEQUENCE type.
type MRDCUsageInformation struct {
	DataUsagePerPDUSessionReport *DataUsagePerPDUSessionReport `aper:"optional,ext"`
	DataUsagePerQOSFlowList      []DataUsagePerQOSFlowItem     `aper:"lb:1,ub:MaxnoofQoSFlows,optional,ext"`
	IEExtensions                 *ProtocolExtensionContainer   `aper:"optional,ext"`
}
