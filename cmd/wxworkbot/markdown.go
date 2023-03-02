package cmd

import (
	"fmt"

	"github.com/bryant-rh/wxworkbot/pkg"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func init() {
	rootCmd.AddCommand(markdownCmd)

	markdownCmd.Flags().StringVarP(&content, "content", "c", "", "指定要发送的文本内容,最长不超过2048个字节[必须指定]")
	markdownCmd.MarkFlagRequired("content")

}

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "wxworkbot send markdown message",
	RunE: func(cmd *cobra.Command, args []string) error {
		//校验参数
		err := Validate(cmd, args)
		if err != nil {
			klog.Fatal(pkg.RedColor(err))
		}
		klog.V(4).Infoln("Send Markdown Message!")

		res, err := Client.SendMarkdown(webHookUrl, content)
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
