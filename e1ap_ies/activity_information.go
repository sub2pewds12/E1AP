package e1ap_ies

// ActivityInformation From: 9_4_5_Information_Element_Definitions.txt:85
// ASN.1 Data Type: CHOICE
const (
	ActivityInformationPresentNothing uint64 = iota
	ActivityInformationPresentDRBActivityList
	ActivityInformationPresentPDUSessionResourceActivityList
	ActivityInformationPresentUEActivity
)

type ActivityInformation struct {
	Choice                         uint64
	DRBActivityList                []DRBActivityItem
	PDUSessionResourceActivityList []PDUSessionResourceActivityItem
	UEActivity                     *UEActivity
}
