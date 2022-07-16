package options

/*
选项模式
*/
import "time"

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultCache   = false
	defaultTimeout = 10
)

type options struct {
	cache   bool
	timeout time.Duration
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.cache = cache
	})
}

func NewConnect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		cache:   defaultCache,
	}
	for _, o := range opts {
		o.apply(&options)
	}
	return &Connection{
		addr:    addr,
		cache:   options.cache,
		timeout: options.timeout,
	}, nil
}
