package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// NGRANAllocationAndRetentionPriority is a generated SEQUENCE type.
type NGRANAllocationAndRetentionPriority struct {
	PriorityLevel           PriorityLevel               `aper:"lb:0,ub:0,mandatory"`
	PreEmptionCapability    PreEmptionCapability        `aper:"mandatory"`
	PreEmptionVulnerability PreEmptionVulnerability     `aper:"mandatory"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *NGRANAllocationAndRetentionPriority) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PriorityLevel.Value), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return fmt.Errorf("Encode PriorityLevel failed: %w", err)
	}
	if err = s.PreEmptionCapability.Encode(w); err != nil {
		return fmt.Errorf("Encode PreEmptionCapability failed: %w", err)
	}
	if err = s.PreEmptionVulnerability.Encode(w); err != nil {
		return fmt.Errorf("Encode PreEmptionVulnerability failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NGRANAllocationAndRetentionPriority) Decode(r *aper.AperReader) (err error) {
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
		s.PriorityLevel.Value = aper.Integer(val)
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
		return fmt.Errorf("Extensions not yet implemented for NGRANAllocationAndRetentionPriority")
	}
	return nil
}
