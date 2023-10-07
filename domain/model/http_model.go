package model

type OK struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Err struct {
	Code    int           `json:"code"`
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Detail  []interface{} `json:"detail,omitempty"`
}
