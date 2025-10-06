package e1ap_ies

// PDCPConfiguration is a generated SEQUENCE type.
type PDCPConfiguration struct {
	PDCPSNSizeUL          PDCPSNSize             `aper:"mandatory,ext"`
	PDCPSNSizeDL          PDCPSNSize             `aper:"mandatory,ext"`
	RLCMode               RLCMode                `aper:"mandatory,ext"`
	ROHCParameters        *ROHCParameters        `aper:"optional,ext"`
	TReorderingTimer      *TReorderingTimer      `aper:"optional,ext"`
	DiscardTimer          *DiscardTimer          `aper:"optional,ext"`
	ULDataSplitThreshold  *ULDataSplitThreshold  `aper:"optional,ext"`
	PDCPDuplication       *PDCPDuplication       `aper:"optional,ext"`
	PDCPReestablishment   *PDCPReestablishment   `aper:"optional,ext"`
	PDCPDataRecovery      *PDCPDataRecovery      `aper:"optional,ext"`
	DuplicationActivation *DuplicationActivation `aper:"optional,ext"`
	OutOfOrderDelivery    *OutOfOrderDelivery    `aper:"optional,ext"`
}
