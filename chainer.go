/*
	Handler builder using chains of builders
*/

package webu

import "net/http"

// ChainFunc
type ChainFunc func(http.Handler) http.Handler

// NewChain create a chainer
func NewChain(chain ...ChainFunc) *ChainBuilder {
	c := &ChainBuilder{}
	c.Add(chain...)
	return c
}

//ChainBuilder chain struct
type ChainBuilder struct {
	// chain list
	chain []ChainFunc
}

// Add 1 or more chainFuncs to list
func (c *ChainBuilder) Add(chain ...ChainFunc) {
	c.chain = append(c.chain, chain...)
}

//Build retrieve handler after building
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
