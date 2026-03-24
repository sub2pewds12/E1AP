package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// EHCParameters is a generated SEQUENCE type.
type EHCParameters struct {
	EhcCommon    EHCCommonParameters
	EhcDownlink  *EHCDownlinkParameters
	EhcUplink    *EHCUplinkParameters
	IEExtensions *EHCParametersExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *EHCParameters) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "ehc-Common", Optional: false},
			per.ComponentInfo{Name: "ehc-Downlink", Optional: true},
			per.ComponentInfo{Name: "ehc-Uplink", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.EhcDownlink != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.EhcUplink != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
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
func (s *EHCParameters) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "ehc-Common", Optional: false},
			per.ComponentInfo{Name: "ehc-Downlink", Optional: true},
			per.ComponentInfo{Name: "ehc-Uplink", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	if err = s.EhcCommon.Decode(r); err != nil {
		return fmt.Errorf("Decode EhcCommon failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(1) {
		s.EhcDownlink = new(EHCDownlinkParameters)
		if err = s.EhcDownlink.Decode(r); err != nil {
			return fmt.Errorf("Decode EhcDownlink failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.EhcUplink = new(EHCUplinkParameters)
		if err = s.EhcUplink.Decode(r); err != nil {
			return fmt.Errorf("Decode EhcUplink failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(EHCParametersExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
