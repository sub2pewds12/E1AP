package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DRBModifiedItemNGRAN is a generated SEQUENCE type.
type DRBModifiedItemNGRAN struct {
	DRBID                   DRBID
	ULUPTransportParameters *UPParameters
	PDCPSNStatusInformation *PDCPSNStatusInformation
	FlowSetupList           *QOSFlowList
	FlowFailedList          *QOSFlowFailedList
	IEExtensions            *DRBModifiedItemNGRANExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBModifiedItemNGRAN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "uL-UP-Transport-Parameters", Optional: true},
			per.ComponentInfo{Name: "pDCP-SN-Status-Information", Optional: true},
			per.ComponentInfo{Name: "flow-Setup-List", Optional: true},
			per.ComponentInfo{Name: "flow-Failed-List", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ULUPTransportParameters != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PDCPSNStatusInformation != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.FlowSetupList != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.FlowFailedList != nil {
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

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBModifiedItemNGRAN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "uL-UP-Transport-Parameters", Optional: true},
			per.ComponentInfo{Name: "pDCP-SN-Status-Information", Optional: true},
			per.ComponentInfo{Name: "flow-Setup-List", Optional: true},
			per.ComponentInfo{Name: "flow-Failed-List", Optional: true},
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

	if seqDecoder.IsComponentPresent(1) {
		s.ULUPTransportParameters = new(UPParameters)
		if err = s.ULUPTransportParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode ULUPTransportParameters failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.PDCPSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PDCPSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode PDCPSNStatusInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.FlowSetupList = new(QOSFlowList)
		if err = s.FlowSetupList.Decode(r); err != nil {
			return fmt.Errorf("Decode FlowSetupList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.FlowFailedList = new(QOSFlowFailedList)
		if err = s.FlowFailedList.Decode(r); err != nil {
			return fmt.Errorf("Decode FlowFailedList failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.IEExtensions = new(DRBModifiedItemNGRANExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
