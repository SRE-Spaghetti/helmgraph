package main

import (
	"fmt"
	"helmgraph/internal/manifest"
	"os"

	"github.com/spf13/cobra"
)

var (
	chartPath   string
	releaseName string
	namespace   string
	outputFile  string
)

var rootCmd = &cobra.Command{
	Use:   "helmgraph",
	Short: "Generate a Cypher script from a Helm chart.",
	Long:  `HelmGraph generates a Cypher script from a Helm chart that can be imported into Neo4j.`,
	Run: func(cmd *cobra.Command, args []string) {
		manifest, err := manifest.Generate(chartPath, releaseName, namespace)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if outputFile != "" {
			err := os.WriteFile(outputFile, []byte(manifest), 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Println(manifest)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&chartPath, "chart", "c", "", "Path to the Helm chart directory")
	rootCmd.Flags().StringVarP(&releaseName, "release", "r", "", "Release name")
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace")
	rootCmd.Flags().StringVarP(&outputFile, "out", "o", "", "Output file name (default: stdout)")

	rootCmd.MarkFlagRequired("chart")
	rootCmd.MarkFlagRequired("release")
}

func main() {
	Execute()
}
