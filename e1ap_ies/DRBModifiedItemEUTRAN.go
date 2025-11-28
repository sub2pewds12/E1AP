package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBModifiedItemEUTRAN is a generated SEQUENCE type.
type DRBModifiedItemEUTRAN struct {
	DRBID                   DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation    *UPTNLInformation           `aper:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation    `aper:"optional,ext"`
	ULUPTransportParameters *UPParameters               `aper:"lb:1,ub:MaxnoofUPParameters,optional,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBModifiedItemEUTRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.S1DLUPTNLInformation != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.PDCPSNStatusInformation != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.ULUPTransportParameters != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if s.S1DLUPTNLInformation != nil {
		if err = s.S1DLUPTNLInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode S1DLUPTNLInformation failed: %w", err)
		}
	}
	if s.PDCPSNStatusInformation != nil {
		if err = s.PDCPSNStatusInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode PDCPSNStatusInformation failed: %w", err)
		}
	}
	if s.ULUPTransportParameters != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.ULUPTransportParameters.Value))
			for i := 0; i < len(s.ULUPTransportParameters.Value); i++ {
				itemPointers[i] = &(s.ULUPTransportParameters.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxnoofUPParameters}, false); err != nil {
				return fmt.Errorf("Encode ULUPTransportParameters failed: %w", err)
			}
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBModifiedItemEUTRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("Decode DRBID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.S1DLUPTNLInformation = new(UPTNLInformation)
		if err = s.S1DLUPTNLInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode S1DLUPTNLInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.PDCPSNStatusInformation = new(PDCPSNStatusInformation)
		if err = s.PDCPSNStatusInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode PDCPSNStatusInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.ULUPTransportParameters = new(UPParameters)
		if err = s.ULUPTransportParameters.Decode(r); err != nil {
			return fmt.Errorf("Decode ULUPTransportParameters failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBModifiedItemEUTRAN */
	}
	return nil
}
