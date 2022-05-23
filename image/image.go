package image

import (
	"go.m3o.com/client"
)

type Image interface {
	Convert(*ConvertRequest) (*ConvertResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Resize(*ResizeRequest) (*ResizeResponse, error)
	Upload(*UploadRequest) (*UploadResponse, error)
}

func NewImageService(token string) *ImageService {
	return &ImageService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ImageService struct {
	client *client.Client
}

// Convert an image from one format (jpeg, png etc.) to an other either on the fly (from base64 to base64),
// or by uploading the conversion result.
// To use the file parameter you need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func (t *ImageService) Convert(request *ConvertRequest) (*ConvertResponse, error) {

	rsp := &ConvertResponse{}
	return rsp, t.client.Call("image", "Convert", request, rsp)

}

// Delete an image previously uploaded.
func (t *ImageService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("image", "Delete", request, rsp)

}

// Resize an image on the fly without storing it (by sending and receiving a base64 encoded image), or resize and upload depending on parameters.
// If one of width or height is 0, the image aspect ratio is preserved.
// Optional cropping.
// To use the file parameter you need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func (t *ImageService) Resize(request *ResizeRequest) (*ResizeResponse, error) {

	rsp := &ResizeResponse{}
	return rsp, t.client.Call("image", "Resize", request, rsp)

}

// Upload an image by either sending a base64 encoded image to this endpoint or a URL.
// To resize an image before uploading, see the Resize endpoint.
// To use the file parameter you need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func (t *ImageService) Upload(request *UploadRequest) (*UploadResponse, error) {

	rsp := &UploadResponse{}
	return rsp, t.client.Call("image", "Upload", request, rsp)

}

type ConvertRequest struct {
	// base64 encoded image to resize,
	Base64 string `json:"base64,omitempty"`
	// The image file to convert
	File string `json:"file,omitempty"`
	// output name of the image including extension, ie. "cat.png"
	Name string `json:"name,omitempty"`
	// make output a URL and not a base64 response
	OutputUrl bool `json:"outputURL,omitempty"`
	// url of the image to resize
	Url string `json:"url,omitempty"`
}

type ConvertResponse struct {
	Base64 string `json:"base64,omitempty"`
	Url    string `json:"url,omitempty"`
}

type CropOptions struct {
	// Crop anchor point: "top", "top left", "top right",
	// "left", "center", "right"
	// "bottom left", "bottom", "bottom right".
	// Optional. Defaults to center.
	Anchor string `json:"anchor,omitempty"`
	// height to crop to
	Height int32 `json:"height,omitempty"`
	// width to crop to
	Width int32 `json:"width,omitempty"`
}

type DeleteRequest struct {
	// url of the image to delete e.g. https://cdn.m3ocontent.com/micro/images/micro/41e23b39-48dd-42b6-9738-79a313414bb8/cat.jpeg
	Url string `json:"url,omitempty"`
}

type DeleteResponse struct {
}

type Point struct {
	X int32 `json:"x,omitempty"`
	Y int32 `json:"y,omitempty"`
}

type Rectangle struct {
	Min *Point `json:"min,omitempty"`
	Max *Point `json:"max,omitempty"`
}

type ResizeRequest struct {
	// url of the image to resize
	Url   string `json:"url,omitempty"`
	Width int64  `json:"width,string,omitempty"`
	// base64 encoded image to resize,
	Base64 string `json:"base64,omitempty"`
	// optional crop options
	// if provided, after resize, the image
	// will be cropped
	CropOptions *CropOptions `json:"cropOptions,omitempty"`
	// The image file to resize
	File   string `json:"file,omitempty"`
	Height int64  `json:"height,string,omitempty"`
	// output name of the image including extension, ie. "cat.png"
	Name string `json:"name,omitempty"`
	// make output a URL and not a base64 response
	OutputUrl bool `json:"outputURL,omitempty"`
}

type ResizeResponse struct {
	Base64 string `json:"base64,omitempty"`
	Url    string `json:"url,omitempty"`
}

type UploadRequest struct {
	// URL of the image to upload
	Url string `json:"url,omitempty"`
	// Base64 encoded image to upload,
	Base64 string `json:"base64,omitempty"`
	// The image file to upload
	File string `json:"file,omitempty"`
	// Output name of the image including extension, ie. "cat.png"
	Name string `json:"name,omitempty"`
}

type UploadResponse struct {
	Url string `json:"url,omitempty"`
}
