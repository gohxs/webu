package chain_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"dev.hexasoftware.com/hxs/webu/chain"
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

func lastHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Final")
}

func TestOrder(t *testing.T) {

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
