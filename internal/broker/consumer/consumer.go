package consumer

type Consumer interface {
	Customer()
	Order()
	Payment()
}
