package e1ap_ies

import (
	"asn1go/per"
)

// RegistrationRequest is a generated ENUMERATED type.
type RegistrationRequest struct {
	Value int64
}

const (
	RegistrationRequestStart int64 = 0
	RegistrationRequestStop  int64 = 1
)

// Encode implements the MessageEncoder interface for RegistrationRequest.
func (e *RegistrationRequest) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for RegistrationRequest.
func (e *RegistrationRequest) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
