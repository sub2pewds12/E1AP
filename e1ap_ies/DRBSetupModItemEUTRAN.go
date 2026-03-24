package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DRBSetupModItemEUTRAN is a generated SEQUENCE type.
type DRBSetupModItemEUTRAN struct {
	DRBID                             DRBID
	S1DLUPTNLInformation              UPTNLInformation
	DataForwardingInformationResponse *DataForwardingInformation
	ULUPTransportParameters           UPParameters
	IEExtensions                      *DRBSetupModItemEUTRANExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBSetupModItemEUTRAN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "s1-DL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "data-Forwarding-Information-Response", Optional: true},
			per.ComponentInfo{Name: "uL-UP-Transport-Parameters", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.DataForwardingInformationResponse != nil {
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

	if err = w.EncodeInteger(int64(s.DRBID.Value), per.ConstrainedExtensible(1, 32)); err != nil {
		return fmt.Errorf("encode DRBID failed: %w", err)
	}
	if err = s.S1DLUPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("encode S1DLUPTNLInformation failed: %w", err)
	}

	if s.DataForwardingInformationResponse != nil {
		if err = s.DataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("encode DataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.ULUPTransportParameters.Encode(w); err != nil {
		return fmt.Errorf("encode ULUPTransportParameters failed: %w", err)
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
func (s *DRBSetupModItemEUTRAN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "s1-DL-UP-TNL-Information", Optional: false},
			per.ComponentInfo{Name: "data-Forwarding-Information-Response", Optional: true},
			per.ComponentInfo{Name: "uL-UP-Transport-Parameters", Optional: false},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 32))
		if err != nil {
			return fmt.Errorf("decode DRBID failed: %w", err)
		}
		s.DRBID.Value = val
	}
	if err = s.S1DLUPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode S1DLUPTNLInformation failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.DataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.DataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("Decode DataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.ULUPTransportParameters.Decode(r); err != nil {
		return fmt.Errorf("Decode ULUPTransportParameters failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(DRBSetupModItemEUTRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
