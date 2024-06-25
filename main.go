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
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	rootCmd.PersistentFlags().StringVar(&bar_2, "bar", "fuck", "Bar Test")
	rootCmd.PersistentFlags().IntVar(&e_2, "e", 1001, "E Test")
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run this code",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	runCmd.PersistentFlags().StringVar(&bar_3, "bar", "fuck", "Bar Test")
	runCmd.PersistentFlags().IntVar(&e_3, "e", 1001, "E Test")
	runCmd.PersistentFlags().StringToStringVar(&params, "params", nil, "Parameters Test")
	rootCmd.AddCommand(runCmd)

	argv, err := shellwords.Parse("@Mr.meeseeks run --bar=baz --e=1 --params='a=v,b=qahdiuowqhduowqjdioqw'")
	if err != nil {
		fmt.Println("牛逼")
	}
	fmt.Println(argv)
	// err = rootCmd.ParseFlags(argv[1:])

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
}
