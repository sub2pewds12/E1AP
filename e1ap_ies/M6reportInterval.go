package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// M6reportInterval is a generated ENUMERATED type.
type M6reportInterval struct {
	Value int64
}

const (
	M6reportIntervalMs120   int64 = 0
	M6reportIntervalMs240   int64 = 1
	M6reportIntervalMs480   int64 = 2
	M6reportIntervalMs640   int64 = 3
	M6reportIntervalMs1024  int64 = 4
	M6reportIntervalMs2048  int64 = 5
	M6reportIntervalMs5120  int64 = 6
	M6reportIntervalMs10240 int64 = 7
	M6reportIntervalMs20480 int64 = 8
	M6reportIntervalMs40960 int64 = 9
	M6reportIntervalMin1    int64 = 10
	M6reportIntervalMin6    int64 = 11
	M6reportIntervalMin12   int64 = 12
	M6reportIntervalMin30   int64 = 13
)

// Encode implements the MessageEncoder interface for M6reportInterval.
func (e *M6reportInterval) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 14), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for M6reportInterval.
func (e *M6reportInterval) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 14), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
