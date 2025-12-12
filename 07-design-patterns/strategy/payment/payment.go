package payment

// interface
type PaymentProcessor interface {
	Run(value float32) float32
}

// equivalente a classes que implementam interface
type PixPayment struct{}

func (pix PixPayment) Run(value float32) float32 {
	return value
}

// equivalente a classes que implementam interface
type CardPayment struct{}

func (c CardPayment) Run(value float32) float32 {
	return value * 1.03
}

func RunPayment(value float32, p PaymentProcessor) float32 {
	return p.Run(value)
}
