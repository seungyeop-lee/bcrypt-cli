package cmd

import (
	"fmt"
	"github.com/seungyeop-lee/bcrypt-cli/app"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "bcrypt-cli",
}

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "generate hash",
	Example: "bcrypt-cli generate -p myPassword",
	RunE: func(cmd *cobra.Command, args []string) error {
		password, _ := cmd.Flags().GetString("password")
		fmt.Printf("input password: %s\n", password)

		cost, _ := cmd.Flags().GetInt("cost")
		generator := app.NewGenerator(cost)
		fmt.Printf("input cost: %d\n", generator.Cost())

		generatedHash, err := generator.Generate(password)
		if err != nil {
			return err
		}

		fmt.Printf("generated hash: %s\n", generatedHash)

		return nil
	},
}

var costCmd = &cobra.Command{
	Use:     "cost",
	Short:   "calculate cost from hash",
	Example: "bcrypt-cli cost -i '$2a$10$yQiVOOeFcuwYRSxjmXxb1.l5EDXn66Mqg0w7Hr/W6vbV1SYcAnwt6'",
	RunE: func(cmd *cobra.Command, args []string) error {
		hash, _ := cmd.Flags().GetString("hash")
		fmt.Printf("input hash: %s\n", hash)

		checker := app.NewChecker()
		cost, err := checker.Cost(hash)
		if err != nil {
			return err
		}

		fmt.Printf("cost: %d\n", cost)

		return nil
	},
}

var checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "check valid password and hash",
	Example: "bcrypt-cli check -p qwer -i '$2a$10$yQiVOOeFcuwYRSxjmXxb1.l5EDXn66Mqg0w7Hr/W6vbV1SYcAnwt6'",
	Run: func(cmd *cobra.Command, args []string) {
		password, _ := cmd.Flags().GetString("password")
		fmt.Printf("input password: %s\n", password)

		hash, _ := cmd.Flags().GetString("hash")
		fmt.Printf("input hash: %s\n", hash)

		checker := app.NewChecker()
		err := checker.Check(password, hash)

		if err != nil {
			fmt.Printf("check result: %s\n", "success")
		} else {
			fmt.Printf("check result: %s\n", "failure")
		}
	},
}

func Execute() {
	generateCmd.Flags().StringP("password", "p", "", "input password")
	generateCmd.Flags().IntP("cost", "c", 10, "input cost for hashing")
	_ = generateCmd.MarkFlagRequired("password")
	rootCmd.AddCommand(generateCmd)

	costCmd.Flags().StringP("hash", "i", "", "input hash")
	_ = costCmd.MarkFlagRequired("hash")
	rootCmd.AddCommand(costCmd)

	checkCmd.Flags().StringP("hash", "i", "", "input hash")
	checkCmd.Flags().StringP("password", "p", "", "input password")
	_ = checkCmd.MarkFlagRequired("hash")
	_ = checkCmd.MarkFlagRequired("password")
	rootCmd.AddCommand(checkCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
