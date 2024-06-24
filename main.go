package main

import (
	"flag"
	"fmt"

	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func main() {
	envs, args, err := shellwords.ParseWithEnvs("FOO=foo BAR=baz ./foo --bar=baz --e=1")
	if err != nil {
		fmt.Printf("Got Problem with %s", err)
	}

	fmt.Println("Fine, this is a HelloWorld String")
	fmt.Println("Args:")
	for _, content := range args {
		fmt.Println(content)
	}
	fmt.Println("Envs:")
	for _, env := range envs {
		fmt.Println(env)
	}
	fmt.Println(args[1:])

	// var fs flag.FlagSet
	fs := flag.NewFlagSet("you", flag.ContinueOnError)
	bar := ""
	e := 0
	fs.StringVar(&bar, "bar", "fuck", "Bar Test")
	fs.IntVar(&e, "e", 1001, "E Test")

	err = fs.Parse(args[1:])

	if err != nil {
		fmt.Printf("Parse command fail cause: %s", err)
	}
	fmt.Println(fs.Parsed())
	fmt.Println()
	fmt.Println(bar)
	fmt.Println(e)

	fmt.Println()
	fs.PrintDefaults()

	fmt.Println("======================================================")

	fmt.Println("Cobra test")
	bar_2 := ""
	e_2 := 0
	bar_3 := ""
	e_3 := 0
	rootCmd := &cobra.Command{
		Use:   "foo",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	rootCmd.PersistentFlags().StringVar(&bar_2, "bar", "fuck", "Bar Test")
	rootCmd.PersistentFlags().IntVar(&e_2, "e", 1001, "E Test")

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	}

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run this code",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Run Code.")
		},
	}
	runCmd.PersistentFlags().StringVar(&bar_3, "bar", "fuck", "Bar Test")
	runCmd.PersistentFlags().IntVar(&e_3, "e", 1001, "E Test")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runCmd)

	argv, err := shellwords.Parse("./foo run --bar=baz --e=1")
	if err != nil {
		fmt.Println("牛逼")
	}
	fmt.Println(argv)
	err = rootCmd.ParseFlags(argv[1:])
	if err != nil {
		fmt.Println("牛逼")
		fmt.Print(err)
	}

	fmt.Println(bar_2)
	fmt.Println(e_2)
	fmt.Println("==============================================")

	argv_2, err := shellwords.Parse("./foo run --help")
	fmt.Println(argv_2)
	if err != nil {
		fmt.Println("牛逼")
		fmt.Print(err)
	}
	argv_3, err := shellwords.Parse("./foo run --bar=bac --e=2")
	fmt.Println(argv_2)
	if err != nil {
		fmt.Println("牛逼")
		fmt.Print(err)
	}
	err = rootCmd.ParseFlags(argv_2[1:])
	if err != nil {
		fmt.Println("牛逼")
	}
	fmt.Println(bar_2)
	fmt.Println(e_2)
	fmt.Println(bar_3)
	fmt.Println(e_3)
	fmt.Println("==============================================")

	// argv_3 := []string{
	// 	"./foo run --bar=baz --e=1",
	// }
	// err = rootCmd.ParseFlags(argv_3)
	// if err != nil {
	// 	fmt.Println("牛逼")
	// }
	rootCmd.SetArgs(argv_2[1:])
	rootCmd.Execute()
	fmt.Println()
	fmt.Println(bar_2)
	fmt.Println(e_2)
	fmt.Println(bar_3)
	fmt.Println(e_3)
	fmt.Println("==============================================")
	rootCmd.SetArgs(argv_3[1:])
	rootCmd.Execute()
	fmt.Println()
	fmt.Println(bar_2)
	fmt.Println(e_2)
	fmt.Println(bar_3)
	fmt.Println(e_3)
	fmt.Println("==============================================")
}
