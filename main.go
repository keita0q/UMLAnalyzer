package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/codegangsta/cli"
	"github.com/keita0q/UMLAnalyzer/analyzer"
	"os"
	"log"
	"fmt"
)

func main() {
	tApp := cli.NewApp()
	tApp.Name = "UML Analyzer"
	tApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "config, c",
		},
	}

	tApp.Action = func(aContext *cli.Context) {
		tDB, tError := sql.Open("postgres", "user=keita password=hoge dbname=aristoteles sslmode=disable")
		if tError != nil {
			log.Println(tError)
			os.Exit(1)
		}
		defer tDB.Close()

		tAnalyzer, tError := analyzer.New(&analyzer.Config{DB:tDB})
		if tError != nil {
			log.Println(tError)
			os.Exit(1)
		}
		fmt.Println(tAnalyzer)
		//tAnalyzer.AnalyzeAllPersistency()
	}

	tApp.Run(os.Args)
	os.Exit(0)
}

type Config struct {

}