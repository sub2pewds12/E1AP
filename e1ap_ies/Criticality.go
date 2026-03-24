package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// Criticality is a generated ENUMERATED type.
type Criticality struct {
	Value int64
}

const (
	CriticalityReject int64 = 0
	CriticalityIgnore int64 = 1
	CriticalityNotify int64 = 2
)

// Encode implements the MessageEncoder interface for Criticality.
func (e *Criticality) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for Criticality.
func (e *Criticality) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
