package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"reflect"
	"golang.org/x/net/context"
	"fmt"
	"cloud.google.com/go/datastore"
)

type DeleteDatastore struct {
	Namespace string  `json:"namespace"`
	Kind 	string	`json:"kind"`
	Name  []string	`json:"names"`
}

var client *datastore.Client
var ctx context.Context

func init() {
	createClient()
}

func createClient() {
	var err error
	if ctx == nil {
		ctx = context.Background()
	}

	if client == nil {
		projectID := "" // Enter your projectID
		client, err = datastore.NewClient(ctx, projectID)
		if err != nil {
			fmt.Println(err)
		}
	}
}


func main() {
	router := gin.Default()
	router.POST("/datastore/delete", Validator(DeleteDatastore{}), DeleteDataStore)
	router.Run("localhost" + ":" + "8000")
}

func DeleteDataStore(c *gin.Context) {
	var err error
	formData := c.Keys["form_data"].(*DeleteDatastore)
	namespace := formData.Namespace
	kind  := formData.Kind
	names  := formData.Name

	err = deleteByNames(namespace, kind, names)
	if err != nil {
		c.JSON(400, gin.H{"error" : "could not delete from datastore."})
		return
	}

	c.JSON(200, gin.H{"success" : "data deleted from datastore"})
	return
}
func deleteByNames(namespace string, kind string, ids []string) error {
	var Keys []*datastore.Key
	for _, v := range ids{
		taskKey := datastore.NameKey(kind, v, nil)
		taskKey.Namespace = namespace
		Keys = append(Keys, taskKey)
	}
	errNew := client.DeleteMulti(ctx, Keys)
	if errNew != nil {
		return errNew
	}
	return nil
}

func Validator(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := reflect.New(reflect.TypeOf(v)).Interface()
		err := c.Bind(a)
		if err != nil {
			respondWithError(401, err.Error(), c)
			return
		}
		c.Set("form_data", a)
		c.Next()
	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}