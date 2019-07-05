package object

import(
	"encoding/json"
)

type Object struct {
	Field string `json:field`
}
func (o *Object) ToMarshal() []byte {
	rec, _ := json.Marshal(o)
	return rec
}
func (o *Object) FromMarshal(v []byte)(error){
	err := json.Unmarshal(v, &o)
	return err
}