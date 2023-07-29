/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crosscheck/aws/checks"
	"crosscheck/general"
	"fmt"
	"io"
	"os"

	crossplanev1beta1 "github.com/crossplane-contrib/provider-aws/apis/s3/v1beta1"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

type Rule struct {
	Name  string `yaml:"name"`
	Check string `yaml:"check"`
}

type Rules struct {
	Rules []Rule `yaml:"rules"`
}

type Check interface {
	ScanResourceConf(bucket *crossplanev1beta1.Bucket) (general.CheckResult, error)
	GetInspectedKey() string
}

// implement rule check mapping
var ruleCheckMapping = map[string]func() checks.Check{
	"NewS3RestrictPublicBucketsCheck": checks.NewS3RestrictPublicBucketsCheck,
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		rulesFilePath, err := cmd.Flags().GetString("rulesFile")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("scan called with rulesFile: " + rulesFilePath)

		rulesFile, err := os.Open(rulesFilePath)
		if err != nil {
			fmt.Errorf("failed to open rules file: %v", err)
		}
		defer rulesFile.Close()

		// Read and parse the rules from the file
		rulesData, err := io.ReadAll(rulesFile)
		if err != nil {
			fmt.Errorf("failed to read rules file: %v", err)
		}

		var rules Rules
		err = yaml.Unmarshal(rulesData, &rules)
		if err != nil {
			fmt.Errorf("failed to unmarshal rules YAML: %v", err)
		}

		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("scan called with file: " + filePath)

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Errorf("failed to load demo_bucket.yaml: %v", err)
		}

		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Errorf("failed to read demo_bucket.yaml: %v", err)
		}

		var bucket crossplanev1beta1.Bucket
		err = yaml.Unmarshal(data, &bucket)
		if err != nil {
			fmt.Errorf("failed to unmarshal YAML: %v", err)
		}

		for _, rule := range rules.Rules {
			checkFunc, ok := ruleCheckMapping[rule.Check]
			if !ok {
				fmt.Errorf("unknown check function: %v", rule.Check)
				continue
			}

			check := checkFunc()
			result, err := check.ScanResourceConf(&bucket)
			if err != nil {
				fmt.Errorf("failed to scan resource: %v", err)
			}
			if result == general.PASSED {
				fmt.Println(rule.Name + " scan succeeded")
			} else {
				fmt.Println(rule.Name + " scan failed")
			}
		}

		// check := checks.NewS3RestrictPublicBucketsCheck()
		// result, err := check.ScanResourceConf(&bucket)
		// if err != nil {
		// 	fmt.Errorf("failed to scan resource: %v", err)
		// }
		// if result == general.PASSED {
		// 	fmt.Println("scan succeeded")
		// } else {
		// 	fmt.Println("scan failed")
		// }
	},
}

func init() {
	var filePath string
	var rulesFilePath string

	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")

	scanCmd.Flags().StringVarP(&rulesFilePath, "rulesFile", "r", "", "Rules file path")
}
