package e1ap_ies

// M4period represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1350
type M4period int32

const (
	M4period_Ms1024  M4period = 0
	M4period_Ms2048  M4period = 1
	M4period_Ms5120  M4period = 2
	M4period_Ms10240 M4period = 3
	M4period_Min1    M4period = 4
)
