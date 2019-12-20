package entities

import "fmt"

type MainKategori struct {
	id           int64
	mainKategori string
}

func (mainKategori MainKategori) toString() string {
	return fmt.Sprintf("id: %d\nname: %s", mainKategori.id, mainKategori.mainKategori)
}
