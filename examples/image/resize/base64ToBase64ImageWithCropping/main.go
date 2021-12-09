package main

import (
	"fmt"
	"os"

	"go.m3o.com/image"
)

// Resize an image on the fly without storing it (by sending and receiving a base64 encoded image), or resize and upload depending on parameters.
// If one of width or height is 0, the image aspect ratio is preserved.
// Optional cropping.
// To use the file parameter you need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func main() {
	imageService := image.NewImageService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := imageService.Resize(&image.ResizeRequest{
		Base64: "data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==",
		CropOptions: &image.CropOptions{
			Height: 50,
			Width:  50,
		},
		Height: 100,
		Width:  100,
	})
	fmt.Println(rsp, err)
}
