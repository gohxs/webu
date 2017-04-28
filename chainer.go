/*
	Handler builder using chains of builders
*/

package webu

import "net/http"

type ChainFunc func(http.Handler) http.Handler

func NewChain(chain ...ChainFunc) *ChainBuilder {
	c := &ChainBuilder{}
	c.Add(chain...)
	return c
}

type ChainBuilder struct {
	// chain list
	chain []ChainFunc
}

func (c *ChainBuilder) Add(chain ...ChainFunc) {
	c.chain = append(c.chain, chain...)
}

func (c *ChainBuilder) Build(handler http.Handler) http.Handler {
	if len(c.chain) == 0 { // Pass trough
		return handler
	}
	finalHandler := c.chain[0](handler)
	for _, v := range c.chain[1:] {
		finalHandler = v(finalHandler)
	}
	return finalHandler
}
