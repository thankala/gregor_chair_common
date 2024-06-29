package configuration

var (
	defaultAddress = "localhost:8080"
)

type TcpOpts struct {
	Address string
}

func DefaultTcpOpts() *TcpOpts {
	return &TcpOpts{
		Address: defaultAddress,
	}
}

type TcpOptFunc func(*TcpOpts)

func WithAddress(address string) TcpOptFunc {
	return func(opts *TcpOpts) {
		opts.Address = address
	}
}
