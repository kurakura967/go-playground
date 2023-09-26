package main

import (
	"fmt"
	"time"
)

type Config struct {
	port int
}

type ConfigBuilder struct {
	port *int
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

const defaultPort = 0000
const randomPort = 0001

func (b *ConfigBuilder) Build() (Config, error) {

	cnf := Config{}

	if b.port == nil {
		cnf.port = defaultPort
	} else {
		if *b.port == 0 {
			cnf.port = randomPort
		} else if *b.port < 0 {
			return Config{}, fmt.Errorf("port must be positive")
		} else {
			cnf.port = *b.port
		}
	}

	return cnf, nil
}

type options struct {
	port    *int
	timeout time.Duration
}

type Option func(*options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return fmt.Errorf("port must be positive")
		}
		options.port = &port
		return nil
	}
}

func WithTimeout(sec time.Duration) Option {
	return func(options *options) error {
		options.timeout = sec
		return nil
	}
}

func NewServer(add string, opts ...Option) {
	var options options

	for _, opt := range opts {
		// 共通のoptions構造体を修正する
		err := opt(&options)
		if err != nil {
			panic(err)
		}
	}

	var port int
	if options.port == nil {
		port = defaultPort
	} else {
		if *options.port == 0 {
			port = randomPort
		} else {
			port = *options.port
		}
	}

	fmt.Println(port)
	fmt.Println(options.timeout)
}

func main() {
	// builderパターン
	builder := ConfigBuilder{}
	cnf, err := builder.Port(8080).Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(cnf)

	// functional optionパターン
	NewServer("localhost", WithPort(8080), WithTimeout(time.Second*10))

}
