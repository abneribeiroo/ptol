package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var delay time.Duration

var rootCmd = &cobra.Command{
	Use:   "file-watcher",
	Short: "File watcher CLI",
	Long:  "CLI to watch and restart processes on file changes.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please provide the script to run")
		}
		script := args[0]
		watchAndRun(script)
	},
}

func init() {
	rootCmd.PersistentFlags().DurationVar(&delay, "delay", 2*time.Second, "Delay before restarting the process")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func watchAndRun(script string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	var cmd *exec.Cmd
	var mu sync.Mutex
	var timer *time.Timer

	restartProcess := func() {
		mu.Lock()
		defer mu.Unlock()

		if cmd != nil {
			_ = cmd.Process.Kill()
		}
		cmd = startProcess(script)
	}

	resetTimer := func() {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(delay, restartProcess)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					// log.Println("File modified:", event.Name)
					resetTimer()
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(script)
	if err != nil {
		log.Fatal(err)
	}

	restartProcess() // Start the initial process
	<-done
}

func startProcess(script string) *exec.Cmd {
	clearTerminal()

	var cmd *exec.Cmd
	if script[len(script)-3:] == ".go" {
		cmd = exec.Command("go", "run", script)
	} else if script[len(script)-3:] == ".js" {
		cmd = exec.Command("node", script)
	} else {
		log.Fatalf("Unsupported file type: %s", script)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error starting process: %v", err)
	}

	return cmd
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
