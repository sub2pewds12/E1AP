package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// PDUSessionResourceActivity is a generated ENUMERATED type.
type PDUSessionResourceActivity struct {
	Value int64
}

const (
	PDUSessionResourceActivityActive    int64 = 0
	PDUSessionResourceActivityNotActive int64 = 1
)

// Encode implements the MessageEncoder interface for PDUSessionResourceActivity.
func (e *PDUSessionResourceActivity) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for PDUSessionResourceActivity.
func (e *PDUSessionResourceActivity) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
