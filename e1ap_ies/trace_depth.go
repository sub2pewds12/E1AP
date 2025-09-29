package e1ap_ies

// TraceDepth represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2285
type TraceDepth int32

const (
	TraceDepth_Minimum                               TraceDepth = 0
	TraceDepth_Medium                                TraceDepth = 1
	TraceDepth_Maximum                               TraceDepth = 2
	TraceDepth_MinimumWithoutVendorSpecificExtension TraceDepth = 3
	TraceDepth_MediumWithoutVendorSpecificExtension  TraceDepth = 4
	TraceDepth_MaximumWithoutVendorSpecificExtension TraceDepth = 5
)
