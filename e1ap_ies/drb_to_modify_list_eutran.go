package e1ap_ies

// DRBToModifyListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:738
type DRBToModifyListEUTRAN []DRBToModifyItemEUTRAN

// DRBToModifyItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:740
type DRBToModifyItemEUTRAN struct {
	DRBID                     int64                      `asn1:"mandatory,ext"`
	PDCPConfiguration         *PDCPConfiguration         `asn1:"optional,ext"`
	EUTRANQOS                 *EUTRANQOS                 `asn1:"optional,ext"`
	S1ULUPTNLInformation      *UPTNLInformation          `asn1:"optional,ext"`
	DataForwardingInformation *DataForwardingInformation `asn1:"optional,ext"`
	PDCPSNStatusRequest       *PDCPSNStatusRequest       `asn1:"optional,ext"`
	PDCPSNStatusInformation   *PDCPSNStatusInformation   `asn1:"optional,ext"`
	DLUPParameters            []UPParametersItem         `asn1:"optional,ext"`
	CellGroupToAdd            []CellGroupInformationItem `asn1:"optional,ext"`
	CellGroupToModify         []CellGroupInformationItem `asn1:"optional,ext"`
	CellGroupToRemove         []CellGroupInformationItem `asn1:"optional,ext"`
	DRBInactivityTimer        *int64                     `asn1:"optional,ext"`
}
