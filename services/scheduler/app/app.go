package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	docs "github.com/MichalK12344321/golang/services/scheduler/docs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
  triggerCollectionStatus := TriggerCollectionStatusType{CollectionId: CollectionIdType(cid.String())}
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

  collectionDataGatewayHost := os.Getenv("COLLECTION_DATA_GATEWAY_HOST")
  if collectionDataGatewayHost == "" {
    collectionDataGatewayHost = "collectiondatagateway"
}

  collectionDataGatewayPort := os.Getenv("COLLECTION_DATA_GATEWAY_PORT")
  if collectionDataGatewayPort == "" {
    collectionDataGatewayPort = "5555"
}
  collectionDataGatewayURL := fmt.Sprintf("http://%s:%s/api/v1/collections", collectionDataGatewayHost, collectionDataGatewayPort)

  resp, err := http.Get(collectionDataGatewayURL)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read response from Get request Collection Data Gateway"})
    return
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to read response from response body Collection Data Gateway"})
    return
  }

  var collections CollectionsType
  if err := json.Unmarshal(body, &collections); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response from Collection Data Gateway"})
    return
  }

  c.JSON(http.StatusOK, collections)
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

type CollectionIdType string

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
