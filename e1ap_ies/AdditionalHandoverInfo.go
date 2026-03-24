package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// AdditionalHandoverInfo is a generated ENUMERATED type.
type AdditionalHandoverInfo struct {
	Value int64
}

const (
	AdditionalHandoverInfoDiscardPdpcSN int64 = 0
)

// Encode implements the MessageEncoder interface for AdditionalHandoverInfo.
func (e *AdditionalHandoverInfo) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for AdditionalHandoverInfo.
func (e *AdditionalHandoverInfo) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
