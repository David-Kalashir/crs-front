package minioapi

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/David-Kalashir/crs-front/foundation/web"
	"github.com/minio/minio-go/v7"
)

type api struct {
	minioclient *minio.Client
}

func newAPI(minioclient *minio.Client) *api {
	return &api{
		minioclient: minioclient,
	}
}

func (api *api) getobject(w http.ResponseWriter, r *http.Request) {
	bn := web.Param(r, "bucketname")
	on := web.Param(r, "objectname")
	object, err := api.minioclient.GetObject(context.Background(), bn, on, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer object.Close()
	// Retrieve object info for content-type and size
	info, err := object.Stat()
	if err != nil {
		http.Error(w, "Unable to retrieve object info", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Set Content-Type header based on the object's metadata
	w.Header().Set("Content-Type", info.ContentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size))

	// Display file inline in the browser (remove 'attachment' to allow inline display)
	w.Header().Set("Content-Disposition", "inline; filename="+on)

	// Stream the object content to the response
	if _, err := io.Copy(w, object); err != nil {
		http.Error(w, "Unable to write object to response", http.StatusInternalServerError)
		fmt.Println(err)
	}
}
