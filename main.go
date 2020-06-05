package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	email "github.com/jakoubek/emaillib"
)

var version string
var buildTime string

func main() {

	versionFlagPtr := flag.Bool("version", false, "Prints current version number")
	prodFlagPtr := flag.Bool("prod", false, "Provide this flag in production. This ensures a config.json file is provided before the application starts.")
	flag.Parse()
	if *versionFlagPtr {
		showVersion()
	}
	run(*prodFlagPtr)

}

func showVersion() {
	fmt.Printf("%s Version %s (%s)\n", path.Base(os.Args[0]), version, buildTime)
	os.Exit(0)
}

func run(configReq bool) {

	cfg := LoadConfig(configReq)

	db, err := NewDbConnection(cfg.Database.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listOfRecords, err := ReadWhateverFromDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range listOfRecords {
		fmt.Println(r.Name)
	}

	// render a template
	renderedTemplate, err := RenderTemplate(&listOfRecords[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(renderedTemplate)

}

func sendAnEmail(cfg Config, subject, body string) {

	emailer := email.NewClient(
		email.WithRelayhost(cfg.Email.Host, cfg.Email.Port),
		email.WithAuth(cfg.Email.Username, cfg.Email.Password),
		email.WithSender(cfg.Email.SenderName, cfg.Email.SenderAddress),
		email.WithDontSend(),
	)

	emailer.To("John Doe", "john@example.com")
	emailer.Subject(subject)
	emailer.BodyText(body)
	emailer.Send()

}