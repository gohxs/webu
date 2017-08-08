package chain_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gohxs/webu/chain"
)

func chainTest(t *testing.T, name string) chain.Func {

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			t.Log(name)
			next(w, r)
			t.Log("After", name)
		}
	}
}

func TestOrder(t *testing.T) {
	lastHandler := func(w http.ResponseWriter, r *http.Request) {
		t.Log("Final")
	}

	mux := http.NewServeMux()

	mainChain := chain.New(chainTest(t, "one"), chainTest(t, "two"), chainTest(t, "three"))
	mux.HandleFunc("/", mainChain.Build(lastHandler))

	// chaining chains
	nchain := chain.New(chainTest(t, "Outer"), mainChain.Build)
	mux.HandleFunc("/sub", nchain.Build(lastHandler))

	testServer := httptest.NewServer(mux)

	http.Get(testServer.URL)
	t.Log("Other test")
	http.Get(testServer.URL + "/sub")
}
