package router

type ResponseBody struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}
