package push

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	offlinepush "github.com/txchat/dtalk/service/offline-push/api"
	"github.com/txchat/dtalk/service/offline-push/tools/mock"
	config "github.com/txchat/dtalk/service/offline-push/tools/mock/etc"
)

var SinglePushCmd = &cobra.Command{
	Use:     "single",
	Short:   "single push",
	Long:    "single push",
	Example: "single -n 1 --hm=true -d true",
	Run:     singlePush,
}

func init() {
	SinglePushCmd.Flags().StringVarP(&config.ConfPath, "conf", "c", "config.toml", "default config path.")
}

func singlePush(cmd *cobra.Command, args []string) {
	err := config.Init()
	if err != nil {
		fmt.Printf("config init error:%v\n", err)
		return
	}
	fmt.Printf("Client Config: %v\n Msg Config: %v\n", config.Conf.Client, config.Conf.Msg)
	c := mock.NewClient(config.Conf.Client.AppId, config.Conf.Client.Brokers)
	m := mock.Msg{
		AppId:       config.Conf.Msg.AppId,
		DeviceType:  offlinepush.Device(config.Conf.Msg.DeviceType),
		Nickname:    config.Conf.Msg.Nickname,
		TargetId:    config.Conf.Msg.TargetId,
		DeviceToken: config.Conf.Msg.DeviceToken,
	}
	data, err := m.Data()
	if err != nil {
		fmt.Printf("msg data error:%v\n", err)
		return
	}
	err = c.PublishOfflineMsg(context.Background(), config.Conf.Client.FromId, data)
	if err != nil {
		fmt.Printf("push msg error:%v\n", err)
		return
	}
	fmt.Println("success")
}
