package internal

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// export: BatchFiles organizes files by given extension into batches and saves them into folders starting from the start date.
// mama
func BatchFiles(dest string, src string, fileExt string, batchSize int, startDateStr string) error {
	// Parse the start date string into a time.Time object
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return fmt.Errorf("invalid start date format: %w", err)
	}

	// Collect all files matching the given extension in the source directory
	var files []string
	err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == fileExt {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking the path %v: %w", src, err)
	}

	// Batch the files into folders
	currentBatch := []string{}
	currentDate := startDate

	for i, file := range files {
		currentBatch = append(currentBatch, file)
		// Check if the current batch has reached the batch size or if we are at the last file
		if len(currentBatch) == batchSize || i == len(files)-1 {
			// Create the folder for the batch
			batchFolder := filepath.Join(dest, currentDate.Format("2006-01-02"))
			err = os.MkdirAll(batchFolder, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory %v: %w", batchFolder, err)
			}

			// Move files into the batch folder
			for _, filePath := range currentBatch {
				fileName := filepath.Base(filePath)
				newPath := filepath.Join(batchFolder, fileName)
				err = os.Rename(filePath, newPath)
				if err != nil {
					return fmt.Errorf("failed to move file %v: %w", filePath, err)
				}
			}

			// Prepare for the next batch
			currentBatch = []string{}
			currentDate = currentDate.Add(24 * time.Hour) // Move to the next day
		}
	}

	return nil
}
