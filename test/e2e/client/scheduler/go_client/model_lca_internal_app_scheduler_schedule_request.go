/*
 * scheduler
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package scheduler_client

type LcaInternalAppSchedulerScheduleRequest struct {
	Host string `json:"host"`
	Password string `json:"password"`
	Port int32 `json:"port"`
	Script string `json:"script"`
	User string `json:"user"`
}
