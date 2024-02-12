package core

type GetModel struct {
	MsgID     string `form:"id"`
	MsgKey    string `form:"mk"`
	SystemKey string `form:"sk"`
	Raw       bool   `form:"r"`
}

type SaveModel struct {
	Msg       string `json:"m"`
	MsgKey    string `json:"mk"`
	EditKey   string `json:"ek"`
	SystemKey string `json:"sk"`
}

type EditModel struct {
	MsgID     string `json:"id"`
	Msg       string `json:"m"`
	EditKey   string `json:"ek"`
	SystemKey string `json:"sk"`
}

type Message struct {
	Msg     string `yaml:"msg"`
	MsgKey  string `yaml:"msgKey"`
	EditKey string `yaml:"editKey"`
}
