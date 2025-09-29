package e1ap_ies

// PDCPConfiguration represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1539
type PDCPConfiguration struct {
	PDCPSNSizeUL          PDCPSNSize             `asn1:"mandatory,ext"`
	PDCPSNSizeDL          PDCPSNSize             `asn1:"mandatory,ext"`
	RLCMode               RLCMode                `asn1:"mandatory,ext"`
	ROHCParameters        *ROHCParameters        `asn1:"optional,ext"`
	TReorderingTimer      *TReorderingTimer      `asn1:"optional,ext"`
	DiscardTimer          *DiscardTimer          `asn1:"optional,ext"`
	ULDataSplitThreshold  *ULDataSplitThreshold  `asn1:"optional,ext"`
	PDCPDuplication       *PDCPDuplication       `asn1:"optional,ext"`
	PDCPReestablishment   *PDCPReestablishment   `asn1:"optional,ext"`
	PDCPDataRecovery      *PDCPDataRecovery      `asn1:"optional,ext"`
	DuplicationActivation *DuplicationActivation `asn1:"optional,ext"`
	OutOfOrderDelivery    *OutOfOrderDelivery    `asn1:"optional,ext"`
}
