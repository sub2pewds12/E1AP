package e1ap_ies

// PrivateMessage is a generated SEQUENCE type.
type PrivateMessage struct {
	PrivateIEs []PrivateIEField `aper:"ub:MaxPrivateIEs,mandatory,ext"`
}
