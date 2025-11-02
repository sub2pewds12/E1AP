package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSFlowQOSParameterItem is a generated SEQUENCE type.
type QOSFlowQOSParameterItem struct {
	QOSFlowIdentifier         QOSFlowIdentifier                  `aper:"lb:0,ub:63,mandatory,ext"`
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters          `aper:"mandatory,ext"`
	QoSFlowMappingIndication  *QOSFlowMappingIndication          `aper:"optional,ext"`
	IEExtensions              *QOSFlowQOSParameterItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowQOSParameterItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.QoSFlowMappingIndication != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.QOSFlowIdentifier.Value), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return fmt.Errorf("Encode QOSFlowIdentifier failed: %w", err)
	}
	if err = s.QoSFlowLevelQoSParameters.Encode(w); err != nil {
		return fmt.Errorf("Encode QoSFlowLevelQoSParameters failed: %w", err)
	}
	if s.QoSFlowMappingIndication != nil {
		if err = w.WriteEnumerate(uint64(s.QoSFlowMappingIndication.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
			return fmt.Errorf("Encode QoSFlowMappingIndication failed: %w", err)
		}
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSFlowQOSParameterItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return fmt.Errorf("Decode QOSFlowIdentifier failed: %w", err)
		}
		s.QOSFlowIdentifier.Value = aper.Integer(val)
	}
	if err = s.QoSFlowLevelQoSParameters.Decode(r); err != nil {
		return fmt.Errorf("Decode QoSFlowLevelQoSParameters failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.QoSFlowMappingIndication = new(QOSFlowMappingIndication)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
				return fmt.Errorf("Decode QoSFlowMappingIndication failed: %w", err)
			}
			s.QoSFlowMappingIndication.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(QOSFlowQOSParameterItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for QOSFlowQOSParameterItem")
	}
	return nil
}
