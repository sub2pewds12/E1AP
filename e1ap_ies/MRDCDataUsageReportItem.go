package e1ap_ies

import (
	"fmt"
	"math"

	"github.com/lvdund/ngap/aper"
)

// MRDCDataUsageReportItem is a generated SEQUENCE type.
type MRDCDataUsageReportItem struct {
	StartTimeStamp MRDCDataUsageReportItemStartTimeStamp `aper:"lb:4,ub:4,mandatory,ext"`
	EndTimeStamp   MRDCDataUsageReportItemEndTimeStamp   `aper:"lb:4,ub:4,mandatory,ext"`
	UsageCountUL   MRDCDataUsageReportItemUsageCountUL   `aper:"lb:0,ub:18446744073709551615,mandatory,ext"`
	UsageCountDL   MRDCDataUsageReportItemUsageCountDL   `aper:"lb:0,ub:18446744073709551615,mandatory,ext"`
	IEExtensions   *ProtocolExtensionContainer           `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *MRDCDataUsageReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.StartTimeStamp.Value), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("encode StartTimeStamp failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.EndTimeStamp.Value), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("encode EndTimeStamp failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.UsageCountUL.Value), &aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false); err != nil {
		return fmt.Errorf("encode UsageCountUL failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.UsageCountDL.Value), &aper.Constraint{Lb: 0, Ub: math.MaxInt64}, false); err != nil {
		return fmt.Errorf("encode UsageCountDL failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCDataUsageReportItem) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.StartTimeStamp.Decode(r); err != nil {
		return fmt.Errorf("decode StartTimeStamp failed: %w", err)
	}
	if err = s.EndTimeStamp.Decode(r); err != nil {
		return fmt.Errorf("decode EndTimeStamp failed: %w", err)
	}
	if err = s.UsageCountUL.Decode(r); err != nil {
		return fmt.Errorf("decode UsageCountUL failed: %w", err)
	}
	if err = s.UsageCountDL.Decode(r); err != nil {
		return fmt.Errorf("decode UsageCountDL failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for MRDCDataUsageReportItem */
	}
	return nil
}
