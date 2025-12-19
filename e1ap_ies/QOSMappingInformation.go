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
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.Dscp != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.FlowLabel != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if s.Dscp != nil {
		if err = w.WriteBitString((*s.Dscp).Value.Bytes, uint((*s.Dscp).Value.NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return fmt.Errorf("encode Dscp failed: %w", err)
		}
	}
	if s.FlowLabel != nil {
		if err = w.WriteBitString((*s.FlowLabel).Value.Bytes, uint((*s.FlowLabel).Value.NumBits), &aper.Constraint{Lb: 20, Ub: 20}, false); err != nil {
			return fmt.Errorf("encode FlowLabel failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSMappingInformation) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.Dscp = new(QOSMappingInformationDscp)
		if err = s.Dscp.Decode(r); err != nil {
			return fmt.Errorf("decode Dscp failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.FlowLabel = new(QOSMappingInformationFlowLabel)
		if err = s.FlowLabel.Decode(r); err != nil {
			return fmt.Errorf("decode FlowLabel failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for QOSMappingInformation */
	}
	return nil
}
