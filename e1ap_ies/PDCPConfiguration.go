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
		return fmt.Errorf("encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.PDCPSNSizeUL.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
		return fmt.Errorf("encode PDCPSNSizeUL failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.PDCPSNSizeDL.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
		return fmt.Errorf("encode PDCPSNSizeDL failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.RLCMode.Value), aper.Constraint{Lb: 0, Ub: 4}, true); err != nil {
		return fmt.Errorf("encode RLCMode failed: %w", err)
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
		if err = w.WriteEnumerate(uint64((*s.DiscardTimer).Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return fmt.Errorf("encode DiscardTimer failed: %w", err)
		}
	}
	if s.ULDataSplitThreshold != nil {
		if err = w.WriteEnumerate(uint64((*s.ULDataSplitThreshold).Value), aper.Constraint{Lb: 0, Ub: 23}, true); err != nil {
			return fmt.Errorf("encode ULDataSplitThreshold failed: %w", err)
		}
	}
	if s.PDCPDuplication != nil {
		if err = w.WriteEnumerate(uint64((*s.PDCPDuplication).Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("encode PDCPDuplication failed: %w", err)
		}
	}
	if s.PDCPReestablishment != nil {
		if err = w.WriteEnumerate(uint64((*s.PDCPReestablishment).Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("encode PDCPReestablishment failed: %w", err)
		}
	}
	if s.PDCPDataRecovery != nil {
		if err = w.WriteEnumerate(uint64((*s.PDCPDataRecovery).Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("encode PDCPDataRecovery failed: %w", err)
		}
	}
	if s.DuplicationActivation != nil {
		if err = w.WriteEnumerate(uint64((*s.DuplicationActivation).Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
			return fmt.Errorf("encode DuplicationActivation failed: %w", err)
		}
	}
	if s.OutOfOrderDelivery != nil {
		if err = w.WriteEnumerate(uint64((*s.OutOfOrderDelivery).Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("encode OutOfOrderDelivery failed: %w", err)
		}
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDCPConfiguration) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 10, Ub: 10}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.PDCPSNSizeUL.Decode(r); err != nil {
		return fmt.Errorf("decode PDCPSNSizeUL failed: %w", err)
	}
	if err = s.PDCPSNSizeDL.Decode(r); err != nil {
		return fmt.Errorf("decode PDCPSNSizeDL failed: %w", err)
	}
	if err = s.RLCMode.Decode(r); err != nil {
		return fmt.Errorf("decode RLCMode failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ROHCParameters = new(ROHCParameters)
		if err = s.ROHCParameters.Decode(r); err != nil {
			return fmt.Errorf("decode ROHCParameters failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.TReorderingTimer = new(TReorderingTimer)
		if err = s.TReorderingTimer.Decode(r); err != nil {
			return fmt.Errorf("decode TReorderingTimer failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.DiscardTimer = new(DiscardTimer)
		if err = s.DiscardTimer.Decode(r); err != nil {
			return fmt.Errorf("decode DiscardTimer failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.ULDataSplitThreshold = new(ULDataSplitThreshold)
		if err = s.ULDataSplitThreshold.Decode(r); err != nil {
			return fmt.Errorf("decode ULDataSplitThreshold failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.PDCPDuplication = new(PDCPDuplication)
		if err = s.PDCPDuplication.Decode(r); err != nil {
			return fmt.Errorf("decode PDCPDuplication failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.PDCPReestablishment = new(PDCPReestablishment)
		if err = s.PDCPReestablishment.Decode(r); err != nil {
			return fmt.Errorf("decode PDCPReestablishment failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<1) > 0 {
		s.PDCPDataRecovery = new(PDCPDataRecovery)
		if err = s.PDCPDataRecovery.Decode(r); err != nil {
			return fmt.Errorf("decode PDCPDataRecovery failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<0) > 0 {
		s.DuplicationActivation = new(DuplicationActivation)
		if err = s.DuplicationActivation.Decode(r); err != nil {
			return fmt.Errorf("decode DuplicationActivation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<7) > 0 {
		s.OutOfOrderDelivery = new(OutOfOrderDelivery)
		if err = s.OutOfOrderDelivery.Decode(r); err != nil {
			return fmt.Errorf("decode OutOfOrderDelivery failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<6) > 0 {
		s.IEExtensions = new(PDCPConfigurationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for PDCPConfiguration */
	}
	return nil
}
