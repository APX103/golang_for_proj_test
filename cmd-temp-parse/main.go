package main

import (
	"fmt"

	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("======= Cobra test =======")
	bar_2 := ""
	e_2 := 0
	bar_3 := ""
	e_3 := 0
	var params map[string]string
	rootCmd := &cobra.Command{
		Use:   "@Mr.meeseeks",
		Short: "Feishu Agent Build By QA Team.",
		Long:  `Feishu Agent Build By QA Team. Any question please try '--help' or ask @李佳伦 for help.`,
	}
	rootCmd.PersistentFlags().StringVar(&bar_2, "bar", "fuck", "总命令行的参数1")
	rootCmd.PersistentFlags().IntVar(&e_2, "e", 1001, "总命令行的参数2")
	var runCmd = &cobra.Command{
		Use:   "Jenkins",
		Short: "Jenkins Task Agent, can build jobs and check status of job.",
		Long:  `Jenkins Task Agent, can build jobs and check status of job.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	runCmd.PersistentFlags().StringVar(&bar_3, "bar", "fuck", "Jenkins 测试参数1")
	runCmd.PersistentFlags().IntVar(&e_3, "e", 1001, "Jenkins 测试参数2")
	runCmd.PersistentFlags().StringToStringVar(&params, "params", nil, "jenkins键值对 Parameters Test")
	rootCmd.AddCommand(runCmd)

	argv, err := shellwords.Parse("@Mr.meeseeks Jenkins --bar=baz --e=1 --params='a=v,b=qahdiuowqhduowqjdioqw'")
	if err != nil {
		fmt.Println("牛逼")
	}
	fmt.Println(argv)

	rootCmd.SetArgs(argv[1:])
	err = rootCmd.Execute()

	if err != nil {
		fmt.Println("牛逼")
		fmt.Print(err)
	}

	fmt.Println(bar_2)
	fmt.Println(e_2)
	fmt.Println(bar_3)
	fmt.Println(e_3)
	fmt.Println(params)

	rootCmd.SetArgs([]string{"--help"})
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println("牛逼")
		fmt.Print(err)
	}

	rootCmd.SetArgs([]string{"Jenkins", "--help"})
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println("牛逼")
		fmt.Print(err)
	}

	fmt.Println("==================")

	rootCmd.SetArgs(argv[1:])
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println("牛逼")
		fmt.Print(err)
	}
}
