package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DataUsagePerQOSFlowItem is a generated LIST type.
type DataUsagePerQOSFlowItem struct {
	Value []MRDCDataUsageReportItem
}

func (s *DataUsagePerQOSFlowItem) Encode(w *per.Encoder) (err error) {

	c := per.SizeConstraints{Extensible: true, Min: int64Ptr(1), Max: int64Ptr(Maxnooftimeperiods)}
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

func (s *DataUsagePerQOSFlowItem) Decode(r *per.Decoder) (err error) {

	c := per.SizeConstraints{Extensible: true, Min: int64Ptr(1), Max: int64Ptr(Maxnooftimeperiods)}
	length, err := r.DecodeLengthDeterminant(c)
	if err != nil {
		return fmt.Errorf("decode length determinant failed: %w", err)
	}
	s.Value = make([]MRDCDataUsageReportItem, length)
	for i := int64(0); i < length; i++ {
		item := new(MRDCDataUsageReportItem)
		if err := item.Decode(r); err != nil {
			return fmt.Errorf("decode list item %d failed: %w", i, err)
		}
		s.Value[i] = *item
	}
	return nil
}
