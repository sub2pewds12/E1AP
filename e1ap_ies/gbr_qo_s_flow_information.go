package e1ap_ies

import (
	"fmt"
	"io"
)

// GBRQoSFlowInformation is a generated SEQUENCE type.
type GBRQoSFlowInformation struct {
	MaxFlowBitRateDownlink        *BitRate                    `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	MaxFlowBitRateUplink          *BitRate                    `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	GuaranteedFlowBitRateDownlink *BitRate                    `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	GuaranteedFlowBitRateUplink   *BitRate                    `aper:"lb:0,ub:4000000000000,optional,reject,ext"`
	MaxPacketLossRateDownlink     *MaxPacketLossRate          `aper:"lb:0,ub:1000,optional,ext"`
	MaxPacketLossRateUplink       *MaxPacketLossRate          `aper:"lb:0,ub:1000,optional,ext"`
	IEExtensions                  *ProtocolExtensionContainer `aper:"optional,ext"`
	// Possible extensions:
	// - AlternativeQoSParaSetList (ID: id-AlternativeQoSParaSetList)
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
