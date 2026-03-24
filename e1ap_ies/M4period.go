package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// M4period is a generated ENUMERATED type.
type M4period struct {
	Value int64
}

const (
	M4periodMs1024  int64 = 0
	M4periodMs2048  int64 = 1
	M4periodMs5120  int64 = 2
	M4periodMs10240 int64 = 3
	M4periodMin1    int64 = 4
)

// Encode implements the MessageEncoder interface for M4period.
func (e *M4period) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for M4period.
func (e *M4period) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
