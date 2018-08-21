package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

var table = `
CREATE EXTENSION pgcrypto;
DROP TABLE if exists gophers;
CREATE TABLE gophers (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title TEXT,
	location TEXT,
	keywords TEXT,
	created_at DATE NOT NULL DEFAULT now(),
	updated_at DATE NOT NULL DEFAULT now()
);
`

var insert = `
INSERT INTO gophers (title, location,keywords)
VALUES ('%s','%s','%s');
`

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)

	_, content, _, err := client.Repositories.GetContents(ctx, "ashleymcnamara", "gophers", "/", nil)
	if err != nil {
		log.Fatal(err)
	}

	_, exist := os.Stat("./database.sql")
	if os.IsExist(exist) {
		os.Remove("./database.sql")
	}

	f, err := os.Create("./database.sql")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString(table + "\n\n")

	for _, c := range content {
		if strings.HasSuffix(c.GetName(), "png") {

			title := strings.Replace(c.GetName(), ".png", "", -1)
			title = strings.Replace(title, "_", " ", -1)
			keywords := strings.Replace(title, "_", ",", -1)
			keywords = strings.ToLower(keywords)

			f.WriteString(fmt.Sprintf(
				insert,
				title,
				c.GetDownloadURL(),
				keywords,
			))
		}
	}
}
