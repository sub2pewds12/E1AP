package e1ap_ies

// DRBModifiedItemEUTRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:546
type DRBModifiedItemEUTRAN struct {
	DRBID                   int64                    `asn1:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation    *UPTNLInformation        `asn1:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation `asn1:"optional,ext"`
	ULUPTransportParameters []UPParametersItem       `asn1:"lb:1,ub:Item,optional,ext"`
}
