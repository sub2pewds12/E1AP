package e1ap_ies

// HWCapacityIndicator is a generated SEQUENCE type.
type HWCapacityIndicator struct {
	OfferedThroughput   HWCapacityIndicatorOfferedThroughput   `aper:"lb:1,ub:16777216,mandatory,ext"`
	AvailableThroughput HWCapacityIndicatorAvailableThroughput `aper:"lb:0,ub:100,mandatory,ext"`
	IEExtensions        ProtocolExtensionContainer             `aper:"mandatory,ext"`
}
