package infra

type MPCServiceAdapter interface {
	Connect(host string, port int) error
	Close()
	PrimeGen(activeMembers int32) (string, error)
}
