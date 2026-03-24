package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// PDCPStatusReportIndication is a generated ENUMERATED type.
type PDCPStatusReportIndication struct {
	Value int64
}

const (
	PDCPStatusReportIndicationDownlink int64 = 0
	PDCPStatusReportIndicationUplink   int64 = 1
	PDCPStatusReportIndicationBoth     int64 = 2
)

// Encode implements the MessageEncoder interface for PDCPStatusReportIndication.
func (e *PDCPStatusReportIndication) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for PDCPStatusReportIndication.
func (e *PDCPStatusReportIndication) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
