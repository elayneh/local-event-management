package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const uploadPath = "./uploads"

// UploadFileHandler handles file uploads
func UploadFileHandler(c *gin.Context) {
	// Only allow POST method
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	// Parse the multipart form
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB file size limit
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing multipart form"})
		return
	}

	// Retrieve the file from the form data
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving file"})
		return
	}
	defer file.Close()

	// Validate file type
	fileExtension := filepath.Ext(fileHeader.Filename)
	if !isValidFileType(fileExtension) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
		return
	}

	// Create the uploads directory if it doesn't exist
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err = os.Mkdir(uploadPath, 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating directory"})
			return
		}
	}

	// Generate a unique file name
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExtension)
	filepath := filepath.Join(uploadPath, filename)

	// Save the file
	destFile, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	// Return the file URL as the response
	fileURL := fmt.Sprintf("/uploads/%s", filename)
	c.JSON(http.StatusOK, gin.H{"url": fileURL})
}

// isValidFileType validates file extensions
func isValidFileType(extension string) bool {
	switch extension {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}
