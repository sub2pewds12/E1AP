package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceSetupModItem is a generated SEQUENCE type.
type PDUSessionResourceSetupModItem struct {
	PDUSessionID                                PDUSessionID                              `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult                           `aper:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation                          `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation                `aper:"optional,ext"`
	DRBSetupModListNGRAN                        []DRBSetupModItemNGRAN                    `aper:"mandatory,ext"`
	DRBFailedModListNGRAN                       []DRBFailedModItemNGRAN                   `aper:"optional,ext"`
	IEExtensions                                *PDUSessionResourceSetupModItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceSetupModItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.SecurityResult != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.PDUSessionDataForwardingInformationResponse != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.DRBFailedModListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PDUSessionID), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return fmt.Errorf("Encode PDUSessionID failed: %w", err)
	}
	if s.SecurityResult != nil {
		if err = s.SecurityResult.Encode(w); err != nil {
			return fmt.Errorf("Encode SecurityResult failed: %w", err)
		}
	}
	if err = s.NGDLUPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("Encode NGDLUPTNLInformation failed: %w", err)
	}
	if s.PDUSessionDataForwardingInformationResponse != nil {
		if err = s.PDUSessionDataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("Encode PDUSessionDataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.DRBSetupModListNGRAN.Encode(w); err != nil {
		return fmt.Errorf("Encode DRBSetupModListNGRAN failed: %w", err)
	}
	if s.DRBFailedModListNGRAN != nil {
		if err = s.DRBFailedModListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBFailedModListNGRAN failed: %w", err)
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
func (s *PDUSessionResourceSetupModItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return fmt.Errorf("Decode PDUSessionID failed: %w", err)
		}
		s.PDUSessionID = PDUSessionID(val)
	}

	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.SecurityResult = new(SecurityResult)
		if err = s.SecurityResult.Decode(r); err != nil {
			return fmt.Errorf("Decode SecurityResult failed: %w", err)
		}
	}
	if err = s.NGDLUPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode NGDLUPTNLInformation failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.PDUSessionDataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.PDUSessionDataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("Decode PDUSessionDataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.DRBSetupModListNGRAN.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBSetupModListNGRAN failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.DRBFailedModListNGRAN = new(DRBFailedModListNGRAN)
		if err = s.DRBFailedModListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedModListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(PDUSessionResourceSetupModItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for PDUSessionResourceSetupModItem")
	}
	return nil
}
