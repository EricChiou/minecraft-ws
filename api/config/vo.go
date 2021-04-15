package config

// response vo

// ListResVo user list data vo
type GetConfigResVo struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
