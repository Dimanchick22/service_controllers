/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CLIConfig хранит аргументы командной строки
type CLIConfig struct {
	ConfigFile string // Путь к файлу конфигурации
}

var cfg CLIConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "service_controllers",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the service controllers",
	Long: `Starts the service controllers and performs necessary actions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Добавьте здесь код для запуска вашего приложения
		fmt.Println("Starting the service controllers...")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately
