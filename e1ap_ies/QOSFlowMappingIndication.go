package e1ap_ies

import (
	"asn1go/per"
)

// QOSFlowMappingIndication is a generated ENUMERATED type.
type QOSFlowMappingIndication struct {
	Value int64
}

const (
	QOSFlowMappingIndicationUl int64 = 0
	QOSFlowMappingIndicationDl int64 = 1
)

// Encode implements the MessageEncoder interface for QOSFlowMappingIndication.
func (e *QOSFlowMappingIndication) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for QOSFlowMappingIndication.
func (e *QOSFlowMappingIndication) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
