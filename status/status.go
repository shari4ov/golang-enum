package status

type status string

const (
	Success status = "success"
	Reject  status = "reject"
	Pending status = "pending"
)

type Status interface {
	getSuccessStatus() status
	getRejectStatus() status
	getPendingStatus() status
}

func (s status) getSuccessStatus() status {
	return Success
}

func (s status) getRejectStatus() status {
	return Reject
}
func (s status) getPendingStatus() status {
	return Pending
}
