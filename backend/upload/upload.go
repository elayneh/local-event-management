package upload

import (
	"fmt"
	"backend/hasura"
	"net/http"
	"os"
	"path/filepath"
	"io"
	"log"
	"time"
)

// Define the upload directory
const uploadPath = "./public/uploads/"

// UploadFileHandler handles file uploads and saves them to the server
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form to retrieve file
	err := r.ParseMultipartForm(10 << 20) // Limit upload size to 10MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		log.Println("Error parsing form:", err)
		return
	}

	// Get the file from the form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		log.Println("Error retrieving file:", err)
		return
	}
	defer file.Close()

	// Extract the original file name and extension
	fileName := fileHeader.Filename
	ext := filepath.Ext(fileName)

	if ext == "" {
		// Default extension if not found
		ext = ".bin"
	}

	// Log extracted file name and extension
	log.Println("Original file name:", fileName)
	log.Println("Extracted extension:", ext)

	// Create a unique file name with timestamp and extension
	newFileName := fmt.Sprintf("uploaded_file_%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadPath, newFileName)

	// Ensure the upload directory exists
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
		log.Println("Error creating upload directory:", err)
		return
	}

	// Create a new file on the server
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		log.Println("Error creating file:", err)
		return
	}
	defer dst.Close()

	// Copy the file content to the newly created file
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		log.Println("Error saving file:", err)
		return
	}

	// Update the database with the filename (Implement this function)
	err = updateProductImage(newFileName)
	if err != nil {
		http.Error(w, "Error updating database", http.StatusInternalServerError)
		log.Println("Error updating database:", err)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded successfully: %s", newFileName)
}

// updateProductImage updates the database with the uploaded file's name
func updateProductImage(fileName string) error {
	query := `
		mutation insertProduct($image: String!) {
			insert_products(
				objects: {
					image: $image
				}
			) {
				returning {
					image
				}
			}
		}
	`

	variables := map[string]interface{}{
		"image": fileName,
	}

	_, err := hasura.SendGraphQLRequest(query, variables)
	if err != nil {
		return err
	}

	return nil
}
