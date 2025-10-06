package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GBRQoSFlowInformation is a generated SEQUENCE type.
type GBRQoSFlowInformation struct {
	MaxFlowBitRateDownlink        aper.Integer  `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
	MaxFlowBitRateUplink          aper.Integer  `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
	GuaranteedFlowBitRateDownlink aper.Integer  `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
	GuaranteedFlowBitRateUplink   aper.Integer  `aper:"lb:0,ub:4000000000000,mandatory,reject,ext"`
	MaxPacketLossRateDownlink     *aper.Integer `aper:"lb:0,ub:1000,optional,ext"`
	MaxPacketLossRateUplink       *aper.Integer `aper:"lb:0,ub:1000,optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GBRQoSFlowInformation) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GBRQoSFlowInformation")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GBRQoSFlowInformation) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GBRQoSFlowInformation")
}
