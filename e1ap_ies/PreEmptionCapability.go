package e1ap_ies

import (
	"asn1go/per"
)

// PreEmptionCapability is a generated ENUMERATED type.
type PreEmptionCapability struct {
	Value int64
}

const (
	PreEmptionCapabilityShallNotTriggerPreEmption int64 = 0
	PreEmptionCapabilityMayTriggerPreEmption      int64 = 1
)

// Encode implements the MessageEncoder interface for PreEmptionCapability.
func (e *PreEmptionCapability) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for PreEmptionCapability.
func (e *PreEmptionCapability) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
