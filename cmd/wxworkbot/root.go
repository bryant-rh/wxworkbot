package cmd

import (
	"fmt"
	"os"

	"github.com/bryant-rh/wxworkbot/pkg"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	enableDebug bool
	botKey      = os.Getenv("BOT_KEY")
	version     string

	content               string
	mentioned_list        string
	mentioned_mobile_list string
	base64                string
	md5                   string
	title                 string
	description           string
	url                   string
	picurl                string
	webHookUrl            string

	Client *pkg.CMClient
)

const (
	defaultWebHookUrlTemplate = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

// versionString returns the version prefixed by 'v'
// or an empty string if no version has been populated by goreleaser.
// In this case, the --version flag will not be added by cobra.
func versionString() string {
	if len(version) == 0 {
		return ""
	}
	return "v" + version
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "wxworkbot",
	Short:   "wxworkbot is a command-line tool for enterprise WeChat robots",
	Version: versionString(),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		err := Validate(cmd, args)
		if err != nil {
			klog.Fatal(pkg.RedColor(err))
		}
		webHookUrl = fmt.Sprintf(defaultWebHookUrlTemplate, botKey)
		Client = pkg.NewReqClient()
		if enableDebug { // Enable debug mode if `--enableDebug=true` or `DEBUG=true`.
			Client.SetDebug(true)
		}

	},
}

func Validate(cmd *cobra.Command, args []string) error {

	if botKey == "" {
		return fmt.Errorf("环境变量BOT_KEY为空:'%s',请设置", botKey)
	}
	return nil

}

func init() {
	rootCmd.PersistentFlags().BoolVar(&enableDebug, "debug", os.Getenv("DEBUG") == "true", "Enable debug mode")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
