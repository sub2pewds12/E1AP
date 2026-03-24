package e1ap_ies

import (
	"asn1go/per"
)

// PDCPReestablishment is a generated ENUMERATED type.
type PDCPReestablishment struct {
	Value int64
}

const (
	PDCPReestablishmentTrue int64 = 0
)

// Encode implements the MessageEncoder interface for PDCPReestablishment.
func (e *PDCPReestablishment) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for PDCPReestablishment.
func (e *PDCPReestablishment) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
