package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EUTRANAllocationAndRetentionPriority is a generated SEQUENCE type.
type EUTRANAllocationAndRetentionPriority struct {
	PriorityLevel           PriorityLevel               `aper:"lb:0,ub:0,mandatory,ext"`
	PreEmptionCapability    PreEmptionCapability        `aper:"mandatory,ext"`
	PreEmptionVulnerability PreEmptionVulnerability     `aper:"mandatory,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *EUTRANAllocationAndRetentionPriority) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteInteger(int64(s.PriorityLevel), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return fmt.Errorf("Encode PriorityLevel failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.PreEmptionCapability.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode PreEmptionCapability failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.PreEmptionVulnerability.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode PreEmptionVulnerability failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EUTRANAllocationAndRetentionPriority) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode PriorityLevel failed: %w", err)
		}
		s.PriorityLevel = PriorityLevel(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return fmt.Errorf("Decode PreEmptionCapability failed: %w", err)
		}
		s.PreEmptionCapability.Value = aper.Enumerated(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return fmt.Errorf("Decode PreEmptionVulnerability failed: %w", err)
		}
		s.PreEmptionVulnerability.Value = aper.Enumerated(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for EUTRANAllocationAndRetentionPriority")
	}
	return nil
}
