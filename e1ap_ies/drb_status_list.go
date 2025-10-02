package e1ap_ies

// DRBStatusList From: 9_4_4_PDU_Definitions.txt:1069
type DRBStatusList []DRBStatusItem

// DRBStatusItem From: 9_4_5_Information_Element_Definitions.txt:684
type DRBStatusItem struct {
	DRBID       int64      `asn1:"mandatory,ext"`
	PDCPDLCount *PDCPCount `asn1:"optional,ext"`
	PDCPULCount *PDCPCount `asn1:"optional,ext"`
}
