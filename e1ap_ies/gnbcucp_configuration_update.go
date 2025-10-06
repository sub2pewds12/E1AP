package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPConfigurationUpdate is a generated SEQUENCE type.
type GNBCUCPConfigurationUpdate struct {
	TransactionID             aper.Integer               `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUCPName               *aper.OctetString          `aper:"optional,ignore,ext"`
	GNBCUCPTNLAToAddList      []GNBCUCPTNLAToAddItem     `aper:"optional,ignore,ext"`
	GNBCUCPTNLAToRemoveList   []GNBCUCPTNLAToRemoveItem  `aper:"optional,ignore,ext"`
	GNBCUCPTNLAToUpdateList   []GNBCUCPTNLAToUpdateItem  `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPConfigurationUpdate) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GNBCUCPConfigurationUpdate")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPConfigurationUpdate) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GNBCUCPConfigurationUpdate")
}
