package e1ap_ies

// ActivityInformation is a generated CHOICE type.
type ActivityInformation struct {
	Choice                         uint64
	DRBActivityList                []DRBActivityItem
	PDUSessionResourceActivityList []PDUSessionResourceActivityItem
	UEActivity                     *UEActivity
	ChoiceExtension                *ProtocolIESingleContainer
}

const (
	ActivityInformationPresentNothing uint64 = iota
	ActivityInformationPresentDRBActivityList
	ActivityInformationPresentPDUSessionResourceActivityList
	ActivityInformationPresentUEActivity
	ActivityInformationPresentChoiceExtension
)
