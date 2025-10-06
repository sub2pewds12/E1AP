package e1ap_ies

// MDTConfiguration is a generated SEQUENCE type.
type MDTConfiguration struct {
	MdtActivation MDTActivation `aper:"mandatory,ext"`
	MDTMode       MDTMode       `aper:"mandatory,ext"`
}
