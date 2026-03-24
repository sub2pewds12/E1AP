package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// MRDCDataUsageReportItem is a generated SEQUENCE type.
type MRDCDataUsageReportItem struct {
	StartTimeStamp MRDCDataUsageReportItemStartTimeStamp
	EndTimeStamp   MRDCDataUsageReportItemEndTimeStamp
	UsageCountUL   MRDCDataUsageReportItemUsageCountUL
	UsageCountDL   MRDCDataUsageReportItemUsageCountDL
	IEExtensions   *MRDCDataUsageReportItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCDataUsageReportItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "startTimeStamp", Optional: false},
			per.ComponentInfo{Name: "endTimeStamp", Optional: false},
			per.ComponentInfo{Name: "usageCountUL", Optional: false},
			per.ComponentInfo{Name: "usageCountDL", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeOctetString([]byte(s.StartTimeStamp.Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)}); err != nil {
		return fmt.Errorf("encode StartTimeStamp failed: %w", err)
	}
	if err = w.EncodeOctetString([]byte(s.EndTimeStamp.Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)}); err != nil {
		return fmt.Errorf("encode EndTimeStamp failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.UsageCountUL.Value), per.Unconstrained()); err != nil {
		return fmt.Errorf("encode UsageCountUL failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.UsageCountDL.Value), per.Unconstrained()); err != nil {
		return fmt.Errorf("encode UsageCountDL failed: %w", err)
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCDataUsageReportItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "startTimeStamp", Optional: false},
			per.ComponentInfo{Name: "endTimeStamp", Optional: false},
			per.ComponentInfo{Name: "usageCountUL", Optional: false},
			per.ComponentInfo{Name: "usageCountDL", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	{
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
		if err != nil {
			return fmt.Errorf("decode StartTimeStamp failed: %w", err)
		}
		s.StartTimeStamp.Value = val
	}

	{
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(4), Max: int64Ptr(4)})
		if err != nil {
			return fmt.Errorf("decode EndTimeStamp failed: %w", err)
		}
		s.EndTimeStamp.Value = val
	}

	{
		val, err := r.DecodeInteger(per.Unconstrained())
		if err != nil {
			return fmt.Errorf("decode UsageCountUL failed: %w", err)
		}
		s.UsageCountUL.Value = val
	}

	{
		val, err := r.DecodeInteger(per.Unconstrained())
		if err != nil {
			return fmt.Errorf("decode UsageCountDL failed: %w", err)
		}
		s.UsageCountDL.Value = val
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(MRDCDataUsageReportItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
