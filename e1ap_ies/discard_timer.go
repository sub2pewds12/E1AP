package e1ap_ies

// DiscardTimer represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:391
type DiscardTimer int32

const (
	DiscardTimer_Ms10     DiscardTimer = 0
	DiscardTimer_Ms20     DiscardTimer = 1
	DiscardTimer_Ms30     DiscardTimer = 2
	DiscardTimer_Ms40     DiscardTimer = 3
	DiscardTimer_Ms50     DiscardTimer = 4
	DiscardTimer_Ms60     DiscardTimer = 5
	DiscardTimer_Ms75     DiscardTimer = 6
	DiscardTimer_Ms100    DiscardTimer = 7
	DiscardTimer_Ms150    DiscardTimer = 8
	DiscardTimer_Ms200    DiscardTimer = 9
	DiscardTimer_Ms250    DiscardTimer = 10
	DiscardTimer_Ms300    DiscardTimer = 11
	DiscardTimer_Ms500    DiscardTimer = 12
	DiscardTimer_Ms750    DiscardTimer = 13
	DiscardTimer_Ms1500   DiscardTimer = 14
	DiscardTimer_Infinity DiscardTimer = 15
)
