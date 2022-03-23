package model

type GetModuleResp struct {
	Name      string   `json:"name" enums:"wallet,oa,redpacket"`
	IsEnabled bool     `json:"isEnabled"`
	EndPoints []string `json:"endPoints"`
}
