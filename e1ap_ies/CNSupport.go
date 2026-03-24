package e1ap_ies

import (
	"asn1go/per"
)

// CNSupport is a generated ENUMERATED type.
type CNSupport struct {
	Value int64
}

const (
	CNSupportCEpc int64 = 0
	CNSupportC5gc int64 = 1
	CNSupportBoth int64 = 2
)

// Encode implements the MessageEncoder interface for CNSupport.
func (e *CNSupport) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for CNSupport.
func (e *CNSupport) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
