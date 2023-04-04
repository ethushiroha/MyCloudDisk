package ApiModels

type MyResponse struct {
	StatusCode int
	Body       RespBody
}

type RespBody interface {
	SetStatus(status string) RespBody
	SetMessage(message string) RespBody
	SetData(data map[string]interface{}) RespBody
	Build() RespBody
}

type BasicRespBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (b *BasicRespBody) SetStatus(status string) RespBody {
	b.Status = status
	return b
}

func (b *BasicRespBody) SetMessage(message string) RespBody {
	b.Message = message
	return b
}

func (b *BasicRespBody) SetData(data map[string]interface{}) RespBody {
	return b
}

func (b *BasicRespBody) Build() RespBody {
	return b
}
