package e1ap_ies

// TSCTrafficCharacteristics is a generated SEQUENCE type.
type TSCTrafficCharacteristics struct {
	TSCTrafficCharacteristicsUL *TSCTrafficInformation `aper:"optional"`
	TSCTrafficCharacteristicsDL *TSCTrafficInformation `aper:"optional"`
}
