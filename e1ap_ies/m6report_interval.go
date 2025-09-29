package e1ap_ies

// M6reportInterval represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1363
type M6reportInterval int32

const (
	M6reportInterval_Ms120   M6reportInterval = 0
	M6reportInterval_Ms240   M6reportInterval = 1
	M6reportInterval_Ms480   M6reportInterval = 2
	M6reportInterval_Ms640   M6reportInterval = 3
	M6reportInterval_Ms1024  M6reportInterval = 4
	M6reportInterval_Ms2048  M6reportInterval = 5
	M6reportInterval_Ms5120  M6reportInterval = 6
	M6reportInterval_Ms10240 M6reportInterval = 7
	M6reportInterval_Ms20480 M6reportInterval = 8
	M6reportInterval_Ms40960 M6reportInterval = 9
	M6reportInterval_Min1    M6reportInterval = 10
	M6reportInterval_Min6    M6reportInterval = 11
	M6reportInterval_Min12   M6reportInterval = 12
	M6reportInterval_Min30   M6reportInterval = 13
)
