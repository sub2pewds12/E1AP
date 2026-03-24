package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// Presence is a generated ENUMERATED type.
type Presence struct {
	Value int64
}

const (
	PresenceOptional    int64 = 0
	PresenceConditional int64 = 1
	PresenceMandatory   int64 = 2
)

// Encode implements the MessageEncoder interface for Presence.
func (e *Presence) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for Presence.
func (e *Presence) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
