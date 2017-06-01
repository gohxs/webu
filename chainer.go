/*
	Handler builder using chains of builders
*/

package webu

import "net/http"

// ChainFunc
type ChainFunc func(http.HandlerFunc) http.HandlerFunc

// NewChain create a chainer
func NewChain(chain ...ChainFunc) *Chain {
	c := &Chain{}
	c.Add(chain...)
	return c
}

//ChainBuilder chain struct
type Chain struct {
	// chain list
	chain []ChainFunc
}

// Add 1 or more chainFuncs to list
func (c *Chain) Add(chain ...ChainFunc) {
	c.chain = append(c.chain, chain...)
}

//Build retrieve handler after building
func (c *Chain) Build(handler http.HandlerFunc) http.HandlerFunc {
	if len(c.chain) == 0 { // Pass trough
		return handler
	}
	finalHandler := c.chain[len(c.chain)-1](handler) // last
	for i := len(c.chain) - 2; i >= 0; i-- {
		v := c.chain[i]
		finalHandler = v(finalHandler)
	}
	return finalHandler
}

/*func (c *Chain) BuildFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return c.Build(handlerFunc).ServeHTTP
}*/
