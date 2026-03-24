package e1ap_ies

import (
	"asn1go/per"
)

// PDCPDataRecovery is a generated ENUMERATED type.
type PDCPDataRecovery struct {
	Value int64
}

const (
	PDCPDataRecoveryTrue int64 = 0
)

// Encode implements the MessageEncoder interface for PDCPDataRecovery.
func (e *PDCPDataRecovery) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for PDCPDataRecovery.
func (e *PDCPDataRecovery) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
