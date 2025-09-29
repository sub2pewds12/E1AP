package e1ap_ies

// DataUsagePerQOSFlowItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:359
type DataUsagePerQOSFlowItem struct {
	QOSFlowIdentifier int64                                   `asn1:"lb:0,ub:63,mandatory,ext"`
	SecondaryRATType  DataUsagePerQOSFlowItemSecondaryRATType `asn1:"mandatory,ext"`
}
