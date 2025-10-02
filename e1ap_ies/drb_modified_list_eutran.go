package e1ap_ies

// DRBModifiedListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:544
type DRBModifiedListEUTRAN []DRBModifiedItemEUTRAN

// DRBModifiedItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:546
type DRBModifiedItemEUTRAN struct {
	DRBID                   int64                    `asn1:"mandatory,ext"`
	S1DLUPTNLInformation    *UPTNLInformation        `asn1:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation `asn1:"optional,ext"`
	ULUPTransportParameters []UPParametersItem       `asn1:"optional,ext"`
}
