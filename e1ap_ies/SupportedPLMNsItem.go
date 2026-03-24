package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// SupportedPLMNsItem is a generated SEQUENCE type.
type SupportedPLMNsItem struct {
	PLMNIdentity             PLMNIdentity
	SliceSupportList         *SliceSupportList
	NRCGISupportList         *NRCGISupportList
	QOSParametersSupportList *QOSParametersSupportList
	IEExtensions             *SupportedPLMNsItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *SupportedPLMNsItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pLMN-Identity", Optional: false},
			per.ComponentInfo{Name: "slice-Support-List", Optional: true},
			per.ComponentInfo{Name: "nR-CGI-Support-List", Optional: true},
			per.ComponentInfo{Name: "qoS-Parameters-Support-List", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.SliceSupportList != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.NRCGISupportList != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.QOSParametersSupportList != nil {
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

	if err = w.EncodeOctetString([]byte(s.PLMNIdentity.Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)}); err != nil {
		return fmt.Errorf("encode PLMNIdentity failed: %w", err)
	}

	if s.SliceSupportList != nil {
		if err = s.SliceSupportList.Encode(w); err != nil {
			return fmt.Errorf("encode SliceSupportList failed: %w", err)
		}
	}

	if s.NRCGISupportList != nil {
		if err = s.NRCGISupportList.Encode(w); err != nil {
			return fmt.Errorf("encode NRCGISupportList failed: %w", err)
		}
	}

	if s.QOSParametersSupportList != nil {
		if err = s.QOSParametersSupportList.Encode(w); err != nil {
			return fmt.Errorf("encode QOSParametersSupportList failed: %w", err)
		}
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SupportedPLMNsItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pLMN-Identity", Optional: false},
			per.ComponentInfo{Name: "slice-Support-List", Optional: true},
			per.ComponentInfo{Name: "nR-CGI-Support-List", Optional: true},
			per.ComponentInfo{Name: "qoS-Parameters-Support-List", Optional: true},
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

	{
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)})
		if err != nil {
			return fmt.Errorf("decode PLMNIdentity failed: %w", err)
		}
		s.PLMNIdentity.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.SliceSupportList = new(SliceSupportList)
		if err = s.SliceSupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode SliceSupportList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.NRCGISupportList = new(NRCGISupportList)
		if err = s.NRCGISupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode NRCGISupportList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.QOSParametersSupportList = new(QOSParametersSupportList)
		if err = s.QOSParametersSupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode QOSParametersSupportList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(SupportedPLMNsItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
