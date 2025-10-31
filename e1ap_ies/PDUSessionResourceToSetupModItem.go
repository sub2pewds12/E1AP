package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceToSetupModItem is a generated SEQUENCE type.
type PDUSessionResourceToSetupModItem struct {
	PDUSessionID                               PDUSessionID                                `aper:"lb:0,ub:255,mandatory,ext"`
	PDUSessionType                             PDUSessionType                              `aper:"mandatory,ext"`
	SNSSAI                                     SNSSAI                                      `aper:"mandatory,ext"`
	SecurityIndication                         SecurityIndication                          `aper:"mandatory,ext"`
	PDUSessionResourceAMBR                     *BitRate                                    `aper:"lb:0,ub:4000000000000,optional,ext"`
	NGULUPTNLInformation                       UPTNLInformation                            `aper:"mandatory,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest           `aper:"optional,ext"`
	PDUSessionInactivityTimer                  *InactivityTimer                            `aper:"lb:1,ub:7200,optional,ext"`
	DRBToSetupModListNGRAN                     []DRBToSetupModItemNGRAN                    `aper:"mandatory,ext"`
	IEExtensions                               *PDUSessionResourceToSetupModItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToSetupModItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.PDUSessionResourceAMBR != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.PDUSessionDataForwardingInformationRequest != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.PDUSessionInactivityTimer != nil {
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
	if err = w.WriteEnumerate(uint64(s.PDUSessionType.Value), aper.Constraint{Lb: 0, Ub: 4}, true); err != nil {
		return fmt.Errorf("Encode PDUSessionType failed: %w", err)
	}
	if err = s.SNSSAI.Encode(w); err != nil {
		return fmt.Errorf("Encode SNSSAI failed: %w", err)
	}
	if err = s.SecurityIndication.Encode(w); err != nil {
		return fmt.Errorf("Encode SecurityIndication failed: %w", err)
	}
	if s.PDUSessionResourceAMBR != nil {
		if err = w.WriteInteger(int64((*s.PDUSessionResourceAMBR)), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode PDUSessionResourceAMBR failed: %w", err)
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
		if err = w.WriteInteger(int64((*s.PDUSessionInactivityTimer)), &aper.Constraint{Lb: 1, Ub: 7200}, true); err != nil {
			return fmt.Errorf("Encode PDUSessionInactivityTimer failed: %w", err)
		}
	}
	if err = s.DRBToSetupModListNGRAN.Encode(w); err != nil {
		return fmt.Errorf("Encode DRBToSetupModListNGRAN failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceToSetupModItem) Decode(r *aper.AperReader) (err error) {
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
				return fmt.Errorf("Decode PDUSessionResourceAMBR failed: %w", err)
			}
			tmp := BitRate(val)
			s.PDUSessionResourceAMBR = &tmp
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
			tmp := InactivityTimer(val)
			s.PDUSessionInactivityTimer = &tmp
		}

	}
	if err = s.DRBToSetupModListNGRAN.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBToSetupModListNGRAN failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(PDUSessionResourceToSetupModItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for PDUSessionResourceToSetupModItem")
	}
	return nil
}
