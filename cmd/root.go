package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abneribeiroo/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	delay   time.Duration
	rootCmd = &cobra.Command{
		Use:   "ptol",
		Short: "File watcher CLI",
		Long:  "CLI to watch and restart processes on file changes.",
		Run:   runRootCmd,
	}
)

// directory := viper.GetString("directory")
// delay := viper.GetDuration("delay")

// WatchAndRun(directory, delay)

func init() {
	viper.SetConfigName("config") // Nome do arquivo de configuração (sem extensão)
	viper.SetConfigType("yaml")   // Tipo do arquivo de configuração
	viper.AddConfigPath(".")      // Caminho para o diretório onde o arquivo de configuração está localizado

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
	}
	rootCmd.PersistentFlags().DurationVar(&delay, "delay", 2*time.Second, "Delay before restarting the process")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runRootCmd(cmd *cobra.Command, args []string) {
	// if len(args) < 1 {
	// 	log.Fatal("Please provide the script to run")
	// }
	directory := viper.GetString("directory")
	delay := viper.GetDuration("delay")

	// // WatchAndRun(directory, delay)
	// script := args[0]
	internal.WatchAndRun(directory, delay)
}
