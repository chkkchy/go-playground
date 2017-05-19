package constant

type basetype struct {
	ID    uint8  `json:"id"`
	Value string `json:"value"`
}

func (b basetype) GetID() uint8 {
	return b.ID
}

func (b basetype) GetValue() string {
	return b.Value
}

type data struct {
	Type    string   `json:"type"`
	Options basetype `json:"options"`
}

type dataiface interface {
	GetID() uint8
	GetValue() string
}

func ToData(t string, di dataiface) *data {
	dt := new(data)
	dt.Type = t
	dt.Options = basetype{di.GetID(), di.GetValue()}
	return dt
}
