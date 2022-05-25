package api

import (
	"github.com/gin-gonic/gin"
	"mngr/reps"
	"mngr/utils"
	"net/http"
	"sort"
)

func RegisterFrImagesEndpoints(router *gin.Engine, rb *reps.RepoBucket) {
	router.GET("/frimagesfolders/:id/:date", func(ctx *gin.Context) {
		sourceId := ctx.Param("id")
		date := ctx.Param("date")
		config, _ := rb.ConfigRep.GetConfig()
		odPath := utils.GetHourlyFrImagesPathBySourceId(config, sourceId, date)
		items, _ := newTree(odPath, true)
		ctx.JSON(http.StatusOK, items)
	})
	// it has potential security risk
	router.POST("frimages", func(ctx *gin.Context) {
		var model ImagesParams
		if err := ctx.ShouldBindJSON(&model); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		config, _ := rb.ConfigRep.GetConfig()
		frhRep := reps.FrHandlerRepository{Config: config}
		jsonObjects := frhRep.GetJsonObjects(model.SourceId, model.RootPath, true)
		items := make([]*ImageItem, 0)
		for _, jsonObject := range jsonObjects {
			fr := jsonObject.FaceRecognition
			item := &ImageItem{Id: fr.Id, ImagePath: fr.ImageFileName, CreatedAt: jsonObject.FaceRecognition.CreatedAt}
			items = append(items, item)
		}
		sort.Slice(items, func(i, j int) bool {
			t1 := utils.StringToTime(items[i].CreatedAt)
			t2 := utils.StringToTime(items[j].CreatedAt)
			return t1.After(t2)
		})
		ctx.JSON(http.StatusOK, items)
	})
}
