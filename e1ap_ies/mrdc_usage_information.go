package e1ap_ies

// MRDCUsageInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1328
type MRDCUsageInformation struct {
	DataUsagePerPDUSessionReport *DataUsagePerPDUSessionReport `asn1:"optional,ext"`
	DataUsagePerQOSFlowList      []DataUsagePerQOSFlowItem     `asn1:"lb:1,ub:Item,optional,ext"`
}
