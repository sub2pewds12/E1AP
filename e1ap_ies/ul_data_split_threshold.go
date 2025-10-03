package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ULDataSplitThreshold From: 9_4_5_Information_Element_Definitions.txt:2393
// ASN.1 Data Type: ENUMERATED
const (
	ULDataSplitThresholdB0       aper.Enumerated = 0
	ULDataSplitThresholdB100     aper.Enumerated = 1
	ULDataSplitThresholdB200     aper.Enumerated = 2
	ULDataSplitThresholdB400     aper.Enumerated = 3
	ULDataSplitThresholdB800     aper.Enumerated = 4
	ULDataSplitThresholdB1600    aper.Enumerated = 5
	ULDataSplitThresholdB3200    aper.Enumerated = 6
	ULDataSplitThresholdB6400    aper.Enumerated = 7
	ULDataSplitThresholdB12800   aper.Enumerated = 8
	ULDataSplitThresholdB25600   aper.Enumerated = 9
	ULDataSplitThresholdB51200   aper.Enumerated = 10
	ULDataSplitThresholdB102400  aper.Enumerated = 11
	ULDataSplitThresholdB204800  aper.Enumerated = 12
	ULDataSplitThresholdB409600  aper.Enumerated = 13
	ULDataSplitThresholdB819200  aper.Enumerated = 14
	ULDataSplitThresholdB1228800 aper.Enumerated = 15
	ULDataSplitThresholdB1638400 aper.Enumerated = 16
	ULDataSplitThresholdB2457600 aper.Enumerated = 17
	ULDataSplitThresholdB3276800 aper.Enumerated = 18
	ULDataSplitThresholdB4096000 aper.Enumerated = 19
	ULDataSplitThresholdB4915200 aper.Enumerated = 20
	ULDataSplitThresholdB5734400 aper.Enumerated = 21
	ULDataSplitThresholdB6553600 aper.Enumerated = 22
	ULDataSplitThresholdInfinity aper.Enumerated = 23
)

type ULDataSplitThreshold struct {
	Value aper.Enumerated
}
