package operations

type Op interface {
}
type UnknownOp struct {
	// return exception message and - list help op
}
type HelpOp struct {
	// list help op
}
type ListOp struct {
	Search string
	Type   string
	Order  string
}
type CreateOp struct {
	Id string
}
type UpdateOp struct {
	Id string
}
type DeleteOp struct {
	Id string
}
