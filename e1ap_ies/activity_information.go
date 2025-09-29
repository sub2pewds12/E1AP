package e1ap_ies

// ActivityInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:85
// ActivityInformation represents a CHOICE type.
type ActivityInformation struct {
	DRBActivityList                []DRBActivityItem                `asn1:"lb:1,ub:Item,mandatory"`
	PDUSessionResourceActivityList []PDUSessionResourceActivityItem `asn1:"lb:1,ub:Item,mandatory"`
	UEActivity                     UEActivity                       `asn1:"mandatory"`
}
