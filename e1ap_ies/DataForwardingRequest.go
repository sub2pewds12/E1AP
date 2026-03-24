package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// DataForwardingRequest is a generated ENUMERATED type.
type DataForwardingRequest struct {
	Value int64
}

const (
	DataForwardingRequestUL   int64 = 0
	DataForwardingRequestDL   int64 = 1
	DataForwardingRequestBoth int64 = 2
)

// Encode implements the MessageEncoder interface for DataForwardingRequest.
func (e *DataForwardingRequest) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for DataForwardingRequest.
func (e *DataForwardingRequest) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
