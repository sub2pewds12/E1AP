package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// EUTRANAllocationAndRetentionPriority is a generated SEQUENCE type.
type EUTRANAllocationAndRetentionPriority struct {
	PriorityLevel           PriorityLevel
	PreEmptionCapability    PreEmptionCapability
	PreEmptionVulnerability PreEmptionVulnerability
	IEExtensions            *EUTRANAllocationAndRetentionPriorityExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *EUTRANAllocationAndRetentionPriority) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "priorityLevel", Optional: false},
			per.ComponentInfo{Name: "pre-emptionCapability", Optional: false},
			per.ComponentInfo{Name: "pre-emptionVulnerability", Optional: false},
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

	if err = w.EncodeInteger(int64(s.PriorityLevel.Value), per.Constrained(0, 15)); err != nil {
		return fmt.Errorf("encode PriorityLevel failed: %w", err)
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.PreEmptionCapability.Value), enumC); err != nil {
			return fmt.Errorf("encode PreEmptionCapability failed: %w", err)
		}
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.PreEmptionVulnerability.Value), enumC); err != nil {
			return fmt.Errorf("encode PreEmptionVulnerability failed: %w", err)
		}
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
func (s *EUTRANAllocationAndRetentionPriority) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "priorityLevel", Optional: false},
			per.ComponentInfo{Name: "pre-emptionCapability", Optional: false},
			per.ComponentInfo{Name: "pre-emptionVulnerability", Optional: false},
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
		val, err := r.DecodeInteger(per.Constrained(0, 15))
		if err != nil {
			return fmt.Errorf("decode PriorityLevel failed: %w", err)
		}
		s.PriorityLevel.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode PreEmptionCapability failed: %w", err)
		}
		s.PreEmptionCapability.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode PreEmptionVulnerability failed: %w", err)
		}
		s.PreEmptionVulnerability.Value = val
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(EUTRANAllocationAndRetentionPriorityExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
