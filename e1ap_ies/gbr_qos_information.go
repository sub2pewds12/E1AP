package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GBRQosInformation is a generated SEQUENCE type.
type GBRQosInformation struct {
	ERABMaximumBitrateDL    aper.Integer                `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
	ERABMaximumBitrateUL    aper.Integer                `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
	ERABGuaranteedBitrateDL aper.Integer                `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
	ERABGuaranteedBitrateUL aper.Integer                `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
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
