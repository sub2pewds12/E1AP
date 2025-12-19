package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QoSFlowLevelQoSParameters is a generated SEQUENCE type.
type QoSFlowLevelQoSParameters struct {
	QOSCharacteristics               QOSCharacteristics                                 `aper:"mandatory,ext"`
	NGRANallocationRetentionPriority NGRANAllocationAndRetentionPriority                `aper:"mandatory,ext"`
	GBRQOSFlowInformation            *GBRQoSFlowInformation                             `aper:"optional,ext"`
	ReflectiveQOSAttribute           *QoSFlowLevelQoSParametersReflectiveQOSAttribute   `aper:"optional,ext"`
	AdditionalQOSInformation         *QoSFlowLevelQoSParametersAdditionalQOSInformation `aper:"optional,ext"`
	PagingPolicyIndicator            *QoSFlowLevelQoSParametersPagingPolicyIndicator    `aper:"lb:1,ub:8,optional,ext"`
	ReflectiveQOSIndicator           *QoSFlowLevelQoSParametersReflectiveQOSIndicator   `aper:"optional,ext"`
	IEExtensions                     *QoSFlowLevelQoSParametersExtensions               `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *QoSFlowLevelQoSParameters) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GBRQOSFlowInformation != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.ReflectiveQOSAttribute != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.AdditionalQOSInformation != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.PagingPolicyIndicator != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.ReflectiveQOSIndicator != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(6), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = s.QOSCharacteristics.Encode(w); err != nil {
		return fmt.Errorf("encode QOSCharacteristics failed: %w", err)
	}
	if err = s.NGRANallocationRetentionPriority.Encode(w); err != nil {
		return fmt.Errorf("encode NGRANallocationRetentionPriority failed: %w", err)
	}
	if s.GBRQOSFlowInformation != nil {
		if err = s.GBRQOSFlowInformation.Encode(w); err != nil {
			return fmt.Errorf("encode GBRQOSFlowInformation failed: %w", err)
		}
	}
	if s.ReflectiveQOSAttribute != nil {
		if err = w.WriteEnumerate(uint64((*s.ReflectiveQOSAttribute).Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("encode ReflectiveQOSAttribute failed: %w", err)
		}
	}
	if s.AdditionalQOSInformation != nil {
		if err = w.WriteEnumerate(uint64((*s.AdditionalQOSInformation).Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("encode AdditionalQOSInformation failed: %w", err)
		}
	}
	if s.PagingPolicyIndicator != nil {
		if err = w.WriteInteger(int64((*s.PagingPolicyIndicator).Value), &aper.Constraint{Lb: 1, Ub: 8}, true); err != nil {
			return fmt.Errorf("encode PagingPolicyIndicator failed: %w", err)
		}
	}
	if s.ReflectiveQOSIndicator != nil {
		if err = w.WriteEnumerate(uint64((*s.ReflectiveQOSIndicator).Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("encode ReflectiveQOSIndicator failed: %w", err)
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
func (s *QoSFlowLevelQoSParameters) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.QOSCharacteristics.Decode(r); err != nil {
		return fmt.Errorf("decode QOSCharacteristics failed: %w", err)
	}
	if err = s.NGRANallocationRetentionPriority.Decode(r); err != nil {
		return fmt.Errorf("decode NGRANallocationRetentionPriority failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.GBRQOSFlowInformation = new(GBRQoSFlowInformation)
		if err = s.GBRQOSFlowInformation.Decode(r); err != nil {
			return fmt.Errorf("decode GBRQOSFlowInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.ReflectiveQOSAttribute = new(QoSFlowLevelQoSParametersReflectiveQOSAttribute)
		if err = s.ReflectiveQOSAttribute.Decode(r); err != nil {
			return fmt.Errorf("decode ReflectiveQOSAttribute failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.AdditionalQOSInformation = new(QoSFlowLevelQoSParametersAdditionalQOSInformation)
		if err = s.AdditionalQOSInformation.Decode(r); err != nil {
			return fmt.Errorf("decode AdditionalQOSInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.PagingPolicyIndicator = new(QoSFlowLevelQoSParametersPagingPolicyIndicator)
		if err = s.PagingPolicyIndicator.Decode(r); err != nil {
			return fmt.Errorf("decode PagingPolicyIndicator failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.ReflectiveQOSIndicator = new(QoSFlowLevelQoSParametersReflectiveQOSIndicator)
		if err = s.ReflectiveQOSIndicator.Decode(r); err != nil {
			return fmt.Errorf("decode ReflectiveQOSIndicator failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.IEExtensions = new(QoSFlowLevelQoSParametersExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for QoSFlowLevelQoSParameters */
	}
	return nil
}
