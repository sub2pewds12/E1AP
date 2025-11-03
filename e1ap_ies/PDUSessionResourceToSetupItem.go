package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceToSetupItem is a generated SEQUENCE type.
type PDUSessionResourceToSetupItem struct {
	PDUSessionID                               PDUSessionID                             `aper:"lb:0,ub:255,mandatory,ext"`
	PDUSessionType                             PDUSessionType                           `aper:"mandatory,ext"`
	SNSSAI                                     SNSSAI                                   `aper:"mandatory,ext"`
	SecurityIndication                         SecurityIndication                       `aper:"mandatory,ext"`
	PDUSessionResourceDLAMBR                   *BitRate                                 `aper:"lb:0,ub:4000000000000,optional,ext"`
	NGULUPTNLInformation                       UPTNLInformation                         `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest        `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *InactivityTimer                         `aper:"lb:1,ub:7200,optional,ext"`
	ExistingAllocatedNGDLUPTNLInfo             *UPTNLInformation                        `aper:"optional,ext"`
	NetworkInstance                            *NetworkInstance                         `aper:"lb:1,ub:256,optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN                    `aper:"mandatory,ext"`
	IEExtensions                               *PDUSessionResourceToSetupItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToSetupItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.PDUSessionResourceDLAMBR != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.PDUSessionDataForwardingInformationRequest != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.PDUSessionInactivityTimer != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.ExistingAllocatedNGDLUPTNLInfo != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.NetworkInstance != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 2
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(6), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PDUSessionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return fmt.Errorf("Encode PDUSessionID failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.PDUSessionType.Value), aper.Constraint{Lb: 0, Ub: 4}, true); err != nil {
		return fmt.Errorf("Encode PDUSessionType failed: %w", err)
	}
	if err = s.SNSSAI.Encode(w); err != nil {
		return fmt.Errorf("Encode SNSSAI failed: %w", err)
	}
	if err = s.SecurityIndication.Encode(w); err != nil {
		return fmt.Errorf("Encode SecurityIndication failed: %w", err)
	}
	if s.PDUSessionResourceDLAMBR != nil {
		if err = w.WriteInteger(int64(s.PDUSessionResourceDLAMBR.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode PDUSessionResourceDLAMBR failed: %w", err)
		}
	}
	if err = s.NGULUPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("Encode NGULUPTNLInformation failed: %w", err)
	}
	if s.PDUSessionDataForwardingInformationRequest != nil {
		if err = s.PDUSessionDataForwardingInformationRequest.Encode(w); err != nil {
			return fmt.Errorf("Encode PDUSessionDataForwardingInformationRequest failed: %w", err)
		}
	}
	if s.PDUSessionInactivityTimer != nil {
		if err = w.WriteInteger(int64(s.PDUSessionInactivityTimer.Value), &aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
			return fmt.Errorf("Encode PDUSessionInactivityTimer failed: %w", err)
		}
	}
	if s.ExistingAllocatedNGDLUPTNLInfo != nil {
		if err = s.ExistingAllocatedNGDLUPTNLInfo.Encode(w); err != nil {
			return fmt.Errorf("Encode ExistingAllocatedNGDLUPTNLInfo failed: %w", err)
		}
	}
	if s.NetworkInstance != nil {
		if err = w.WriteInteger(int64(s.NetworkInstance.Value), &aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
			return fmt.Errorf("Encode NetworkInstance failed: %w", err)
		}
	}
	if err = s.DRBToSetupListNGRAN.Encode(w); err != nil {
		return fmt.Errorf("Encode DRBToSetupListNGRAN failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceToSetupItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return fmt.Errorf("Decode PDUSessionID failed: %w", err)
		}
		s.PDUSessionID.Value = aper.Integer(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true); err != nil {
			return fmt.Errorf("Decode PDUSessionType failed: %w", err)
		}
		s.PDUSessionType.Value = aper.Enumerated(val)
	}
	if err = s.SNSSAI.Decode(r); err != nil {
		return fmt.Errorf("Decode SNSSAI failed: %w", err)
	}
	if err = s.SecurityIndication.Decode(r); err != nil {
		return fmt.Errorf("Decode SecurityIndication failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode PDUSessionResourceDLAMBR failed: %w", err)
			}
			s.PDUSessionResourceDLAMBR = new(BitRate)
			s.PDUSessionResourceDLAMBR.Value = aper.Integer(val)
		}
	}
	if err = s.NGULUPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode NGULUPTNLInformation failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.PDUSessionDataForwardingInformationRequest = new(DataForwardingInformationRequest)
		if err = s.PDUSessionDataForwardingInformationRequest.Decode(r); err != nil {
			return fmt.Errorf("Decode PDUSessionDataForwardingInformationRequest failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
				return fmt.Errorf("Decode PDUSessionInactivityTimer failed: %w", err)
			}
			s.PDUSessionInactivityTimer = new(InactivityTimer)
			s.PDUSessionInactivityTimer.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.ExistingAllocatedNGDLUPTNLInfo = new(UPTNLInformation)
		if err = s.ExistingAllocatedNGDLUPTNLInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode ExistingAllocatedNGDLUPTNLInfo failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
				return fmt.Errorf("Decode NetworkInstance failed: %w", err)
			}
			s.NetworkInstance = new(NetworkInstance)
			s.NetworkInstance.Value = aper.Integer(val)
		}
	}
	if err = s.DRBToSetupListNGRAN.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBToSetupListNGRAN failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<2) > 0 {
		s.IEExtensions = new(PDUSessionResourceToSetupItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for PDUSessionResourceToSetupItem")
	}
	return nil
}
