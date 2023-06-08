package dto

type ResponseMeta struct {
	Success      bool   `json:"success"`
	MessageTitle string `json:"messageTitle"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Token        string `json:"token"`
}

type Response struct {
	Success      bool   `json:"success"`
	MessageTitle string `json:"messageTitle"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
}
