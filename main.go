/*
Create a CLI backup tool in Golang with the following features:
1. Running with no flag should copy source directory into a backup directory.

2. Once a directory has been copied, the original directory should have meta-
data to track which files were changed, so that only changed files are
recopied in the backup directory next time.

3. Provide a flag to encrypt files while backing up. Additional flags can be
used in combination with encryption to recursively encrypt, selectively
encrypt files, etc.
4. Create a log file in the backup directory that logs all backup activities.
5. Use an efficient backing mechanism such that if a subdirectory has already been backed up, only new directories/files are copied.


no flags => copy source directory into a backup directory
			have one default backup directory (?)


*/


package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func backupDirectory(sourceDir, backupDir string) error {
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		return fmt.Errorf("source dir does not exist, %w", err)
	}

	timestamp := time.Now().Format("20060102-150405") 
	versionedBackupDir := filepath.Join(backupDir, timestamp)
	err := os.MkdirAll(versionedBackupDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating backup: %w", err)
	}

	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil 
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(versionedBackupDir, relPath)

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(destPath, data, 0644)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error during backup:", err)
	}

	// 4. Logging (Omitted for brevity)
	timestamp := time.Now().Format("20060102-150405")
	logFile := filepath.Join(backupDir, timestamp + ".dat")
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error creating log file: %w", err)
	}
	f.Close()


	return nil

	return nil
}

func main() {

	clargs := os.Args[1:]

	
	var sourceDir, backupDir string

    // Check if any argument starts with a hyphen
    hasFlag := false
    for _, arg := range clargs {
        if arg[0] == '-' {
            hasFlag = true
			fmt.Println()
            break
        }
    }

    //no flag present => arguments for backup and source dir given
    if !hasFlag && len(clargs) == 2 {
        sourceDir = clargs[0]
        backupDir = clargs[1]
    } else { //flag present => toDo
        fmt.Println("Reached here")
        os.Exit(1)
    }

    // Print the values of sourceDir and backupDir
    fmt.Println("Source Directory:", sourceDir)
    fmt.Println("Backup Directory:", backupDir)


	err := backupDirectory(sourceDir, backupDir)
	if err != nil {
		fmt.Println("Backup failed:", err)
	} else {
		fmt.Println("Backup successful!")
	}
}