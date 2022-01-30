package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mngr/models"
	"mngr/utils"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func RegisterVideoEndpoints(router *gin.Engine) {
	router.GET("/videos/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		//id = "QLma6mWR3V8"
		recordingFolderPath, _ := utils.GetRecordingFolderPath()
		files, _ := ioutil.ReadDir(path.Join(recordingFolderPath, id))
		var list = make([]*models.VideoFile, 0)
		for _, file := range files {
			videoFile := models.VideoFile{}
			videoFile.SourceId = id
			videoFile.Name = file.Name()
			videoFile.Path = path.Join("/playback", id, file.Name()) //utils.RelativePlaybackFolderPath + "/" + file.Name()
			videoFile.Size = utils.Round(float64(file.Size()) * 0.000001)
			videoFile.CreatedAt = strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			videoFile.ModifiedAt = utils.FromDateToString(file.ModTime())
			list = append(list, &videoFile)
		}
		ctx.JSON(http.StatusOK, list)
	})

	router.DELETE("/videos/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var fileNames []string
		ctx.BindJSON(&fileNames)
		//id = "QLma6mWR3V8"
		recordingFolderPath, _ := utils.GetRecordingFolderPath()
		for _, fileName := range fileNames {
			err := os.Remove(path.Join(recordingFolderPath, id, fileName))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
		}
		ctx.JSON(http.StatusOK, nil)
	})

	//router.GET("/streaming/:id", func(ctx *gin.Context) {
	//	id := ctx.Param("id")
	//	rep := reps.StreamingRepository{Connection: utils.ConnSources}
	//	model, err := rep.Get(id)
	//	if err != nil {
	//		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	//		return
	//	}
	//	ctx.JSON(http.StatusOK, model)
	//})
}
