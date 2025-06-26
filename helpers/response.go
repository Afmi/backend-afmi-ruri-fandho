package helpers

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Info    InfoPage    `json:"info"`
	Data    interface{} `json:"data,omitempty"`
}

type InfoPage struct {
	Page       int `json:"page,omitempty"`
	Length     int `json:"length,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
	TotalData  int `json:"totalData,omitempty"`
}
