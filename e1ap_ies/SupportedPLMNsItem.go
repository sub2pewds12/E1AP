package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SupportedPLMNsItem is a generated SEQUENCE type.
type SupportedPLMNsItem struct {
	PLMNIdentity             PLMNIdentity                  `aper:"lb:3,ub:3,mandatory,ext"`
	SliceSupportList         *SliceSupportList             `aper:"optional,ext"`
	NRCGISupportList         *NRCGISupportList             `aper:"optional,ext"`
	QOSParametersSupportList *QOSParametersSupportList     `aper:"optional,ext"`
	IEExtensions             *SupportedPLMNsItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *SupportedPLMNsItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.PLMNIdentity.Value), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode PLMNIdentity failed: %w", err)
	}
	if s.SliceSupportList != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.SliceSupportList.Value))
			for i := 0; i < len(s.SliceSupportList.Value); i++ {
				itemPointers[i] = &(s.SliceSupportList.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Encode SliceSupportList failed: %w", err)
			}
		}
	}
	if s.NRCGISupportList != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.NRCGISupportList.Value))
			for i := 0; i < len(s.NRCGISupportList.Value); i++ {
				itemPointers[i] = &(s.NRCGISupportList.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Encode NRCGISupportList failed: %w", err)
			}
		}
	}
	if s.QOSParametersSupportList != nil {
		if err = s.QOSParametersSupportList.Encode(w); err != nil {
			return fmt.Errorf("Encode QOSParametersSupportList failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SupportedPLMNsItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.PLMNIdentity.Decode(r); err != nil {
		return fmt.Errorf("Decode PLMNIdentity failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.SliceSupportList = new(SliceSupportList)
		if err = s.SliceSupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode SliceSupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.NRCGISupportList = new(NRCGISupportList)
		if err = s.NRCGISupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode NRCGISupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.QOSParametersSupportList = new(QOSParametersSupportList)
		if err = s.QOSParametersSupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode QOSParametersSupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(SupportedPLMNsItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for SupportedPLMNsItem */
	}
	return nil
}
