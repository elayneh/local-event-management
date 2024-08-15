package controller

import (
	"backend/utils"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/machinebox/graphql"
)

func HandleAddProduct(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	var reqBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       string `json:"price"`
		Category    string `json:"category"`
		Image       string `json:"image"`
	}
	if err := json.Unmarshal(body, &reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	imageData, err := decodeBase64Image(reqBody.Image)
	if err != nil {
		http.Error(w, "Invalid base64 image data", http.StatusBadRequest)
		return
	}

	imageURL, err := uploadToCloudinary(imageData)
	if err != nil {
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}

	req := graphql.NewRequest(`
		mutation($object: products_insert_input!) {
			insert_products_one(object: $object) {
				id
				name
				description
				price
				category
				image_url
			}
		}
	`)

	req.Var("object", map[string]interface{}{
		"name":        reqBody.Name,
		"description": reqBody.Description,
		"price":       reqBody.Price,
		"category":    reqBody.Category,
		"image_url":   imageURL,
	})

	var respData struct {
		InsertProductsOne struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Price       string `json:"price"`
			Category    string `json:"category"`
			ImageURL    string `json:"image_url"`
		} `json:"insert_products_one"`
	}

	client := utils.Client()

	err = client.Run(context.Background(), req, &respData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error executing GraphQL mutation: ", err)
		return
	}
	if respData.InsertProductsOne.ID == "" {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product uploaded successfully"))
}

func decodeBase64Image(base64Image string) ([]byte, error) {

	if strings.Contains(base64Image, "base64,") {
		base64Image = strings.Split(base64Image, "base64,")[1]
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}

func uploadToCloudinary(imageData []byte) (string, error) {
	CLD_NAME := os.Getenv("CLD_NAME")
	CLD_API_KEY := os.Getenv("CLD_API_KEY")
	CLD_SECRET := os.Getenv("CLD_SECRET")

	cld, _ := cloudinary.NewFromParams(CLD_NAME, CLD_API_KEY, CLD_SECRET)

	uploadParams := uploader.UploadParams{}
	uploadResult, err := cld.Upload.Upload(context.TODO(), imageData, uploadParams)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
