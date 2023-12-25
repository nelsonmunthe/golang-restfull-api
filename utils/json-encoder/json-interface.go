package json_encoder

type Json2 interface {
	Unmarshal(obj []byte, caster interface{}) error
	Marshal(obj interface{}) ([]byte, error)
}
