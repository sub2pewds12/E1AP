package e1ap_ies

import (
	"asn1go/per"
)

// PDCPSNStatusRequest is a generated ENUMERATED type.
type PDCPSNStatusRequest struct {
	Value int64
}

const (
	PDCPSNStatusRequestRequested int64 = 0
)

// Encode implements the MessageEncoder interface for PDCPSNStatusRequest.
func (e *PDCPSNStatusRequest) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for PDCPSNStatusRequest.
func (e *PDCPSNStatusRequest) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
