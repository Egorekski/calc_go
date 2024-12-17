package application

const (
	ErrUnmarshalNumber = "json: cannot unmarshal bool into Go struct field Request.expression of type string"
	ErrUnmarshalBool   = "json: cannot unmarshal number into Go struct field Request.expression of type string"
	ErrUnmarshalObject = "json: cannot unmarshal object into Go struct field Request.expression of type string"
)
