package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSMappingInformation is a generated SEQUENCE type.
type QOSMappingInformation struct {
	Dscp      *QOSMappingInformationDscp      `aper:"lb:6,ub:6,optional,ext"`
	FlowLabel *QOSMappingInformationFlowLabel `aper:"lb:20,ub:20,optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSMappingInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.Dscp != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.FlowLabel != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.Dscp != nil {
		if err = w.WriteBitString((*s.Dscp).Bytes, uint((*s.Dscp).NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return fmt.Errorf("Encode Dscp failed: %w", err)
		}
	}
	if s.FlowLabel != nil {
		if err = w.WriteBitString((*s.FlowLabel).Bytes, uint((*s.FlowLabel).NumBits), &aper.Constraint{Lb: 20, Ub: 20}, false); err != nil {
			return fmt.Errorf("Encode FlowLabel failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSMappingInformation) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
				return fmt.Errorf("Decode Dscp failed: %w", err)
			}
			tmp := QOSMappingInformationDscp(val)
			s.Dscp = &tmp
		}

	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 20, Ub: 20}, false); err != nil {
				return fmt.Errorf("Decode FlowLabel failed: %w", err)
			}
			tmp := QOSMappingInformationFlowLabel(val)
			s.FlowLabel = &tmp
		}

	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for QOSMappingInformation")
	}
	return nil
}
