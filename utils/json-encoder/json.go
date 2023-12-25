//go:generate rm -fr mocks
//go:generate mockery --all

package json_encoder

import (
	"encoding/json"
)

type json2 struct {
}

func New() Json2 {
	return json2{}
}

func (j json2) Unmarshal(obj []byte, caster interface{}) error {
	return json.Unmarshal(obj, &caster)
}

func (j json2) Marshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}
