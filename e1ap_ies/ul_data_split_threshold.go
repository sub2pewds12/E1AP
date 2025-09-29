package e1ap_ies

// ULDataSplitThreshold represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2393
type ULDataSplitThreshold int32

const (
	ULDataSplitThreshold_B0       ULDataSplitThreshold = 0
	ULDataSplitThreshold_B100     ULDataSplitThreshold = 1
	ULDataSplitThreshold_B200     ULDataSplitThreshold = 2
	ULDataSplitThreshold_B400     ULDataSplitThreshold = 3
	ULDataSplitThreshold_B800     ULDataSplitThreshold = 4
	ULDataSplitThreshold_B1600    ULDataSplitThreshold = 5
	ULDataSplitThreshold_B3200    ULDataSplitThreshold = 6
	ULDataSplitThreshold_B6400    ULDataSplitThreshold = 7
	ULDataSplitThreshold_B12800   ULDataSplitThreshold = 8
	ULDataSplitThreshold_B25600   ULDataSplitThreshold = 9
	ULDataSplitThreshold_B51200   ULDataSplitThreshold = 10
	ULDataSplitThreshold_B102400  ULDataSplitThreshold = 11
	ULDataSplitThreshold_B204800  ULDataSplitThreshold = 12
	ULDataSplitThreshold_B409600  ULDataSplitThreshold = 13
	ULDataSplitThreshold_B819200  ULDataSplitThreshold = 14
	ULDataSplitThreshold_B1228800 ULDataSplitThreshold = 15
	ULDataSplitThreshold_B1638400 ULDataSplitThreshold = 16
	ULDataSplitThreshold_B2457600 ULDataSplitThreshold = 17
	ULDataSplitThreshold_B3276800 ULDataSplitThreshold = 18
	ULDataSplitThreshold_B4096000 ULDataSplitThreshold = 19
	ULDataSplitThreshold_B4915200 ULDataSplitThreshold = 20
	ULDataSplitThreshold_B5734400 ULDataSplitThreshold = 21
	ULDataSplitThreshold_B6553600 ULDataSplitThreshold = 22
	ULDataSplitThreshold_Infinity ULDataSplitThreshold = 23
)
