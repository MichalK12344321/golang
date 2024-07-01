package types

type ErrorResponse struct {
  Error string `json:"error" example:""`
}

type CollectionIdType string
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

type SshInfoType struct {
  Connection ConnectionInfoType `json:"connection" binding:"required"`
  Credentials CredentialsInfoType `json:"credentials" binding:"required"`
}

type ConnectionInfoType struct {
  Host string `json:"host" binding:"required" example:"example.host.pl"`
  Port int `json:"port" example:"2222"`
}

type CredentialsInfoType struct {
  User string `json:"user" binding:"required" example:"myuser"`
  Password string `json:"pass" binding:"required" example:"mypass"`
}
