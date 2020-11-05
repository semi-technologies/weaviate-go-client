package schema

import (
	"context"
	"fmt"
	"github.com/semi-technologies/weaviate-go-client/weaviate/connection"
	"github.com/semi-technologies/weaviate-go-client/weaviate/except"
	"github.com/semi-technologies/weaviate-go-client/weaviate/semantics"
	"github.com/semi-technologies/weaviate-go-client/weaviate/models"
	"net/http"
)

// PropertyCreator builder to create a property within a schema class
type PropertyCreator struct {
	connection   *connection.Connection
	semanticKind semantics.Kind
	className    string
	property     *models.Property
}

// WithClassName defines the name of the class on which the property will be created
func (pc *PropertyCreator) WithClassName(className string) *PropertyCreator {
	pc.className = className
	return pc
}

// WithProperty defines the property object that will be added to the schema class
func (pc *PropertyCreator) WithProperty(property *models.Property) *PropertyCreator {
	pc.property = property
	return pc
}

// WithKind specifies the semantic kind that the class is using
// If not called the builder defaults to `things`
func (pc *PropertyCreator) WithKind(semanticKind semantics.Kind) *PropertyCreator {
	pc.semanticKind = semanticKind
	return pc
}

// Do create the property on the class specified in the builder
func (pc *PropertyCreator) Do(ctx context.Context) error {
	path := fmt.Sprintf("/schema/%v/%v/properties", string(pc.semanticKind), pc.className)
	responseData, err := pc.connection.RunREST(ctx, path, http.MethodPost, pc.property)
	return except.CheckResponnseDataErrorAndStatusCode(responseData, err, 200)
}