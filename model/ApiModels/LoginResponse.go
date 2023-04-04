package ApiModels

// loginData 是 /Login 接口响应 Data 字段的结构体
type loginData struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginRespBody struct {
	BasicRespBody
	loginData `json:"data"`
}

func (login *LoginRespBody) SetData(data map[string]interface{}) RespBody {
	login.Username = data["username"].(string)
	login.Token = data["token"].(string)
	return login
}
