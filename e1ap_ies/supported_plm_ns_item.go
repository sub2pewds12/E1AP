package e1ap_ies

// SupportedPLMNsItem represents the ASN.1 definition from 9_4_4_PDU_Definitions.txt:363
type SupportedPLMNsItem struct {
	PLMNIdentity             []byte                    `asn1:"mandatory,ext"`
	SliceSupportList         []SliceSupportItem        `asn1:"lb:1,ub:Item,optional,ext"`
	NRCGISupportList         []NRCGISupportItem        `asn1:"lb:1,ub:Item,optional,ext"`
	QOSParametersSupportList *QOSParametersSupportList `asn1:"optional,ext"`
}
