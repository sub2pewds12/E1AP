package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionType is a generated ENUMERATED type.
type PDUSessionType struct {
	Value aper.Enumerated
}

const (
	PDUSessionTypeIpv4         aper.Enumerated = 0
	PDUSessionTypeIpv6         aper.Enumerated = 1
	PDUSessionTypeIpv4v6       aper.Enumerated = 2
	PDUSessionTypeEthernet     aper.Enumerated = 3
	PDUSessionTypeUnstructured aper.Enumerated = 4
)

func (e *PDUSessionType) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PDUSessionType to be generated here.
	return fmt.Errorf("Encode not implemented for enum PDUSessionType")
}

func (e *PDUSessionType) Decode(r *aper.AperReader) error {
	// Decode logic for enum PDUSessionType to be generated here.
	return fmt.Errorf("Decode not implemented for enum PDUSessionType")
}
