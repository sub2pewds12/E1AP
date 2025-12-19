package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EHCParameters is a generated SEQUENCE type.
type EHCParameters struct {
	EhcCommon    EHCCommonParameters         `aper:"mandatory"`
	EhcDownlink  *EHCDownlinkParameters      `aper:"optional"`
	EhcUplink    *EHCUplinkParameters        `aper:"optional"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *EHCParameters) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.EhcDownlink != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.EhcUplink != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = s.EhcCommon.Encode(w); err != nil {
		return fmt.Errorf("encode EhcCommon failed: %w", err)
	}
	if s.EhcDownlink != nil {
		if err = s.EhcDownlink.Encode(w); err != nil {
			return fmt.Errorf("encode EhcDownlink failed: %w", err)
		}
	}
	if s.EhcUplink != nil {
		if err = s.EhcUplink.Encode(w); err != nil {
			return fmt.Errorf("encode EhcUplink failed: %w", err)
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
func (s *EHCParameters) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.EhcCommon.Decode(r); err != nil {
		return fmt.Errorf("decode EhcCommon failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.EhcDownlink = new(EHCDownlinkParameters)
		if err = s.EhcDownlink.Decode(r); err != nil {
			return fmt.Errorf("decode EhcDownlink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.EhcUplink = new(EHCUplinkParameters)
		if err = s.EhcUplink.Decode(r); err != nil {
			return fmt.Errorf("decode EhcUplink failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for EHCParameters */
	}
	return nil
}
