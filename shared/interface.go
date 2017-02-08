package shared

type Git interface {
	GetVersion(_, reply *Version) error
}
