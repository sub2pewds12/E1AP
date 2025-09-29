package e1ap_ies

// PDUSessionResourceActivity represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1647
type PDUSessionResourceActivity int32

const (
	PDUSessionResourceActivity_Active    PDUSessionResourceActivity = 0
	PDUSessionResourceActivity_NotActive PDUSessionResourceActivity = 1
)
