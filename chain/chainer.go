/*
	Handler builder using chains of builders
*/
package chain

import (
	"net/http"
	"reflect"
)

// Func chain Func type
type Func func(http.HandlerFunc) http.HandlerFunc

//New create a new chainer
// Params can be chain.Func, http.HandlerFunc or another chainer
func New(chainFunc ...interface{}) *Chain {
	c := &Chain{}
	c.Add(chainFunc...)
	return c
}

// NewChain create a chainer
/*func NewChain(chain ...Func) *Chain {
	c := &Chain{}
	c.Add(chain...)
	return c
}*/

//Chain struct
type Chain struct {
	// chain list
	chain []Func
}

// Add 1 or more chainFuncs to list
func (c *Chain) Add(chain ...interface{}) {
	// Convert whatever to chainFunc
	for _, v := range chain {
		switch vt := v.(type) {
		case func(handler http.HandlerFunc) http.HandlerFunc:
			c.chain = append(c.chain, vt)
		case Func:
			c.chain = append(c.chain, vt)
		default:
			panic("Unsupported now: " + reflect.TypeOf(v).String())
		}

	}

	//c.chain = append(c.chain, chain...)
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
