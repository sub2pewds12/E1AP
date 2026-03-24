package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// DiscardTimer is a generated ENUMERATED type.
type DiscardTimer struct {
	Value int64
}

const (
	DiscardTimerMs10     int64 = 0
	DiscardTimerMs20     int64 = 1
	DiscardTimerMs30     int64 = 2
	DiscardTimerMs40     int64 = 3
	DiscardTimerMs50     int64 = 4
	DiscardTimerMs60     int64 = 5
	DiscardTimerMs75     int64 = 6
	DiscardTimerMs100    int64 = 7
	DiscardTimerMs150    int64 = 8
	DiscardTimerMs200    int64 = 9
	DiscardTimerMs250    int64 = 10
	DiscardTimerMs300    int64 = 11
	DiscardTimerMs500    int64 = 12
	DiscardTimerMs750    int64 = 13
	DiscardTimerMs1500   int64 = 14
	DiscardTimerInfinity int64 = 15
)

// Encode implements the MessageEncoder interface for DiscardTimer.
func (e *DiscardTimer) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 16), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for DiscardTimer.
func (e *DiscardTimer) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 16), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
