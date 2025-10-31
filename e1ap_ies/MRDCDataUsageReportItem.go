package e1ap_ies

import (
	"fmt"

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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.StartTimeStamp), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode StartTimeStamp failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.EndTimeStamp), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode EndTimeStamp failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.UsageCountUL), &aper.Constraint{Lb: 0, Ub: 18446744073709551615}, false); err != nil {
		return fmt.Errorf("Encode UsageCountUL failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.UsageCountDL), &aper.Constraint{Lb: 0, Ub: 18446744073709551615}, false); err != nil {
		return fmt.Errorf("Encode UsageCountDL failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MRDCDataUsageReportItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val []byte
		if val, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return fmt.Errorf("Decode StartTimeStamp failed: %w", err)
		}
		s.StartTimeStamp = MRDCDataUsageReportItemStartTimeStamp(val)
	}

	{
		var val []byte
		if val, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return fmt.Errorf("Decode EndTimeStamp failed: %w", err)
		}
		s.EndTimeStamp = MRDCDataUsageReportItemEndTimeStamp(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 18446744073709551615}, false); err != nil {
			return fmt.Errorf("Decode UsageCountUL failed: %w", err)
		}
		s.UsageCountUL = MRDCDataUsageReportItemUsageCountUL(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 18446744073709551615}, false); err != nil {
			return fmt.Errorf("Decode UsageCountDL failed: %w", err)
		}
		s.UsageCountDL = MRDCDataUsageReportItemUsageCountDL(val)
	}

	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for MRDCDataUsageReportItem")
	}
	return nil
}
