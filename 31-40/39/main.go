package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	// NewWatcher establishes a new watcher with the underlying OS and begins waiting for events.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	// Close removes all watches and closes the events channel.
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			// Return, name of file, type of operation (Create Write Remove Rename Chmod) and error if any.
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Notify Create Write Remove Rename Chmod events
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					// Notify write file events
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Watch for file or directory for change.
	//Add starts watching the named file or directory (non-recursively).
	err = watcher.Add("/tmp")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
