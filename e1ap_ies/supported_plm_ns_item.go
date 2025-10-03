package e1ap_ies

import "github.com/lvdund/ngap/aper"

// SupportedPLMNsItem From: 9_4_4_PDU_Definitions.txt:363
// ASN.1 Data Type: SEQUENCE
type SupportedPLMNsItem struct {
	PLMNIdentity             aper.OctetString          `aper:"lb:3,ub:3,mandatory,ignore,ext"`
	SliceSupportList         []SliceSupportItem        `aper:"optional,ext"`
	NRCGISupportList         []NRCGISupportItem        `aper:"optional,ext"`
	QOSParametersSupportList *QOSParametersSupportList `aper:"optional,ext"`
}
