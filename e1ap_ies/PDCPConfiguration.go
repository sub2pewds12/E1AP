package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// PDCPConfiguration is a generated SEQUENCE type.
type PDCPConfiguration struct {
	PDCPSNSizeUL          PDCPSNSize
	PDCPSNSizeDL          PDCPSNSize
	RLCMode               RLCMode
	ROHCParameters        *ROHCParameters
	TReorderingTimer      *TReorderingTimer
	DiscardTimer          *DiscardTimer
	ULDataSplitThreshold  *ULDataSplitThreshold
	PDCPDuplication       *PDCPDuplication
	PDCPReestablishment   *PDCPReestablishment
	PDCPDataRecovery      *PDCPDataRecovery
	DuplicationActivation *DuplicationActivation
	OutOfOrderDelivery    *OutOfOrderDelivery
	IEExtensions          *PDCPConfigurationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDCPConfiguration) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDCP-SN-Size-UL", Optional: false},
			per.ComponentInfo{Name: "pDCP-SN-Size-DL", Optional: false},
			per.ComponentInfo{Name: "rLC-Mode", Optional: false},
			per.ComponentInfo{Name: "rOHC-Parameters", Optional: true},
			per.ComponentInfo{Name: "t-ReorderingTimer", Optional: true},
			per.ComponentInfo{Name: "discardTimer", Optional: true},
			per.ComponentInfo{Name: "uLDataSplitThreshold", Optional: true},
			per.ComponentInfo{Name: "pDCP-Duplication", Optional: true},
			per.ComponentInfo{Name: "pDCP-Reestablishment", Optional: true},
			per.ComponentInfo{Name: "pDCP-DataRecovery", Optional: true},
			per.ComponentInfo{Name: "duplication-Activation", Optional: true},
			per.ComponentInfo{Name: "outOfOrderDelivery", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ROHCParameters != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.TReorderingTimer != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DiscardTimer != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.ULDataSplitThreshold != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDCPDuplication != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDCPReestablishment != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDCPDataRecovery != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.DuplicationActivation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.OutOfOrderDelivery != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.PDCPSNSizeUL.Value), enumC); err != nil {
			return fmt.Errorf("encode PDCPSNSizeUL failed: %w", err)
		}
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.PDCPSNSizeDL.Value), enumC); err != nil {
			return fmt.Errorf("encode PDCPSNSizeDL failed: %w", err)
		}
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.RLCMode.Value), enumC); err != nil {
			return fmt.Errorf("encode RLCMode failed: %w", err)
		}
	}

	if s.ROHCParameters != nil {
		if err = s.ROHCParameters.Encode(w); err != nil {
			return fmt.Errorf("encode ROHCParameters failed: %w", err)
		}
	}

	if s.TReorderingTimer != nil {
		if err = s.TReorderingTimer.Encode(w); err != nil {
			return fmt.Errorf("encode TReorderingTimer failed: %w", err)
		}
	}

	if s.DiscardTimer != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 16), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.DiscardTimer).Value), enumC); err != nil {
				return fmt.Errorf("encode DiscardTimer failed: %w", err)
			}
		}
	}

	if s.ULDataSplitThreshold != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 24), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.ULDataSplitThreshold).Value), enumC); err != nil {
				return fmt.Errorf("encode ULDataSplitThreshold failed: %w", err)
			}
		}
	}

	if s.PDCPDuplication != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.PDCPDuplication).Value), enumC); err != nil {
				return fmt.Errorf("encode PDCPDuplication failed: %w", err)
			}
		}
	}

	if s.PDCPReestablishment != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.PDCPReestablishment).Value), enumC); err != nil {
				return fmt.Errorf("encode PDCPReestablishment failed: %w", err)
			}
		}
	}

	if s.PDCPDataRecovery != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.PDCPDataRecovery).Value), enumC); err != nil {
				return fmt.Errorf("encode PDCPDataRecovery failed: %w", err)
			}
		}
	}

	if s.DuplicationActivation != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.DuplicationActivation).Value), enumC); err != nil {
				return fmt.Errorf("encode DuplicationActivation failed: %w", err)
			}
		}
	}

	if s.OutOfOrderDelivery != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.OutOfOrderDelivery).Value), enumC); err != nil {
				return fmt.Errorf("encode OutOfOrderDelivery failed: %w", err)
			}
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
func (s *PDCPConfiguration) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDCP-SN-Size-UL", Optional: false},
			per.ComponentInfo{Name: "pDCP-SN-Size-DL", Optional: false},
			per.ComponentInfo{Name: "rLC-Mode", Optional: false},
			per.ComponentInfo{Name: "rOHC-Parameters", Optional: true},
			per.ComponentInfo{Name: "t-ReorderingTimer", Optional: true},
			per.ComponentInfo{Name: "discardTimer", Optional: true},
			per.ComponentInfo{Name: "uLDataSplitThreshold", Optional: true},
			per.ComponentInfo{Name: "pDCP-Duplication", Optional: true},
			per.ComponentInfo{Name: "pDCP-Reestablishment", Optional: true},
			per.ComponentInfo{Name: "pDCP-DataRecovery", Optional: true},
			per.ComponentInfo{Name: "duplication-Activation", Optional: true},
			per.ComponentInfo{Name: "outOfOrderDelivery", Optional: true},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode PDCPSNSizeUL failed: %w", err)
		}
		s.PDCPSNSizeUL.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode PDCPSNSizeDL failed: %w", err)
		}
		s.PDCPSNSizeDL.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode RLCMode failed: %w", err)
		}
		s.RLCMode.Value = val
	}

	if seqDecoder.IsComponentPresent(3) {
		s.ROHCParameters = new(ROHCParameters)
		if err = s.ROHCParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode ROHCParameters failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.TReorderingTimer = new(TReorderingTimer)
		if err = s.TReorderingTimer.Decode(r); err != nil {
			return fmt.Errorf("Decode TReorderingTimer failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.DiscardTimer = new(DiscardTimer)

		{
			enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 16), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode DiscardTimer failed: %w", err)
			}
			s.DiscardTimer.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(6) {
		s.ULDataSplitThreshold = new(ULDataSplitThreshold)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 24), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode ULDataSplitThreshold failed: %w", err)
			}
			s.ULDataSplitThreshold.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(7) {
		s.PDCPDuplication = new(PDCPDuplication)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode PDCPDuplication failed: %w", err)
			}
			s.PDCPDuplication.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(8) {
		s.PDCPReestablishment = new(PDCPReestablishment)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode PDCPReestablishment failed: %w", err)
			}
			s.PDCPReestablishment.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(9) {
		s.PDCPDataRecovery = new(PDCPDataRecovery)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode PDCPDataRecovery failed: %w", err)
			}
			s.PDCPDataRecovery.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(10) {
		s.DuplicationActivation = new(DuplicationActivation)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode DuplicationActivation failed: %w", err)
			}
			s.DuplicationActivation.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(11) {
		s.OutOfOrderDelivery = new(OutOfOrderDelivery)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode OutOfOrderDelivery failed: %w", err)
			}
			s.OutOfOrderDelivery.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(12) {
		s.IEExtensions = new(PDCPConfigurationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
