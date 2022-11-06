package swagger

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed dist
var fsys embed.FS

func SwaggerUI() gin.HandlerFunc {
	dist, _ := fs.Sub(fsys, "dist")

	staticServer := http.FileServer(http.FS(dist))
	sh := http.StripPrefix("/swaggerui", staticServer)

	return gin.WrapH(sh)

}
