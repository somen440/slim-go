package exception

type DomainException struct{}

func (d *DomainException) Error() string {
	return "domain exception."
}
