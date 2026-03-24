package e1ap_ies

import (
	"asn1go/per"
)

// ULDataSplitThreshold is a generated ENUMERATED type.
type ULDataSplitThreshold struct {
	Value int64
}

const (
	ULDataSplitThresholdB0       int64 = 0
	ULDataSplitThresholdB100     int64 = 1
	ULDataSplitThresholdB200     int64 = 2
	ULDataSplitThresholdB400     int64 = 3
	ULDataSplitThresholdB800     int64 = 4
	ULDataSplitThresholdB1600    int64 = 5
	ULDataSplitThresholdB3200    int64 = 6
	ULDataSplitThresholdB6400    int64 = 7
	ULDataSplitThresholdB12800   int64 = 8
	ULDataSplitThresholdB25600   int64 = 9
	ULDataSplitThresholdB51200   int64 = 10
	ULDataSplitThresholdB102400  int64 = 11
	ULDataSplitThresholdB204800  int64 = 12
	ULDataSplitThresholdB409600  int64 = 13
	ULDataSplitThresholdB819200  int64 = 14
	ULDataSplitThresholdB1228800 int64 = 15
	ULDataSplitThresholdB1638400 int64 = 16
	ULDataSplitThresholdB2457600 int64 = 17
	ULDataSplitThresholdB3276800 int64 = 18
	ULDataSplitThresholdB4096000 int64 = 19
	ULDataSplitThresholdB4915200 int64 = 20
	ULDataSplitThresholdB5734400 int64 = 21
	ULDataSplitThresholdB6553600 int64 = 22
	ULDataSplitThresholdInfinity int64 = 23
)

// Encode implements the MessageEncoder interface for ULDataSplitThreshold.
func (e *ULDataSplitThreshold) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 24), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for ULDataSplitThreshold.
func (e *ULDataSplitThreshold) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 24), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
