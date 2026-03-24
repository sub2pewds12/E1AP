package e1ap_ies

import (
	"asn1go/per"
)

// Dynamic5QIDescriptorDelayCritical is a generated ENUMERATED type.
type Dynamic5QIDescriptorDelayCritical struct {
	Value int64
}

const (
	Dynamic5QIDescriptorDelayCriticalDelayCritical    int64 = 0
	Dynamic5QIDescriptorDelayCriticalNonDelayCritical int64 = 1
)

// Encode implements the MessageEncoder interface for Dynamic5QIDescriptorDelayCritical.
func (e *Dynamic5QIDescriptorDelayCritical) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for Dynamic5QIDescriptorDelayCritical.
func (e *Dynamic5QIDescriptorDelayCritical) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
