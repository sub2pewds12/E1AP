package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// UEAssociatedLogicalE1ConnectionListResAck is a generated LIST type.
type UEAssociatedLogicalE1ConnectionListResAck struct {
	Value []UEAssociatedLogicalE1ConnectionListResItemAckItem
}

func (s *UEAssociatedLogicalE1ConnectionListResAck) Encode(w *per.Encoder) (err error) {

	c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxnoofIndividualE1ConnectionsToReset)}
	if err = w.EncodeLengthDeterminant(int64(len(s.Value)), c); err != nil {
		return fmt.Errorf("encode length determinant failed: %w", err)
	}
	for i := 0; i < len(s.Value); i++ {
		if err = s.Value[i].Encode(w); err != nil {
			return fmt.Errorf("encode list item %d failed: %w", i, err)
		}
	}
	return nil
}

func (s *UEAssociatedLogicalE1ConnectionListResAck) Decode(r *per.Decoder) (err error) {

	c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxnoofIndividualE1ConnectionsToReset)}
	length, err := r.DecodeLengthDeterminant(c)
	if err != nil {
		return fmt.Errorf("decode length determinant failed: %w", err)
	}
	s.Value = make([]UEAssociatedLogicalE1ConnectionListResItemAckItem, length)
	for i := int64(0); i < length; i++ {
		item := new(UEAssociatedLogicalE1ConnectionListResItemAckItem)
		if err := item.Decode(r); err != nil {
			return fmt.Errorf("decode list item %d failed: %w", i, err)
		}
		s.Value[i] = *item
	}
	return nil
}
