package infra

type TenancyRestAdapter interface {
	Connect()
	Close()
}
