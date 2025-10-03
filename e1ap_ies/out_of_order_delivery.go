package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// OutOfOrderDelivery From: 9_4_5_Information_Element_Definitions.txt:1515
// ASN.1 Data Type: ENUMERATED
const (
	OutOfOrderDeliveryTrue aper.Enumerated = 0
)

type OutOfOrderDelivery struct {
	Value aper.Enumerated
}
