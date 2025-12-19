package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBSetupModItemEUTRAN is a generated SEQUENCE type.
type DRBSetupModItemEUTRAN struct {
	DRBID                             DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation              UPTNLInformation            `aper:"mandatory,ext"`
	DataForwardingInformationResponse *DataForwardingInformation  `aper:"optional,ext"`
	ULUPTransportParameters           UPParameters                `aper:"lb:1,ub:MaxnoofUPParameters,mandatory,ext"`
	IEExtensions                      *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBSetupModItemEUTRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.DataForwardingInformationResponse != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
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
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBSetupModItemEUTRAN) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("decode DRBID failed: %w", err)
	}
	if err = s.S1DLUPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("decode S1DLUPTNLInformation failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.DataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("decode DataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.ULUPTransportParameters.Decode(r); err != nil {
		return fmt.Errorf("decode ULUPTransportParameters failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBSetupModItemEUTRAN */
	}
	return nil
}
