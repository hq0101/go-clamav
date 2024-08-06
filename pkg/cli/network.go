package cli

type NetworkType string

const (
	Unix NetworkType = "unix"
	TCP  NetworkType = "tcp"
)

func (t NetworkType) String() string {
	return string(t)
}
