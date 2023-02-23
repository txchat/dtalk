package model

type Device struct {
	Uid         string `json:"uid"`
	ConnectId   string `json:"connectId"`
	DeviceUuid  string `json:"deviceUuid"`
	DeviceType  int32  `json:"deviceType"`
	DeviceName  string `json:"deviceName"`
	Username    string `json:"username"`
	DeviceToken string `json:"deviceToken"`
	IsEnabled   bool   `json:"isEnabled"`
	AddTime     int64  `json:"addTime"`
	DTUid       string `json:"-"`
}
