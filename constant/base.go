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
	Type    string     `json:"type"`
	Options []basetype `json:"options"`
}

type Dataiface interface {
	GetID() uint8
	GetValue() string
}

func ToData(t string, di []Dataiface) *data {
	dt := new(data)
	dt.Type = t
	for _, v := range di {
		dt.Options = append(dt.Options, basetype{v.GetID(), v.GetValue()})
	}
	return dt
}
