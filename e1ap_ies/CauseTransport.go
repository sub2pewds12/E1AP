package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// CauseTransport is a generated ENUMERATED type.
type CauseTransport struct {
	Value int64
}

const (
	CauseTransportUnspecified                  int64 = 0
	CauseTransportTransportResourceUnavailable int64 = 1
	CauseTransportUnknownTNLAddressForIAB      int64 = 2
)

// Encode implements the MessageEncoder interface for CauseTransport.
func (e *CauseTransport) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for CauseTransport.
func (e *CauseTransport) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
