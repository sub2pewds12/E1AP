package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// M6reportInterval From: 9_4_5_Information_Element_Definitions.txt:1363
// ASN.1 Data Type: ENUMERATED
const (
	M6reportIntervalMs120   aper.Enumerated = 0
	M6reportIntervalMs240   aper.Enumerated = 1
	M6reportIntervalMs480   aper.Enumerated = 2
	M6reportIntervalMs640   aper.Enumerated = 3
	M6reportIntervalMs1024  aper.Enumerated = 4
	M6reportIntervalMs2048  aper.Enumerated = 5
	M6reportIntervalMs5120  aper.Enumerated = 6
	M6reportIntervalMs10240 aper.Enumerated = 7
	M6reportIntervalMs20480 aper.Enumerated = 8
	M6reportIntervalMs40960 aper.Enumerated = 9
	M6reportIntervalMin1    aper.Enumerated = 10
	M6reportIntervalMin6    aper.Enumerated = 11
	M6reportIntervalMin12   aper.Enumerated = 12
	M6reportIntervalMin30   aper.Enumerated = 13
)

type M6reportInterval struct {
	Value aper.Enumerated
}
