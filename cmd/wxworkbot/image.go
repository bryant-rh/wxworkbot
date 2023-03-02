package cmd

import (
	"fmt"

	"github.com/bryant-rh/wxworkbot/pkg"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func init() {
	rootCmd.AddCommand(imageCmd)

	imageCmd.Flags().StringVarP(&base64, "base64", "b", "", "图片内容的base64编码, [必须指定]")
	imageCmd.Flags().StringVarP(&md5, "md5", "m", "", "图片内容(base64编码前)的md5值,[必须指定]")
	imageCmd.MarkFlagRequired("base64")
	imageCmd.MarkFlagRequired("md5")

}

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "wxworkbot send image message",
	RunE: func(cmd *cobra.Command, args []string) error {
		//校验参数
		err := Validate(cmd, args)
		if err != nil {
			klog.Fatal(pkg.RedColor(err))
		}
		klog.V(4).Infoln("Send Image Message!")

		res, err := Client.SendImage(webHookUrl, base64, md5)
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
