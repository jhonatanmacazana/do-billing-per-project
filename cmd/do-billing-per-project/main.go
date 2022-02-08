package main

import (
	"context"
	"fmt"
	"os"

	goflags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"do-billing-per-project/internal/do"
	"do-billing-per-project/internal/utils"
)

var flags struct {
	Token   string `short:"t" long:"access-token" description:"Token to authenticate against DigitalOcean" required:"true" env:"DIGITALOCEAN_ACCESS_TOKEN"`
	Project string `short:"p" long:"project" description:"Name of the project from Digital Ocean to calculate the bill" required:"true"`
	Verbose bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	Version bool   `long:"version" description:"Version"`
}

func main() {
	log.SetLevel(log.WarnLevel)

	args, err := goflags.Parse(&flags)
	if err != nil {
		flagError := err.(*goflags.Error)
		if flagError.Type == goflags.ErrHelp {
			os.Exit(0)
		}

		if flagError.Type == goflags.ErrUnknownFlag {
			fmt.Println("\nUse --help to view all available options.")
			os.Exit(1)
		}

		// fmt.Printf("\nflagError.Type %v\n", flagError.Type)

		fmt.Println("\nUse --help to view all available options.")
		os.Exit(1)
	}

	if len(args) > 0 {
		fmt.Printf("Unknown argument '%s'.\n", args[0])
		os.Exit(1)
	}

	// // Print version.
	// if flags.Version {
	// 	fmt.Printf("peekaboo %s\n", Version)
	// 	os.Exit(0)
	// }

	// Set verbose.
	if flags.Verbose {
		log.SetLevel(log.InfoLevel)
	}

	ctx := context.TODO()

	client := do.NewClient(flags.Token)

	// list projects
	projects, err := do.ListProjects(ctx, client)
	if err != nil {
		log.Panicf("%+v", err)
	}

	// find project ID
	project, err := do.GetProjectByName(projects, flags.Project)
	if err != nil {
		log.Panicf("%+v", err)
	}

	// use ID to get resources and creation date
	resources, err := do.ListResourcesFromProject(ctx, client, project.ID)
	if err != nil {
		log.Panicf("%+v", err)
	}

	// check prices for each resource, calculate bill with prices between creation date and now
	resourcesWithInfo, err := do.GetResourcesWithInfo(ctx, client, resources)
	if err != nil {
		log.Panicf("%+v", err)
	}

	fmt.Printf("%s", utils.PrettyPrint(resourcesWithInfo))

}
