package scheduler

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  types "github.com/MichalK12344321/golang/services/common/types"
  swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
  docs "github.com/MichalK12344321/golang/services/scheduler/docs"
  interfaces "github.com/MichalK12344321/golang/services/scheduler/interfaces"
)

type Server struct {
  config *Config
  router *gin.Engine
  messageQueue interfaces.TriggerMessageQueue
}

func NewServer(config *Config, messageQueue interfaces.TriggerMessageQueue) *Server {
  s := Server{}
  s.messageQueue = messageQueue
  s.router = gin.Default()

  docs.SwaggerInfo.BasePath = "/api/v1"

  v1 := s.router.Group("/api/v1")
  {
    v1.POST("/trigger", s.triggerCollectionHandler)
//    v1.GET("/:collectionId/log", s.forwardGetLogHandler)
    v1.GET("/collections", s.forwardListCollectionsHandler)
  }
  s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
 
  return &s
}

func (s *Server) Start() {
  s.router.Run(":8080")
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
func (s *Server) triggerCollectionHandler(c *gin.Context) {
  var triggerCollectionInfo TriggerCollectionInfoType
  if err := c.BindJSON(&triggerCollectionInfo); err != nil {
    c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
    return
  }
  cid := uuid.New().String()

  err := s.messageQueue.SendTrigger(types.CollectionIdType(cid), triggerCollectionInfo.Ssh)
  if err != nil {
    c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: fmt.Errorf("Could not send trigger to message queue: %v", err).Error()})
    return
  }

  triggerCollectionStatus := TriggerCollectionStatusType{CollectionId: types.CollectionIdType(cid)}
  triggerCollectionStatus.Message = "Triggered collection: " + cid
  c.JSON(http.StatusOK, triggerCollectionStatus)
}

// @Summary Forward Get Collections
// @Description Forward Get List of collection info to data-collection-gateway
// @Tags scheduler
// @Accept json
// @Produce json
// @Success 200 {object} types.CollectionsType
// @Router /collections [get]
func (s *Server) forwardListCollectionsHandler(c *gin.Context)  {
	//TODO:
	// get config for Collection Data Gateway URL and PORT
	// Forward GET request to Collection Data Gateway
	// handle response
}

type TriggerCollectionInfoType struct {
  Tag string `json:"tag" example:"Testing"`
  Ssh types.SshInfoType `json:"ssh" binding:"required"`
}

type TriggerCollectionStatusType struct {
  Message string `json:"message" binding:"required"`
  CollectionId types.CollectionIdType `json:"collectionId" binding:"required"`
}
