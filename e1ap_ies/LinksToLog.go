package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// LinksToLog is a generated ENUMERATED type.
type LinksToLog struct {
	Value int64
}

const (
	LinksToLogUplink                int64 = 0
	LinksToLogDownlink              int64 = 1
	LinksToLogBothUplinkAndDownlink int64 = 2
)

// Encode implements the MessageEncoder interface for LinksToLog.
func (e *LinksToLog) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for LinksToLog.
func (e *LinksToLog) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
