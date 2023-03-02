package cmd

import (
	"fmt"

	"github.com/bryant-rh/wxworkbot/pkg"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func init() {
	rootCmd.AddCommand(newsCmd)

	newsCmd.Flags().StringVarP(&title, "title", "t", "", "标题不超过128个字节,超过会自动截断,[必须指定]")
	newsCmd.Flags().StringVarP(&description, "description", "d", "", "描述,不超过512个字节,超过会自动截断,[必须指定]")
	newsCmd.Flags().StringVarP(&url, "url", "u", "", "点击后跳转的链接。[必须指定]")
	newsCmd.Flags().StringVarP(&picurl, "picurl", "p", "", "图文消息的图片链接,支持JPG、PNG格式,较好的效果为大图 1068*455,小图150*150。")
	newsCmd.MarkFlagRequired("title")
	newsCmd.MarkFlagRequired("description")
	newsCmd.MarkFlagRequired("url")

}

var newsCmd = &cobra.Command{
	Use:   "news",
	Short: "wxworkbot send news message",
	RunE: func(cmd *cobra.Command, args []string) error {
		//校验参数
		err := Validate(cmd, args)
		if err != nil {
			klog.Fatal(pkg.RedColor(err))
		}
		klog.V(4).Infoln("Send News Message!")

		res, err := Client.SendNews(webHookUrl, title, description, url, picurl)
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
