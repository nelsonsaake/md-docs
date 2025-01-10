package env

import (
	"os"
	"fmt"
	"strconv"
)

func Get(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	v, _ := strconv.Atoi(Get(key))
	return v
}

func Set(key string, value any) {
	os.Setenv(key, fmt.Sprintf("%v", value))
}

func Password() string {
	return Get("PASSWORD")
}

func SetPassword(value any) {
	Set("PASSWORD", value)
}

func Db() string {
	return Get("DB")
}

func SetDb(value any) {
	Set("DB", value)
}

func DocPath() string {
	return Get("DOC_PATH")
}

func SetDocPath(value any) {
	Set("DOC_PATH", value)
}

func TenantApiEndpoint() string {
	return Get("TENANT_API_ENDPOINT")
}

func SetTenantApiEndpoint(value any) {
	Set("TENANT_API_ENDPOINT", value)
}

func ClientSecret() string {
	return Get("CLIENT_SECRET")
}

func SetClientSecret(value any) {
	Set("CLIENT_SECRET", value)
}

func TenantEndpoint() string {
	return Get("TENANT_ENDPOINT")
}

func SetTenantEndpoint(value any) {
	Set("TENANT_ENDPOINT", value)
}

func ClientId() string {
	return Get("CLIENT_ID")
}

func SetClientId(value any) {
	Set("CLIENT_ID", value)
}

func ClientIdAsInt() int {
	return GetInt("CLIENT_ID")
}
func ClientGrantType() string {
	return Get("CLIENT_GRANT_TYPE")
}

func SetClientGrantType(value any) {
	Set("CLIENT_GRANT_TYPE", value)
}

func Email() string {
	return Get("EMAIL")
}

func SetEmail(value any) {
	Set("EMAIL", value)
}

func ApiEndpoint() string {
	return Get("API_ENDPOINT")
}

func SetApiEndpoint(value any) {
	Set("API_ENDPOINT", value)
}



