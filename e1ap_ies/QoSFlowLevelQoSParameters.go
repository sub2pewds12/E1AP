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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = s.QOSCharacteristics.Encode(w); err != nil {
		return fmt.Errorf("Encode QOSCharacteristics failed: %w", err)
	}
	if err = s.NGRANallocationRetentionPriority.Encode(w); err != nil {
		return fmt.Errorf("Encode NGRANallocationRetentionPriority failed: %w", err)
	}
	if s.GBRQOSFlowInformation != nil {
		if err = s.GBRQOSFlowInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode GBRQOSFlowInformation failed: %w", err)
		}
	}
	if s.ReflectiveQOSAttribute != nil {
		if err = w.WriteEnumerate(uint64(s.ReflectiveQOSAttribute.Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("Encode ReflectiveQOSAttribute failed: %w", err)
		}
	}
	if s.AdditionalQOSInformation != nil {
		if err = w.WriteEnumerate(uint64(s.AdditionalQOSInformation.Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("Encode AdditionalQOSInformation failed: %w", err)
		}
	}
	if s.PagingPolicyIndicator != nil {
		if err = w.WriteInteger(int64(s.PagingPolicyIndicator.Value), &aper.Constraint{Lb: 1, Ub: 8}, true); err != nil {
			return fmt.Errorf("Encode PagingPolicyIndicator failed: %w", err)
		}
	}
	if s.ReflectiveQOSIndicator != nil {
		if err = w.WriteEnumerate(uint64(s.ReflectiveQOSIndicator.Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("Encode ReflectiveQOSIndicator failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QoSFlowLevelQoSParameters) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.QOSCharacteristics.Decode(r); err != nil {
		return fmt.Errorf("Decode QOSCharacteristics failed: %w", err)
	}
	if err = s.NGRANallocationRetentionPriority.Decode(r); err != nil {
		return fmt.Errorf("Decode NGRANallocationRetentionPriority failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.GBRQOSFlowInformation = new(GBRQoSFlowInformation)
		if err = s.GBRQOSFlowInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode GBRQOSFlowInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.ReflectiveQOSAttribute = new(QoSFlowLevelQoSParametersReflectiveQOSAttribute)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode ReflectiveQOSAttribute failed: %w", err)
			}
			s.ReflectiveQOSAttribute.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.AdditionalQOSInformation = new(QoSFlowLevelQoSParametersAdditionalQOSInformation)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode AdditionalQOSInformation failed: %w", err)
			}
			s.AdditionalQOSInformation.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, true); err != nil {
				return fmt.Errorf("Decode PagingPolicyIndicator failed: %w", err)
			}
			s.PagingPolicyIndicator = new(QoSFlowLevelQoSParametersPagingPolicyIndicator)
			s.PagingPolicyIndicator.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.ReflectiveQOSIndicator = new(QoSFlowLevelQoSParametersReflectiveQOSIndicator)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode ReflectiveQOSIndicator failed: %w", err)
			}
			s.ReflectiveQOSIndicator.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.IEExtensions = new(QoSFlowLevelQoSParametersExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for QoSFlowLevelQoSParameters")
	}
	return nil
}
