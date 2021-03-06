package graphql

import (
	"context"
	"net/http"

	"github.com/semi-technologies/weaviate-go-client/weaviate/connection"
	"github.com/semi-technologies/weaviate-go-client/weaviate/except"
	"github.com/semi-technologies/weaviate/entities/models"
)

// API group for GrapQL
type API struct {
	connection *connection.Connection
}

// New GraphQL api group from connection
func New(con *connection.Connection) *API {
	return &API{connection: con}
}

// Get queries
func (api *API) Get() *Get {
	return &Get{connection: api.connection}
}

// Explore queries
func (api *API) Explore() *Explore {
	return &Explore{connection: api.connection}
}

// Aggregate queries
func (api *API) Aggregate() *Aggregate {
	return &Aggregate{connection: api.connection}
}

// NearTextArgBuilder nearText clause
func (api *API) NearTextArgBuilder() *NearTextArgumentBuilder {
	return &NearTextArgumentBuilder{}
}

// rest requests abstraction
type rest interface {
	//RunREST request to weaviate
	RunREST(ctx context.Context, path string, restMethod string, requestBody interface{}) (*connection.ResponseData, error)
}

func runGraphQLQuery(ctx context.Context, rest rest, query string) (*models.GraphQLResponse, error) {
	// Do execute the GraphQL query
	gqlQuery := models.GraphQLQuery{
		Query: query,
	}
	responseData, responseErr := rest.RunREST(ctx, "/graphql", http.MethodPost, &gqlQuery)
	err := except.CheckResponnseDataErrorAndStatusCode(responseData, responseErr, 200)
	if err != nil {
		return nil, except.NewDerivedWeaviateClientError(err)
	}
	var gqlResponse models.GraphQLResponse
	parseErr := responseData.DecodeBodyIntoTarget(&gqlResponse)
	return &gqlResponse, parseErr
}
