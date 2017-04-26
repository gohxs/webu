package webu

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

//Controller interface
type Controller interface {
}

// Aliast for handleFunc
type Action http.HandlerFunc

func (m *Manager) AddController(name string, ctrl Controller) {

	ctrlVal := reflect.Indirect(reflect.ValueOf(ctrl))
	ctrlTyp := ctrlVal.Type()

	log.Println("Analysing controller for", name)
	log.Println("Fields:", ctrlTyp.NumField())
	for i := 0; i < ctrlTyp.NumField(); i++ {
		fieldTyp := ctrlTyp.Field(i)
		fieldVal := ctrlVal.Field(i)

		log.Printf("Inspecting '%s'", fieldTyp.Name)
		if fieldVal.Kind() != reflect.Func {
			log.Println("..Fail not a func")
			continue
		}

		actionFunc, ok := fieldVal.Interface().(Action)
		if !ok {
			log.Println("..Fail not an Action")
			continue
		}
		log.Println("..OK")

		// Check func reflect for type

		actionMethods := []string{}
		actionName := fieldTyp.Name

		log.Println("Analysing tags")
		tag, ok := fieldTyp.Tag.Lookup("webu")
		if ok {
			log.Println("Parsing tags", tag)
			tagPart := strings.Split(tag, ";")

			actionName = tagPart[0]
			actionMethods = strings.Split(tagPart[1], ",") // In case we have methods
		}

		if len(actionMethods) == 0 {
			entryStr := fmt.Sprintf("#%s%s/%s", m.mainName, name, actionName)
			m.entry[entryStr] = actionFunc
			log.Printf("Register route: %s", entryStr)
			return
		}
		for _, v := range actionMethods {
			entryStr := fmt.Sprintf("%s#%s%s/%s", v, m.mainName, name, actionName)
			m.entry[entryStr] = actionFunc
			log.Printf("Register route: %s", entryStr)
		}

	}

}
