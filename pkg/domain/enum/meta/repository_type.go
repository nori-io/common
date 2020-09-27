package meta

type RepositoryType uint8

const (
	Git RepositoryType = iota
	Subversion
)
