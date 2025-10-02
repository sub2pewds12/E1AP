package e1ap_ies

// SNSSAI From: 9_4_5_Information_Element_Definitions.txt:2185
type SNSSAI struct {
	SST []byte  `asn1:"mandatory,ext"`
	SD  *[]byte `asn1:"optional,ext"`
}
