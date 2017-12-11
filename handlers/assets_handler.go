package handlers

import (
	"net/http"
	"path"

	"github.com/shreyaganguly/chirper/packed"
)

const (
	applicationJavaScript = "application/javascript"
	contentType           = "Content-Type"
	textCSS               = "text/css"
	imageSVG              = "image/svg+xml"
	imageJPG              = "image/jpeg"
	imagePNG              = "image/png"
)

var assetMap = map[string]string{
	".js":  applicationJavaScript,
	".css": textCSS,
	".svg": imageSVG,
	".jpg": imageJPG,
	".png": imagePNG,
}

//AssetHandler handles the asset and set the header content
func AssetHandler(w http.ResponseWriter, r *http.Request) {
	asset := r.URL.Path[1:]

	ext := path.Ext(r.URL.Path)
	data, err := packed.Asset(asset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	mt, ok := assetMap[ext]
	if ok {
		w.Header().Set(contentType, mt)
	} else {
		w.Header().Set(contentType, "text/plain")
	}

	w.Header().Set("Cache-Control", "public, max-age=31536000")
	w.Write(data)
}
