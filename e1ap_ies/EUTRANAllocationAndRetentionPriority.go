package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EUTRANAllocationAndRetentionPriority is a generated SEQUENCE type.
type EUTRANAllocationAndRetentionPriority struct {
	PriorityLevel           PriorityLevel               `aper:"lb:0,ub:15,mandatory,ext"`
	PreEmptionCapability    PreEmptionCapability        `aper:"mandatory,ext"`
	PreEmptionVulnerability PreEmptionVulnerability     `aper:"mandatory,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *EUTRANAllocationAndRetentionPriority) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteInteger(int64(s.PriorityLevel.Value), &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return fmt.Errorf("encode PriorityLevel failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.PreEmptionCapability.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return fmt.Errorf("encode PreEmptionCapability failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.PreEmptionVulnerability.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return fmt.Errorf("encode PreEmptionVulnerability failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EUTRANAllocationAndRetentionPriority) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.PriorityLevel.Decode(r); err != nil {
		return fmt.Errorf("decode PriorityLevel failed: %w", err)
	}
	if err = s.PreEmptionCapability.Decode(r); err != nil {
		return fmt.Errorf("decode PreEmptionCapability failed: %w", err)
	}
	if err = s.PreEmptionVulnerability.Decode(r); err != nil {
		return fmt.Errorf("decode PreEmptionVulnerability failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for EUTRANAllocationAndRetentionPriority */
	}
	return nil
}
