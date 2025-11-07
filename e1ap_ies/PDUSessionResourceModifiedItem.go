package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceModifiedItem is a generated SEQUENCE type.
type PDUSessionResourceModifiedItem struct {
	PDUSessionID                                PDUSessionID                              `aper:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation                        *UPTNLInformation                         `aper:"optional,ext"`
	SecurityResult                              *SecurityResult                           `aper:"optional,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation                `aper:"optional,ext"`
	DRBSetupListNGRAN                           *DRBSetupListNGRAN                        `aper:"optional,ext"`
	DRBFailedListNGRAN                          *DRBFailedListNGRAN                       `aper:"optional,ext"`
	DRBModifiedListNGRAN                        *DRBModifiedListNGRAN                     `aper:"optional,ext"`
	DRBFailedToModifyListNGRAN                  *DRBFailedToModifyListNGRAN               `aper:"optional,ext"`
	IEExtensions                                *PDUSessionResourceModifiedItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceModifiedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.NGDLUPTNLInformation != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.SecurityResult != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.PDUSessionDataForwardingInformationResponse != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.DRBSetupListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.DRBFailedListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.DRBModifiedListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if s.DRBFailedToModifyListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 1
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 0
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(8), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PDUSessionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return fmt.Errorf("Encode PDUSessionID failed: %w", err)
	}
	if s.NGDLUPTNLInformation != nil {
		if err = s.NGDLUPTNLInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode NGDLUPTNLInformation failed: %w", err)
		}
	}
	if s.SecurityResult != nil {
		if err = s.SecurityResult.Encode(w); err != nil {
			return fmt.Errorf("Encode SecurityResult failed: %w", err)
		}
	}
	if s.PDUSessionDataForwardingInformationResponse != nil {
		if err = s.PDUSessionDataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("Encode PDUSessionDataForwardingInformationResponse failed: %w", err)
		}
	}
	if s.DRBSetupListNGRAN != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DRBSetupListNGRAN.Value))
			for i := 0; i < len(s.DRBSetupListNGRAN.Value); i++ {
				itemPointers[i] = &(s.DRBSetupListNGRAN.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Encode DRBSetupListNGRAN failed: %w", err)
			}
		}
	}
	if s.DRBFailedListNGRAN != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DRBFailedListNGRAN.Value))
			for i := 0; i < len(s.DRBFailedListNGRAN.Value); i++ {
				itemPointers[i] = &(s.DRBFailedListNGRAN.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Encode DRBFailedListNGRAN failed: %w", err)
			}
		}
	}
	if s.DRBModifiedListNGRAN != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DRBModifiedListNGRAN.Value))
			for i := 0; i < len(s.DRBModifiedListNGRAN.Value); i++ {
				itemPointers[i] = &(s.DRBModifiedListNGRAN.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Encode DRBModifiedListNGRAN failed: %w", err)
			}
		}
	}
	if s.DRBFailedToModifyListNGRAN != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DRBFailedToModifyListNGRAN.Value))
			for i := 0; i < len(s.DRBFailedToModifyListNGRAN.Value); i++ {
				itemPointers[i] = &(s.DRBFailedToModifyListNGRAN.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Encode DRBFailedToModifyListNGRAN failed: %w", err)
			}
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceModifiedItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.PDUSessionID.Decode(r); err != nil {
		return fmt.Errorf("Decode PDUSessionID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.NGDLUPTNLInformation = new(UPTNLInformation)
		if err = s.NGDLUPTNLInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode NGDLUPTNLInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.SecurityResult = new(SecurityResult)
		if err = s.SecurityResult.Decode(r); err != nil {
			return fmt.Errorf("Decode SecurityResult failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.PDUSessionDataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.PDUSessionDataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("Decode PDUSessionDataForwardingInformationResponse failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.DRBSetupListNGRAN = new(DRBSetupListNGRAN)
		if err = s.DRBSetupListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBSetupListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.DRBFailedListNGRAN = new(DRBFailedListNGRAN)
		if err = s.DRBFailedListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.DRBModifiedListNGRAN = new(DRBModifiedListNGRAN)
		if err = s.DRBModifiedListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBModifiedListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<1) > 0 {
		s.DRBFailedToModifyListNGRAN = new(DRBFailedToModifyListNGRAN)
		if err = s.DRBFailedToModifyListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedToModifyListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<0) > 0 {
		s.IEExtensions = new(PDUSessionResourceModifiedItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for PDUSessionResourceModifiedItem")
	}
	return nil
}
