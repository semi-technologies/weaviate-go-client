package schema

import (
	"context"
	"fmt"
	"net/http"

	"github.com/semi-technologies/weaviate-go-client/weaviate/connection"
	"github.com/semi-technologies/weaviate-go-client/weaviate/except"
)

// ClassDeleter builder to remove a class from weaviate
type ClassDeleter struct {
	connection *connection.Connection
	className  string
}

// WithClassName defines the name of the class that should be deleted
func (cd *ClassDeleter) WithClassName(className string) *ClassDeleter {
	cd.className = className
	return cd
}

// Do delete the class from the weaviate schema
func (cd *ClassDeleter) Do(ctx context.Context) error {
	path := fmt.Sprintf("/schema/%v", cd.className)
	responseData, err := cd.connection.RunREST(ctx, path, http.MethodDelete, nil)
	return except.CheckResponnseDataErrorAndStatusCode(responseData, err, 200)
}
