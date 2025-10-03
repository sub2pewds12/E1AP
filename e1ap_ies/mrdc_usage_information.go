package e1ap_ies

// MRDCUsageInformation From: 9_4_5_Information_Element_Definitions.txt:1328
// ASN.1 Data Type: SEQUENCE
type MRDCUsageInformation struct {
	DataUsagePerPDUSessionReport *DataUsagePerPDUSessionReport `aper:"optional,ext"`
	DataUsagePerQOSFlowList      []DataUsagePerQOSFlowItem     `aper:"lb:1,ub:MaxnoofQoSFlows,optional,ext"`
}
