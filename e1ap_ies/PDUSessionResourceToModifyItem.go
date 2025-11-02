package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceToModifyItem struct {
	PDUSessionID                               PDUSessionID                              `aper:"lb:0,ub:255,mandatory,ext"`
	SecurityIndication                         *SecurityIndication                       `aper:"optional,ext"`
	PDUSessionResourceDLAMBR                   *BitRate                                  `aper:"lb:0,ub:4000000000000,optional,ext"`
	NGULUPTNLInformation                       *UPTNLInformation                         `aper:"optional,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest         `aper:"optional,ext"`
	PDUSessionDataForwardingInformation        *DataForwardingInformation                `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *InactivityTimer                          `aper:"lb:1,ub:7200,optional,ext"`
	NetworkInstance                            *NetworkInstance                          `aper:"lb:1,ub:256,optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN                     `aper:"optional,ext"`
	DRBToModifyListNGRAN                       []DRBToModifyItemNGRAN                    `aper:"optional,ext"`
	DRBToRemoveListNGRAN                       []DRBToRemoveItemNGRAN                    `aper:"optional,ext"`
	IEExtensions                               *PDUSessionResourceToModifyItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToModifyItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [2]byte
	if s.SecurityIndication != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.PDUSessionResourceDLAMBR != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.NGULUPTNLInformation != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.PDUSessionDataForwardingInformationRequest != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.PDUSessionDataForwardingInformation != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.PDUSessionInactivityTimer != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if s.NetworkInstance != nil {
		optionalityBitmap[0] |= 1 << 1
	}
	if s.DRBToSetupListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 0
	}
	if s.DRBToModifyListNGRAN != nil {
		optionalityBitmap[1] |= 1 << 7
	}
	if s.DRBToRemoveListNGRAN != nil {
		optionalityBitmap[1] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[1] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(11), &aper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PDUSessionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return fmt.Errorf("Encode PDUSessionID failed: %w", err)
	}
	if s.SecurityIndication != nil {
		if err = s.SecurityIndication.Encode(w); err != nil {
			return fmt.Errorf("Encode SecurityIndication failed: %w", err)
		}
	}
	if s.PDUSessionResourceDLAMBR != nil {
		if err = w.WriteInteger(int64(s.PDUSessionResourceDLAMBR.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode PDUSessionResourceDLAMBR failed: %w", err)
		}
	}
	if s.NGULUPTNLInformation != nil {
		if err = s.NGULUPTNLInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode NGULUPTNLInformation failed: %w", err)
		}
	}
	if s.PDUSessionDataForwardingInformationRequest != nil {
		if err = s.PDUSessionDataForwardingInformationRequest.Encode(w); err != nil {
			return fmt.Errorf("Encode PDUSessionDataForwardingInformationRequest failed: %w", err)
		}
	}
	if s.PDUSessionDataForwardingInformation != nil {
		if err = s.PDUSessionDataForwardingInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode PDUSessionDataForwardingInformation failed: %w", err)
		}
	}
	if s.PDUSessionInactivityTimer != nil {
		if err = w.WriteInteger(int64(s.PDUSessionInactivityTimer.Value), &aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
			return fmt.Errorf("Encode PDUSessionInactivityTimer failed: %w", err)
		}
	}
	if s.NetworkInstance != nil {
		if err = w.WriteInteger(int64(s.NetworkInstance.Value), &aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
			return fmt.Errorf("Encode NetworkInstance failed: %w", err)
		}
	}
	if s.DRBToSetupListNGRAN != nil {
		if err = s.DRBToSetupListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToSetupListNGRAN failed: %w", err)
		}
	}
	if s.DRBToModifyListNGRAN != nil {
		if err = s.DRBToModifyListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToModifyListNGRAN failed: %w", err)
		}
	}
	if s.DRBToRemoveListNGRAN != nil {
		if err = s.DRBToRemoveListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToRemoveListNGRAN failed: %w", err)
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
func (s *PDUSessionResourceToModifyItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return fmt.Errorf("Decode PDUSessionID failed: %w", err)
		}
		s.PDUSessionID.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.SecurityIndication = new(SecurityIndication)
		if err = s.SecurityIndication.Decode(r); err != nil {
			return fmt.Errorf("Decode SecurityIndication failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode PDUSessionResourceDLAMBR failed: %w", err)
			}
			s.PDUSessionResourceDLAMBR = new(BitRate)
			s.PDUSessionResourceDLAMBR.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.NGULUPTNLInformation = new(UPTNLInformation)
		if err = s.NGULUPTNLInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode NGULUPTNLInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.PDUSessionDataForwardingInformationRequest = new(DataForwardingInformationRequest)
		if err = s.PDUSessionDataForwardingInformationRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode PDUSessionDataForwardingInformationRequest failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.PDUSessionDataForwardingInformation = new(DataForwardingInformation)
		if err = s.PDUSessionDataForwardingInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode PDUSessionDataForwardingInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
				return fmt.Errorf("Decode PDUSessionInactivityTimer failed: %w", err)
			}
			s.PDUSessionInactivityTimer = new(InactivityTimer)
			s.PDUSessionInactivityTimer.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<1) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
				return fmt.Errorf("Decode NetworkInstance failed: %w", err)
			}
			s.NetworkInstance = new(NetworkInstance)
			s.NetworkInstance.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<0) > 0 {
		s.DRBToSetupListNGRAN = new(DRBToSetupListNGRAN)
		if err = s.DRBToSetupListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBToSetupListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<7) > 0 {
		s.DRBToModifyListNGRAN = new(DRBToModifyListNGRAN)
		if err = s.DRBToModifyListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBToModifyListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<6) > 0 {
		s.DRBToRemoveListNGRAN = new(DRBToRemoveListNGRAN)
		if err = s.DRBToRemoveListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBToRemoveListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 1 && optionalityBitmap[1]&(1<<5) > 0 {
		s.IEExtensions = new(PDUSessionResourceToModifyItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for PDUSessionResourceToModifyItem")
	}
	return nil
}
