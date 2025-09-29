package e1ap_ies

// TNLAssociationUsage represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2226
type TNLAssociationUsage int32

const (
	TNLAssociationUsage_Ue    TNLAssociationUsage = 0
	TNLAssociationUsage_NonUe TNLAssociationUsage = 1
	TNLAssociationUsage_Both  TNLAssociationUsage = 2
)
