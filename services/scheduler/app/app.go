package app

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
  docs "github.com/MichalK12344321/golang/services/scheduler/docs"
)
// /trigger POST
// /collections GET
// /{collectionId}/log GET


func Start() {
  router := gin.Default()
	
  docs.SwaggerInfo.BasePath = "/api/v1"
  v1 := router.Group("/api/v1")
  {
    v1.POST("/trigger", triggerCollectionHandler)
//    v1.GET("/:collectionId/log", forwardGetLogHandler)
    v1.GET("/collections", forwardListCollectionsHandler)
  }
  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
  router.Run(":8080")
}

// @BasePath /api/v1

// @Summary Trigger Collection
// @Schemes
// @Description Start collection
// @Tags scheduler
// @Accept json
// @Produce json
// @Param triggerCollectionInfo body TriggerCollectionInfoType true "Trigger Collection Info"
// @Success 200 {object} TriggerCollectionStatusType
// @Router /trigger [post]
func triggerCollectionHandler(c *gin.Context)  {
  var triggerCollectionInfo TriggerCollectionInfoType
  if err := c.BindJSON(&triggerCollectionInfo); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  cid := uuid.New()
  triggerCollectionStatus := TriggerCollectionStatusType{CollectionId: CollectionIdType(cid)}
  triggerCollectionStatus.Message = "Triggered collection: " + cid.String()
  c.JSON(http.StatusOK, triggerCollectionStatus)
}

// @Summary Forward Collections
// @Description Forward Get List of collection info
// @Tags scheduler
// @Accept json
// @Produce json
// @Success 200 {object} CollectionsType
// @Router /collections [get]
func forwardListCollectionsHandler(c *gin.Context)  {
	//TODO:
	// get config for Collection Data Gateway URL and PORT
	// Forward GET request to Collection Data Gateway
	// handle response
}

type ConnectionInfoType struct {
  Host string `json:"host" binding:"required" example:"example.host.pl"`
  Port int `json:"port" example:"22"`
}

type CredentialsInfoType struct {
  User string `json:"user" binding:"required" example:"myuser"`
  Pass string `json:"pass" binding:"required" example:"mypass"`
}

type TriggerCollectionInfoType struct {
  Connection ConnectionInfoType `json:"connection" binding:"required"`
  Credentials CredentialsInfoType `json:"credentials" binding:"required"`
}

type CollectionIdType uuid.UUID

type TriggerCollectionStatusType struct {
  Message string `json:"message" binding:"required"`
  CollectionId CollectionIdType `json:"collectionId" binding:"required"`
}

type CollectionStatusType string
const (
	CollectionStatusCreated = "created"
	CollectionStatusRunning = "running"
	CollectionStatusSuccess = "success"
	CollectionStatusFailure = "failure"
)

type CollectionType struct {
  Id CollectionIdType `json:"id" example:""`
  Status CollectionStatusType `json:"status" example:""`
}

type CollectionsType []CollectionType
