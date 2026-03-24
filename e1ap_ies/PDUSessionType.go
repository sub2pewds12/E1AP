package e1ap_ies

import (
	"asn1go/per"
)

// PDUSessionType is a generated ENUMERATED type.
type PDUSessionType struct {
	Value int64
}

const (
	PDUSessionTypeIpv4         int64 = 0
	PDUSessionTypeIpv6         int64 = 1
	PDUSessionTypeIpv4v6       int64 = 2
	PDUSessionTypeEthernet     int64 = 3
	PDUSessionTypeUnstructured int64 = 4
)

// Encode implements the MessageEncoder interface for PDUSessionType.
func (e *PDUSessionType) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for PDUSessionType.
func (e *PDUSessionType) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
