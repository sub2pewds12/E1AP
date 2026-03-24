package e1ap_ies

import (
	"asn1go/per"
)

// TimeToWait is a generated ENUMERATED type.
type TimeToWait struct {
	Value int64
}

const (
	TimeToWaitV1s  int64 = 0
	TimeToWaitV2s  int64 = 1
	TimeToWaitV5s  int64 = 2
	TimeToWaitV10s int64 = 3
	TimeToWaitV20s int64 = 4
	TimeToWaitV60s int64 = 5
)

// Encode implements the MessageEncoder interface for TimeToWait.
func (e *TimeToWait) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 6), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for TimeToWait.
func (e *TimeToWait) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 6), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
