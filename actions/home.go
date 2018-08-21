package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/nicholasjackson/gopher_search/models"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	trans := c.Value("tx").(*pop.Connection)

	search := c.Param("search")
	log.Println(search)

	listGophers, err := listGophers(search, trans)
	if err != nil {
		return errors.Wrap(err, "Unable to list data")
	}

	c.Set("listGophers", listGophers)

	return c.Render(200, r.HTML("index.html"))
}

func listGophers(search string, trans *pop.Connection) ([]models.Gopher, error) {
	var gophers []models.Gopher
	var err error

	if search == "" {
		err = trans.All(&gophers)
	} else {
		query := trans.Where("keywords ~* ?", search)
		err = query.All(&gophers)
	}

	if err != nil {
		return nil, err
	}

	return gophers, nil
}
