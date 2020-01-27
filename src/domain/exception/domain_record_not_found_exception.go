package exception

type DomainRecordNotFoundException struct{}

func (d *DomainRecordNotFoundException) Error() string {
	return "domain record not found exception."
}
