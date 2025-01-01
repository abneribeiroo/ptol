package internal

import (
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func WatchAndRun(directory string, delay time.Duration) {
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

        // Use Viper to get the script to run
        script := viper.GetString("script")
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
                    resetTimer()
                }
            case err := <-watcher.Errors:
                log.Println("Error:", err)
            }
        }
    }()

    if err := watcher.Add(directory); err != nil {
        log.Fatal(err)
    }

    restartProcess() // Start the initial process
    <-done
}
