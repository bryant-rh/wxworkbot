package cmd

import (
	"fmt"
	"strings"

	"github.com/bryant-rh/wxworkbot/pkg"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func init() {
	rootCmd.AddCommand(textCmd)

	textCmd.Flags().StringVarP(&content, "content", "c", "", "指定要发送的文本内容,最长不超过2048个字节,[必须指定]")
	textCmd.Flags().StringVarP(&mentioned_list, "mentioned_list", "u", "", "userid的列表,提醒群中的指定成员(@某个成员),@all表示提醒所有人,如果开发者获取不到userid,可以使用mentioned_mobile_list")
	textCmd.Flags().StringVarP(&mentioned_mobile_list, "mentioned_mobile_list", "m", "", "手机号列表，提醒手机号对应的群成员(@某个成员),@all表示提醒所有人")
	textCmd.MarkFlagRequired("content")

}

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "wxworkbot send text message",
	RunE: func(cmd *cobra.Command, args []string) error {
		//校验参数
		err := Validate(cmd, args)
		if err != nil {
			klog.Fatal(pkg.RedColor(err))
		}
		klog.V(4).Infoln("Send Text Message!")

		mentioned_list_temp := strings.Split(mentioned_list, ",")
		mentioned_mobile_list_temp := strings.Split(mentioned_mobile_list, ",")
		res, err := Client.SendText(webHookUrl, content, mentioned_list_temp, mentioned_mobile_list_temp)
		if err != nil {
			klog.Fatal(err)
			return err
		}
		if res.ErrorCode != 0 && res.ErrorMessage != "" {
			return fmt.Errorf("发送失败,err: %s", res.ErrorMessage)
		}

		return nil

	},
}
