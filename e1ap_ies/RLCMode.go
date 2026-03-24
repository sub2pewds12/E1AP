package e1ap_ies

import (
	"asn1go/per"
)

// RLCMode is a generated ENUMERATED type.
type RLCMode struct {
	Value int64
}

const (
	RLCModeRlcTm                 int64 = 0
	RLCModeRlcAm                 int64 = 1
	RLCModeRlcUmBidirectional    int64 = 2
	RLCModeRlcUmUnidirectionalUl int64 = 3
	RLCModeRlcUmUnidirectionalDl int64 = 4
)

// Encode implements the MessageEncoder interface for RLCMode.
func (e *RLCMode) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for RLCMode.
func (e *RLCMode) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
