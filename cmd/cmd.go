package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	logger "github.com/seungyeop-lee/bcrypt-cli/log"

	"github.com/seungyeop-lee/bcrypt-cli/app"
	"github.com/spf13/cobra"
)

var executeFileName = filepath.Base(os.Args[0])

var rootCmd = &cobra.Command{
	Use: executeFileName,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		logger.IsVerbose = verbose
	},
}

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen", "g"},
	Short:   "Generate hash",
	Example: fmt.Sprintf("%s generate -p myPassword", executeFileName),
	RunE: func(cmd *cobra.Command, args []string) error {
		password, _ := cmd.Flags().GetString("password")
		logger.Info(fmt.Sprintf("input password: %s", password))

		cost, _ := cmd.Flags().GetInt("cost")
		generator := app.NewGenerator(cost)
		logger.Info(fmt.Sprintf("input cost: %d", generator.Cost()))

		generatedHash, err := generator.Generate(password)
		if err != nil {
			return err
		}

		logger.Info(fmt.Sprintf("generated hash: %s", generatedHash))

		fmt.Print(generatedHash)

		return nil
	},
}

var costCmd = &cobra.Command{
	Use:     "cost",
	Short:   "Calculate cost from hash",
	Example: fmt.Sprintf("%s cost -i '$2a$10$iJ/CnWkU8efsEKnnR14vl.MYVfy9adcAXxpPeiLrGaHTaKx5JBbse'", executeFileName),
	RunE: func(cmd *cobra.Command, args []string) error {
		hash, _ := cmd.Flags().GetString("hash")
		logger.Info(fmt.Sprintf("input hash: %s", hash))

		checker := app.NewChecker()
		cost, err := checker.Cost(hash)
		if err != nil {
			return err
		}

		logger.Info(fmt.Sprintf("cost: %d", cost))

		fmt.Print(cost)

		return nil
	},
}

var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"c"},
	Short:   "Check valid password and hash",
	Long:    "As a result of checking the password and hash, 0 if valid and 1 if invalid are returned as the status code",
	Example: fmt.Sprintf("%s check -p myPassword -i '$2a$10$iJ/CnWkU8efsEKnnR14vl.MYVfy9adcAXxpPeiLrGaHTaKx5JBbse'", executeFileName),
	Run: func(cmd *cobra.Command, args []string) {
		password, _ := cmd.Flags().GetString("password")
		logger.Info(fmt.Sprintf("input password: %s", password))

		hash, _ := cmd.Flags().GetString("hash")
		logger.Info(fmt.Sprintf("input hash: %s", hash))

		checker := app.NewChecker()
		err := checker.Check(password, hash)

		if err == nil {
			logger.Info(fmt.Sprintf("check result: %s", "success"))
		} else {
			logger.Info(fmt.Sprintf("check result: %s", "failure"))
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "run by verbose mode (for debug)")

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
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
