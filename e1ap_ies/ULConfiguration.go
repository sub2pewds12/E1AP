package e1ap_ies

import (
	"asn1go/per"
)

// ULConfiguration is a generated ENUMERATED type.
type ULConfiguration struct {
	Value int64
}

const (
	ULConfigurationNoData int64 = 0
	ULConfigurationShared int64 = 1
	ULConfigurationOnly   int64 = 2
)

// Encode implements the MessageEncoder interface for ULConfiguration.
func (e *ULConfiguration) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for ULConfiguration.
func (e *ULConfiguration) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
