package webu

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Manager - controller Manager
type Manager struct {
	mainName   string
	slashCount int
	// Routes
	entry map[string]Action // Basic entry
}

// New Manager
func CreateManager(name string) *Manager {
	if name[0] != '/' {
		name = "/" + name
	}
	if name[len(name)-1] != '/' {
		name = name + "/"
	}
	slashCount := strings.Count(name, "/")
	log.Println("There are", slashCount, " slashes")

	m := &Manager{name, slashCount, map[string]Action{}}

	return m
}

/*
 count will check from the path name
*/
func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Do magic
	log.Println("Receiving request for path:", r.URL.Path)
	// Find by indexing '/'
	count := 0
	pathI := strings.IndexFunc(r.URL.Path, func(c rune) bool {
		if c == '/' {
			count++
		}
		if count > (m.slashCount + 1) { //  can be optimized
			return true
		}
		return false
	})

	if pathI == -1 {
		pathI = len(r.URL.Path)
	}
	rPath := r.URL.Path[:pathI]
	// Index adder
	if count == 2 {
		rPath += "/index"
	}
	if count == 3 && rPath[pathI-1] == '/' {
		rPath += "index" // Already has dash
	}

	log.Println("Slash count:", count)
	log.Println("Relevant path", rPath)
	log.Printf("Search for  entry %s#%s", r.Method, rPath)
	entryStr := fmt.Sprintf("%s#%s", r.Method, rPath)

	handler, ok := m.entry[entryStr]
	if ok {
		handler(w, r)
		return
	}
	//XXX: Search without method here:

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found\r\n"))
}
func (m *Manager) Handler() (string, http.Handler) {

	// Return handler for muxer with proper path
	return m.mainName, m
}
