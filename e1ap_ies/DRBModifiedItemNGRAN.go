package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBModifiedItemNGRAN is a generated SEQUENCE type.
type DRBModifiedItemNGRAN struct {
	DRBID                   DRBID                           `aper:"lb:1,ub:32,mandatory,ext"`
	ULUPTransportParameters *UPParameters                   `aper:"lb:1,ub:MaxnoofUPParameters,optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation        `aper:"optional,ext"`
	FlowSetupList           *QOSFlowList                    `aper:"lb:1,ub:MaxnoofQoSFlows,optional,ext"`
	FlowFailedList          *QOSFlowFailedList              `aper:"lb:1,ub:MaxnoofQoSFlows,optional,ext"`
	IEExtensions            *DRBModifiedItemNGRANExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBModifiedItemNGRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ULUPTransportParameters != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.PDCPSNStatusInformation != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.FlowSetupList != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.FlowFailedList != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(5), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("encode DRBID failed: %w", err)
	}
	if s.ULUPTransportParameters != nil {
		if err = s.ULUPTransportParameters.Encode(w); err != nil {
			return fmt.Errorf("encode ULUPTransportParameters failed: %w", err)
		}
	}
	if s.PDCPSNStatusInformation != nil {
		if err = s.PDCPSNStatusInformation.Encode(w); err != nil {
			return fmt.Errorf("encode PDCPSNStatusInformation failed: %w", err)
		}
	}
	if s.FlowSetupList != nil {
		if err = s.FlowSetupList.Encode(w); err != nil {
			return fmt.Errorf("encode FlowSetupList failed: %w", err)
		}
	}
	if s.FlowFailedList != nil {
		if err = s.FlowFailedList.Encode(w); err != nil {
			return fmt.Errorf("encode FlowFailedList failed: %w", err)
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
func (s *DRBModifiedItemNGRAN) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("decode DRBID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ULUPTransportParameters = new(UPParameters)
		if err = s.ULUPTransportParameters.Decode(r); err != nil {
			return fmt.Errorf("decode ULUPTransportParameters failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.PDCPSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PDCPSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("decode PDCPSNStatusInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.FlowSetupList = new(QOSFlowList)
		if err = s.FlowSetupList.Decode(r); err != nil {
			return fmt.Errorf("decode FlowSetupList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.FlowFailedList = new(QOSFlowFailedList)
		if err = s.FlowFailedList.Decode(r); err != nil {
			return fmt.Errorf("decode FlowFailedList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(DRBModifiedItemNGRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBModifiedItemNGRAN */
	}
	return nil
}
