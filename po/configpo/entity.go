package configpo

// table name
const Table string = "config"

// pk name
const PK string = "key"

// column name
const (
	Key         string = "config.key"
	Value       string = "config.value"
	Description string = "config.description"
	CreatedDate string = "config.created_date"
	UpdatedDate string = "config.updated_date"
)

// Entity codemap table entity
type Entity struct {
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedDate string `json:"creatDate,omitempty"`
	UpdatedDate string `json:"updateDate,omitempty"`
}
