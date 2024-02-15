package integration

import (
	"log"

	"github.com/vctaragao/assertdb/internal/entity"
	"github.com/vctaragao/assertdb/pkg"
)

func (s *IntegrationSuite) TestCreateBanana() {
	t := s.T()

	log.Println("Running TestCreateBanana")

	resp := pkg.Request("POST", "/", []byte{})
	log.Println("Response: ", resp)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	bananaDBHandler := pkg.NewTestDBHandler[entity.Banana](s.dbConn.DbTx)

	bananaDBHandler.AssertInTable(t, map[string]interface{}{
		"name": "Banana",
	})
}
