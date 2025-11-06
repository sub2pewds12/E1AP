package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDCPConfiguration is a generated SEQUENCE type.
type PDCPConfiguration struct {
	PDCPSNSizeUL          PDCPSNSize                   `aper:"mandatory,ext"`
	PDCPSNSizeDL          PDCPSNSize                   `aper:"mandatory,ext"`
	RLCMode               RLCMode                      `aper:"mandatory,ext"`
	ROHCParameters        *ROHCParameters              `aper:"optional,ext"`
	TReorderingTimer      *TReorderingTimer            `aper:"optional,ext"`
	DiscardTimer          *DiscardTimer                `aper:"optional,ext"`
	ULDataSplitThreshold  *ULDataSplitThreshold        `aper:"optional,ext"`
	PDCPDuplication       *PDCPDuplication             `aper:"optional,ext"`
	PDCPReestablishment   *PDCPReestablishment         `aper:"optional,ext"`
	PDCPDataRecovery      *PDCPDataRecovery            `aper:"optional,ext"`
	DuplicationActivation *DuplicationActivation       `aper:"optional,ext"`
	OutOfOrderDelivery    *OutOfOrderDelivery          `aper:"optional,ext"`
	IEExtensions          *PDCPConfigurationExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDCPConfiguration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [2]byte
	if s.ROHCParameters != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.TReorderingTimer != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.DiscardTimer != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.ULDataSplitThreshold != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.PDCPDuplication != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.PDCPReestablishment != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if s.PDCPDataRecovery != nil {
		optionalityBitmap[0] |= 1 << 1
	}
	if s.DuplicationActivation != nil {
		optionalityBitmap[0] |= 1 << 0
	}
	if s.OutOfOrderDelivery != nil {
		optionalityBitmap[1] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[1] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(10), &aper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = s.PDCPSNSizeUL.Encode(w); err != nil {
		return fmt.Errorf("Encode PDCPSNSizeUL failed: %w", err)
	}
	if err = s.PDCPSNSizeDL.Encode(w); err != nil {
		return fmt.Errorf("Encode PDCPSNSizeDL failed: %w", err)
	}
	if err = s.RLCMode.Encode(w); err != nil {
		return fmt.Errorf("Encode RLCMode failed: %w", err)
	}
	if s.ROHCParameters != nil {
		if err = s.ROHCParameters.Encode(w); err != nil {
			return fmt.Errorf("Encode ROHCParameters failed: %w", err)
		}
	}
	if s.TReorderingTimer != nil {
		if err = s.TReorderingTimer.Encode(w); err != nil {
			return fmt.Errorf("Encode TReorderingTimer failed: %w", err)
		}
	}
	if s.DiscardTimer != nil {
		if err = s.DiscardTimer.Encode(w); err != nil {
			return fmt.Errorf("Encode DiscardTimer failed: %w", err)
		}
	}
	if s.ULDataSplitThreshold != nil {
		if err = s.ULDataSplitThreshold.Encode(w); err != nil {
			return fmt.Errorf("Encode ULDataSplitThreshold failed: %w", err)
		}
	}
	if s.PDCPDuplication != nil {
		if err = s.PDCPDuplication.Encode(w); err != nil {
			return fmt.Errorf("Encode PDCPDuplication failed: %w", err)
		}
	}
	if s.PDCPReestablishment != nil {
		if err = s.PDCPReestablishment.Encode(w); err != nil {
			return fmt.Errorf("Encode PDCPReestablishment failed: %w", err)
		}
	}
	if s.PDCPDataRecovery != nil {
		if err = s.PDCPDataRecovery.Encode(w); err != nil {
			return fmt.Errorf("Encode PDCPDataRecovery failed: %w", err)
		}
	}
	if s.DuplicationActivation != nil {
		if err = s.DuplicationActivation.Encode(w); err != nil {
			return fmt.Errorf("Encode DuplicationActivation failed: %w", err)
		}
	}
	if s.OutOfOrderDelivery != nil {
		if err = s.OutOfOrderDelivery.Encode(w); err != nil {
			return fmt.Errorf("Encode OutOfOrderDelivery failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDCPConfiguration) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
			return fmt.Errorf("Decode PDCPSNSizeUL failed: %w", err)
		}
		s.PDCPSNSizeUL.Value = aper.Enumerated(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
			return fmt.Errorf("Decode PDCPSNSizeDL failed: %w", err)
		}
		s.PDCPSNSizeDL.Value = aper.Enumerated(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true); err != nil {
			return fmt.Errorf("Decode RLCMode failed: %w", err)
		}
		s.RLCMode.Value = aper.Enumerated(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ROHCParameters = new(ROHCParameters)
		if err = s.ROHCParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode ROHCParameters failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.TReorderingTimer = new(TReorderingTimer)
		if err = s.TReorderingTimer.Decode(r); err != nil {
			return fmt.Errorf("Decode TReorderingTimer failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.DiscardTimer = new(DiscardTimer)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
				return fmt.Errorf("Decode DiscardTimer failed: %w", err)
			}
			s.DiscardTimer.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.ULDataSplitThreshold = new(ULDataSplitThreshold)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 23}, true); err != nil {
				return fmt.Errorf("Decode ULDataSplitThreshold failed: %w", err)
			}
			s.ULDataSplitThreshold.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.PDCPDuplication = new(PDCPDuplication)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode PDCPDuplication failed: %w", err)
			}
			s.PDCPDuplication.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.PDCPReestablishment = new(PDCPReestablishment)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode PDCPReestablishment failed: %w", err)
			}
			s.PDCPReestablishment.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<1) > 0 {
		s.PDCPDataRecovery = new(PDCPDataRecovery)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode PDCPDataRecovery failed: %w", err)
			}
			s.PDCPDataRecovery.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<0) > 0 {
		s.DuplicationActivation = new(DuplicationActivation)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return fmt.Errorf("Decode DuplicationActivation failed: %w", err)
			}
			s.DuplicationActivation.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<7) > 0 {
		s.OutOfOrderDelivery = new(OutOfOrderDelivery)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode OutOfOrderDelivery failed: %w", err)
			}
			s.OutOfOrderDelivery.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<6) > 0 {
		s.IEExtensions = new(PDCPConfigurationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for PDCPConfiguration")
	}
	return nil
}
