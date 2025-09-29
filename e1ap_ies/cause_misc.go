package e1ap_ies

// CauseMisc represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:153
type CauseMisc int32

const (
	CauseMisc_ControlProcessingOverload             CauseMisc = 0
	CauseMisc_NotEnoughUserPlaneProcessingResources CauseMisc = 1
	CauseMisc_HardwareFailure                       CauseMisc = 2
	CauseMisc_OmIntervention                        CauseMisc = 3
	CauseMisc_Unspecified                           CauseMisc = 4
)
