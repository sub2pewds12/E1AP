package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// EHCUplinkParametersDRBContinueEHCUL is a generated ENUMERATED type.
type EHCUplinkParametersDRBContinueEHCUL struct {
	Value int64
}

const (
	EHCUplinkParametersDRBContinueEHCULTrue int64 = 0
)

// Encode implements the MessageEncoder interface for EHCUplinkParametersDRBContinueEHCUL.
func (e *EHCUplinkParametersDRBContinueEHCUL) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for EHCUplinkParametersDRBContinueEHCUL.
func (e *EHCUplinkParametersDRBContinueEHCUL) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
