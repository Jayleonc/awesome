package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"os/signal"
)

var (
	password string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "A simple CLI app",
		Run: func(cmd *cobra.Command, args []string) {
			// 在这里可以使用变量 password 进行后续操作
			fmt.Println()
			fmt.Println("Password:", password)
		},
	}

	// 添加 password 参数
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password")

	// 在执行之前，要求用户输入密码
	rootCmd.PreRun = func(cmd *cobra.Command, args []string) {
		if password == "" {

			fmt.Println("Enter password: ")

			statOut, _ := os.Stdout.Stat()
			statErr, _ := os.Stderr.Stat()

			if (statOut.Mode()&os.ModeCharDevice) == 0 && (statErr.Mode()&os.ModeCharDevice) == 0 {
				reader := bufio.NewReader(os.Stdin)
				read, _ := reader.ReadString('\n')
				password = read
				fmt.Println("程序是通过 nohup 方式启动的")
			} else {
				fd := int(os.Stdin.Fd())
				bytePassword, err := term.ReadPassword(fd)
				if err != nil {
					return
				}
				password = string(bytePassword)
				fmt.Println("程序不是通过 nohup 方式启动的")
			}

			//reader := bufio.NewReader(cmd.InOrStdin())
			//read, err := reader.ReadString('\n')
			//if err != nil {
			//	return
			//}

		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
