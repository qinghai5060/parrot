package export

import (
	"encoding/json"

	"github.com/anthonynsimon/parrot/parrot-api/model"
)

type JSON struct{}

func (e *JSON) FileExtension() string {
	return "json"
}

func (e *JSON) ContentType() string {
	return "application/json; charset=UTF-8"
}

func (e *JSON) Export(locale *model.Locale) ([]byte, error) {
	return json.MarshalIndent(locale.Pairs, "", "    ")
}
