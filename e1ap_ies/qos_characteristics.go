package e1ap_ies

// QOSCharacteristics is a generated CHOICE type.
type QOSCharacteristics struct {
	Choice        uint64
	NonDynamic5QI *NonDynamic5QIDescriptor
	Dynamic5QI    *Dynamic5QIDescriptor
}

const (
	QOSCharacteristicsPresentNothing uint64 = iota
	QOSCharacteristicsPresentNonDynamic5QI
	QOSCharacteristicsPresentDynamic5QI
)
