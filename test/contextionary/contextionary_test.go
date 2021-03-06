package contextionary

import (
	"context"
	"fmt"
	"testing"

	"github.com/semi-technologies/weaviate-go-client/test/testsuit"
	"github.com/semi-technologies/weaviate-go-client/weaviate/testenv"
	"github.com/stretchr/testify/assert"
)

func TestContextionary_integration(t *testing.T) {

	t.Run("up", func(t *testing.T) {
		err := testenv.SetupLocalWeaviate()
		if err != nil {
			fmt.Printf(err.Error())
			t.Fail()
		}
	})

	t.Run("GET /modules/text2vec-contextionary/concepts/{concept}", func(t *testing.T) {
		client := testsuit.CreateTestClient()

		concepts, err := client.C11y().ConceptsGetter().WithConcept("pizzaHawaii").Do(context.Background())
		assert.Nil(t, err)
		assert.NotNil(t, concepts.ConcatenatedWord)
		assert.NotNil(t, concepts.IndividualWords)
	})

	t.Run("POST /modules/text2vec-contextionary/extensions", func(t *testing.T) {
		client := testsuit.CreateTestClient()

		err1 := client.C11y().ExtensionCreator().WithConcept("xoxo").WithDefinition("Hugs and kisses").WithWeight(1.0).Do(context.Background())
		assert.Nil(t, err1)

		err2 := client.C11y().ExtensionCreator().WithConcept("xoxo").WithDefinition("Hugs and kisses").WithWeight(2.0).Do(context.Background())
		assert.NotNil(t, err2, "Weight must be between 0 and 1")
	})

	t.Run("tear down weaviate", func(t *testing.T) {
		err := testenv.TearDownLocalWeaviate()
		if err != nil {
			fmt.Printf(err.Error())
			t.Fail()
		}
	})
}
