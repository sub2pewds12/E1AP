package e1ap_ies

import (
	"asn1go/per"
)

// PrivacyIndicator is a generated ENUMERATED type.
type PrivacyIndicator struct {
	Value int64
}

const (
	PrivacyIndicatorImmediateMDT int64 = 0
	PrivacyIndicatorLoggedMDT    int64 = 1
)

// Encode implements the MessageEncoder interface for PrivacyIndicator.
func (e *PrivacyIndicator) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for PrivacyIndicator.
func (e *PrivacyIndicator) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
