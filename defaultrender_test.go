package render

import (
	"fmt"
	"testing"
	"time"
)

func TestDefaultRender(t *testing.T) {
	type Contacto struct {
		Nombre          string
		Telefono        int
		FechaNacimiento time.Time
	}
	contacto := Contacto{}
	texto := string(DefaultRender(contacto))
	fmt.Println(texto)
}
