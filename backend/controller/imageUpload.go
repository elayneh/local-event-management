package controller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary

func init() {
	// CLD_SECRET := os.Getenv("CLD_SECRET")
	// CLD_NAME := os.Getenv("CLD_NAME")
	// CLD_API_KEY := os.Getenv("CLD_API_KEY")

	var err error
	cld, err = cloudinary.NewFromParams("dci964xm7", "419582463847597", "R_IXfnVEVFIydSiWVJPaYT32j0w")
	if err != nil {
		fmt.Println("Error initializing Cloudinary:", err)
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
	fmt.Println("this is from backend1")
	var imageUrls []string
	for _, fileData := range req.Input.UserInput.Files {
		decodedFile, err := base64.StdEncoding.DecodeString(fileData)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode file"}`, http.StatusBadRequest)
			return
		}
		fmt.Println("RESPfff: ", fileData)

		resp, err := cld.Upload.Upload(r.Context(), bytes.NewReader(decodedFile), uploader.UploadParams{Folder: "tasks"})
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
