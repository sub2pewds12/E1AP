package e1ap_ies

import (
	"fmt"
	"io"
)

// GBRQosInformation is a generated SEQUENCE type.
type GBRQosInformation struct {
	ERABMaximumBitrateDL    *BitRate                    `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	ERABMaximumBitrateUL    *BitRate                    `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	ERABGuaranteedBitrateDL *BitRate                    `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	ERABGuaranteedBitrateUL *BitRate                    `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GBRQosInformation) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GBRQosInformation")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GBRQosInformation) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GBRQosInformation")
}
