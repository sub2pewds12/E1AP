package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SupportedPLMNsItem is a generated SEQUENCE type.
type SupportedPLMNsItem struct {
	PLMNIdentity             PLMNIdentity                  `aper:"lb:3,ub:3,mandatory,ext"`
	SliceSupportList         *SliceSupportList             `aper:"lb:1,ub:MaxnoofSliceItems,optional,ext"`
	NRCGISupportList         *NRCGISupportList             `aper:"lb:1,ub:MaxnoofNRCGI,optional,ext"`
	QOSParametersSupportList *QOSParametersSupportList     `aper:"optional,ext"`
	IEExtensions             *SupportedPLMNsItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *SupportedPLMNsItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.SliceSupportList != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.NRCGISupportList != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.QOSParametersSupportList != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.PLMNIdentity.Value), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("encode PLMNIdentity failed: %w", err)
	}
	if s.SliceSupportList != nil {
		if err = s.SliceSupportList.Encode(w); err != nil {
			return fmt.Errorf("encode SliceSupportList failed: %w", err)
		}
	}
	if s.NRCGISupportList != nil {
		if err = s.NRCGISupportList.Encode(w); err != nil {
			return fmt.Errorf("encode NRCGISupportList failed: %w", err)
		}
	}
	if s.QOSParametersSupportList != nil {
		if err = s.QOSParametersSupportList.Encode(w); err != nil {
			return fmt.Errorf("encode QOSParametersSupportList failed: %w", err)
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
func (s *SupportedPLMNsItem) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.PLMNIdentity.Decode(r); err != nil {
		return fmt.Errorf("decode PLMNIdentity failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.SliceSupportList = new(SliceSupportList)
		if err = s.SliceSupportList.Decode(r); err != nil {
			return fmt.Errorf("decode SliceSupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.NRCGISupportList = new(NRCGISupportList)
		if err = s.NRCGISupportList.Decode(r); err != nil {
			return fmt.Errorf("decode NRCGISupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.QOSParametersSupportList = new(QOSParametersSupportList)
		if err = s.QOSParametersSupportList.Decode(r); err != nil {
			return fmt.Errorf("decode QOSParametersSupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(SupportedPLMNsItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for SupportedPLMNsItem */
	}
	return nil
}
