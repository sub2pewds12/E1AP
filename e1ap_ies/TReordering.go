package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// TReordering is a generated ENUMERATED type.
type TReordering struct {
	Value int64
}

const (
	TReorderingMs0    int64 = 0
	TReorderingMs1    int64 = 1
	TReorderingMs2    int64 = 2
	TReorderingMs4    int64 = 3
	TReorderingMs5    int64 = 4
	TReorderingMs8    int64 = 5
	TReorderingMs10   int64 = 6
	TReorderingMs15   int64 = 7
	TReorderingMs20   int64 = 8
	TReorderingMs30   int64 = 9
	TReorderingMs40   int64 = 10
	TReorderingMs50   int64 = 11
	TReorderingMs60   int64 = 12
	TReorderingMs80   int64 = 13
	TReorderingMs100  int64 = 14
	TReorderingMs120  int64 = 15
	TReorderingMs140  int64 = 16
	TReorderingMs160  int64 = 17
	TReorderingMs180  int64 = 18
	TReorderingMs200  int64 = 19
	TReorderingMs220  int64 = 20
	TReorderingMs240  int64 = 21
	TReorderingMs260  int64 = 22
	TReorderingMs280  int64 = 23
	TReorderingMs300  int64 = 24
	TReorderingMs500  int64 = 25
	TReorderingMs750  int64 = 26
	TReorderingMs1000 int64 = 27
	TReorderingMs1250 int64 = 28
	TReorderingMs1500 int64 = 29
	TReorderingMs1750 int64 = 30
	TReorderingMs2000 int64 = 31
	TReorderingMs2250 int64 = 32
	TReorderingMs2500 int64 = 33
	TReorderingMs2750 int64 = 34
	TReorderingMs3000 int64 = 35
)

// Encode implements the MessageEncoder interface for TReordering.
func (e *TReordering) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 36), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for TReordering.
func (e *TReordering) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 36), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
