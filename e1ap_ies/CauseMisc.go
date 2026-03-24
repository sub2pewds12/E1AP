package e1ap_ies

import (
	"asn1go/per"
)

// CauseMisc is a generated ENUMERATED type.
type CauseMisc struct {
	Value int64
}

const (
	CauseMiscControlProcessingOverload             int64 = 0
	CauseMiscNotEnoughUserPlaneProcessingResources int64 = 1
	CauseMiscHardwareFailure                       int64 = 2
	CauseMiscOmIntervention                        int64 = 3
	CauseMiscUnspecified                           int64 = 4
)

// Encode implements the MessageEncoder interface for CauseMisc.
func (e *CauseMisc) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for CauseMisc.
func (e *CauseMisc) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
