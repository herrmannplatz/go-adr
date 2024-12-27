package adr

import (
	"fmt"
	"strings"
	"time"
)

type Adr struct {
	Template []byte
	Number   int
	Title    string
}

func (adr Adr) Data() []byte {
	now := time.Now().Local()
	d := strings.Replace(string(adr.Template), "[DATE]", now.Format("2006-01-02"), -1)
	title := strings.Title(strings.Replace(adr.Title, "-", " ", -1))
	c := strings.Replace(d, "[TITLE]", fmt.Sprintf("%d. %s", adr.Number, title), -1)
	return []byte(c)
}

func (adr Adr) Filename() string {
	return fmt.Sprintf("%04d-%s.md", adr.Number, adr.Title)
}
