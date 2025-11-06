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
	ULUPTransportParameters           UPParameters                `aper:"mandatory,ext"`
	IEExtensions                      *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBSetupModItemEUTRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.DataForwardingInformationResponse != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if err = s.S1DLUPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("Encode S1DLUPTNLInformation failed: %w", err)
	}
	if s.DataForwardingInformationResponse != nil {
		if err = s.DataForwardingInformationResponse.Encode(w); err != nil {
			return fmt.Errorf("Encode DataForwardingInformationResponse failed: %w", err)
		}
	}

	{
		itemPointers := make([]aper.AperMarshaller, len(s.ULUPTransportParameters.Value))
		for i := 0; i < len(s.ULUPTransportParameters.Value); i++ {
			itemPointers[i] = &(s.ULUPTransportParameters.Value[i])
		}
		if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Encode ULUPTransportParameters failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBSetupModItemEUTRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
			return fmt.Errorf("Decode DRBID failed: %w", err)
		}
		s.DRBID.Value = aper.Integer(val)
	}
	if err = s.S1DLUPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode S1DLUPTNLInformation failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DataForwardingInformationResponse = new(DataForwardingInformation)
		if err = s.DataForwardingInformationResponse.Decode(r); err != nil {
			return fmt.Errorf("Decode DataForwardingInformationResponse failed: %w", err)
		}
	}
	if err = s.ULUPTransportParameters.Decode(r); err != nil {
		return fmt.Errorf("Decode ULUPTransportParameters failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DRBSetupModItemEUTRAN")
	}
	return nil
}
