package controller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var (
	CLD_API_KEY string
	CLD_NAME    string
	CLD_SECRET  string
	cld         *cloudinary.Cloudinary
)

func init() {
	CLD_SECRET := os.Getenv("CLD_SECRET")
	CLD_NAME := os.Getenv("CLD_NAME")
	CLD_API_KEY := os.Getenv("CLD_API_KEY")

	var err error
	cld, err = cloudinary.NewFromParams(CLD_NAME, CLD_API_KEY, CLD_SECRET)
	if err != nil {
		fmt.Println("Error initializing Cloudinary:", err)
	}

	// Load environment variables
	CLD_API_KEY = os.Getenv("CLD_API_KEY")
	CLD_NAME = os.Getenv("CLD_NAME")
	CLD_SECRET = os.Getenv("CLD_SECRET")

	// Check if any of the environment variables are empty
	if CLD_API_KEY == "" || CLD_NAME == "" || CLD_SECRET == "" {
		fmt.Println("Error: Cloudinary environment variables are not set properly")
		return
	}

	cld, err = cloudinary.NewFromParams(CLD_NAME, CLD_API_KEY, CLD_SECRET)
	if err != nil {
		fmt.Println("Error initializing Cloudinary:", err)
		return
	}
}

func UploadImagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid request method"}`, http.StatusMethodNotAllowed)
		return
	}

	var req UploadImagesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Failed to decode request"}`, http.StatusBadRequest)
		return
	}

	if len(req.Input.UserInput.Files) == 0 {
		http.Error(w, `{"error": "No file data provided"}`, http.StatusBadRequest)
		return
	}

	var imageUrls []string
	for _, fileData := range req.Input.UserInput.Files {
		decodedFile, err := base64.StdEncoding.DecodeString(fileData)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode file"}`, http.StatusBadRequest)
			return
		}
		fmt.Println("RESPfff: ", fileData)

		resp, err := cld.Upload.Upload(r.Context(), bytes.NewReader(decodedFile), uploader.UploadParams{Folder: "EventImages"})
		if err != nil {
			http.Error(w, `{"error": "Failed to upload to Cloudinary"}`, http.StatusInternalServerError)
			return
		}
		fmt.Println("RESP: ", resp)
		imageUrls = append(imageUrls, resp.SecureURL)
	}
	fmt.Println("this is from req.Files: ", req.Input.UserInput.Files)

	response := UploadImagesResponse{
		ImageURLs: imageUrls,
	}

	fmt.Println("Responseeeeee: ", response)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}
