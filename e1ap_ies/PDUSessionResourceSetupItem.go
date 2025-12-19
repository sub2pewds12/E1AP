package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceSetupItem is a generated SEQUENCE type.
type PDUSessionResourceSetupItem struct {
	PDUSessionID                                PDUSessionID                                `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityResult                              *SecurityResult                             `aper:"optional,ext"`
	NGDLUPTNLInformation                        UPTNLInformation                            `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationResponse *DataForwardingInformation                  `aper:"optional,ext"`
	NGDLUPUnchanged                             *PDUSessionResourceSetupItemNGDLUPUnchanged `aper:"optional,ext"`
	DRBSetupListNGRAN                           DRBSetupListNGRAN                           `aper:"lb:1,ub:MaxnoofDRBs,mandatory,ext"`
	DRBFailedListNGRAN                          *DRBFailedListNGRAN                         `aper:"lb:1,ub:MaxnoofDRBs,optional,ext"`
	IEExtensions                                *PDUSessionResourceSetupItemExtensions      `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceSetupItem) Encode(w *aper.AperWriter) (err error) {
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
	if s.NGDLUPUnchanged != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.DRBFailedListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(5), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PDUSessionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
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
	if s.NGDLUPUnchanged != nil {
		if err = w.WriteEnumerate(uint64((*s.NGDLUPUnchanged).Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("Encode NGDLUPUnchanged failed: %w", err)
		}
	}
	if err = s.DRBSetupListNGRAN.Encode(w); err != nil {
		return fmt.Errorf("Encode DRBSetupListNGRAN failed: %w", err)
	}
	if s.DRBFailedListNGRAN != nil {
		if err = s.DRBFailedListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBFailedListNGRAN failed: %w", err)
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
func (s *PDUSessionResourceSetupItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.PDUSessionID.Decode(r); err != nil {
		return fmt.Errorf("Decode PDUSessionID failed: %w", err)
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
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.NGDLUPUnchanged = new(PDUSessionResourceSetupItemNGDLUPUnchanged)
		if err = s.NGDLUPUnchanged.Decode(r); err != nil {
			return fmt.Errorf("Decode NGDLUPUnchanged failed: %w", err)
		}
	}
	if err = s.DRBSetupListNGRAN.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBSetupListNGRAN failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.DRBFailedListNGRAN = new(DRBFailedListNGRAN)
		if err = s.DRBFailedListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(PDUSessionResourceSetupItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for PDUSessionResourceSetupItem */
	}
	return nil
}
