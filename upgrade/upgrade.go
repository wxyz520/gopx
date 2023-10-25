package upgrade

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"gopx/rpc/execx"
)

// upgrade gets the latest goctl by
// go install gopx@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install gopx@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn,direct && go install gopx@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
