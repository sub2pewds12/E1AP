package e1ap_ies

// RetainabilityMeasurementsInfo From: 9_4_5_Information_Element_Definitions.txt:2079
type RetainabilityMeasurementsInfo []DRBRemovedItem

// DRBRemovedItem From: 9_4_5_Information_Element_Definitions.txt:577
type DRBRemovedItem struct {
	DRBID                     int64                               `asn1:"mandatory,ext"`
	DRBReleasedInSession      *DRBRemovedItemDRBReleasedInSession `asn1:"optional,ext"`
	DRBAccumulatedSessionTime *[]byte                             `asn1:"optional,ext"`
	QOSFlowRemovedList        *SEQUENCE                           `asn1:"optional,ext"`
}
