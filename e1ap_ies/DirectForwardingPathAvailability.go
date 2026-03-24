package e1ap_ies

import (
	"asn1go/per"
)

// DirectForwardingPathAvailability is a generated ENUMERATED type.
type DirectForwardingPathAvailability struct {
	Value int64
}

const (
	DirectForwardingPathAvailabilityInterSystemDirectPathAvailable int64 = 0
	DirectForwardingPathAvailabilityIntraSystemDirectPathAvailable int64 = 1
)

// Encode implements the MessageEncoder interface for DirectForwardingPathAvailability.
func (e *DirectForwardingPathAvailability) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for DirectForwardingPathAvailability.
func (e *DirectForwardingPathAvailability) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
