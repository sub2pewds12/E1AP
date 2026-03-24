package e1ap_ies

import (
	"asn1go/per"
)

// TraceDepth is a generated ENUMERATED type.
type TraceDepth struct {
	Value int64
}

const (
	TraceDepthMinimum                               int64 = 0
	TraceDepthMedium                                int64 = 1
	TraceDepthMaximum                               int64 = 2
	TraceDepthMinimumWithoutVendorSpecificExtension int64 = 3
	TraceDepthMediumWithoutVendorSpecificExtension  int64 = 4
	TraceDepthMaximumWithoutVendorSpecificExtension int64 = 5
)

// Encode implements the MessageEncoder interface for TraceDepth.
func (e *TraceDepth) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 6), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for TraceDepth.
func (e *TraceDepth) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 6), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
