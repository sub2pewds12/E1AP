package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// DLTXStop is a generated ENUMERATED type.
type DLTXStop struct {
	Value int64
}

const (
	DLTXStopStop   int64 = 0
	DLTXStopResume int64 = 1
)

// Encode implements the MessageEncoder interface for DLTXStop.
func (e *DLTXStop) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for DLTXStop.
func (e *DLTXStop) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
