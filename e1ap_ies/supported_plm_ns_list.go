package e1ap_ies

// SupportedPLMNsList From: 9_4_4_PDU_Definitions.txt:361
type SupportedPLMNsList []SupportedPLMNsItem

// SupportedPLMNsItem From: 9_4_4_PDU_Definitions.txt:363
type SupportedPLMNsItem struct {
	PLMNIdentity             []byte                    `asn1:"lb:3,ub:3,mandatory,ext"`
	SliceSupportList         []SliceSupportItem        `asn1:"optional,ext"`
	NRCGISupportList         []NRCGISupportItem        `asn1:"optional,ext"`
	QOSParametersSupportList *QOSParametersSupportList `asn1:"optional,ext"`
}
