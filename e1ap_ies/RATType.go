package e1ap_ies

import (
	"asn1go/per"
)

// RATType is a generated ENUMERATED type.
type RATType struct {
	Value int64
}

const (
	RATTypeEUTRA int64 = 0
	RATTypeNR    int64 = 1
)

// Encode implements the MessageEncoder interface for RATType.
func (e *RATType) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for RATType.
func (e *RATType) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
