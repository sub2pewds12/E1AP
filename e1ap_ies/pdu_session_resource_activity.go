package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceActivity is a generated ENUMERATED type.
type PDUSessionResourceActivity struct {
	Value aper.Enumerated
}

const (
	PDUSessionResourceActivityActive    aper.Enumerated = 0
	PDUSessionResourceActivityNotActive aper.Enumerated = 1
)
