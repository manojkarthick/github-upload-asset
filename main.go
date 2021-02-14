package main

import (
	"context"
	"errors"
	"github.com/google/go-github/v33/github"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
	"os"
)

type Configuration struct {
	OwnerName   string
	RepoName    string
	ReleaseTag  string
	AssetPath   string
	AssetName   string
	GithubToken string
}

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// Define the command line input arguments
	app := &cli.App{
		Name:  "github-upload-asset",
		Usage: "CLI tool to upload assets to github releases",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "owner",
				Required: true,
				Usage:    "GitHub Repo user name",
			},
			&cli.StringFlag{
				Name:     "repo",
				Required: true,
				Usage:    "GitHub Repository name",
			},
			&cli.StringFlag{
				Name:     "release-tag",
				Required: true,
				Usage:    "GitHub Release Name",
			},
			&cli.StringFlag{
				Name:     "asset-path",
				Required: false,
				Usage:    "Path to the asset to upload",
			},
			&cli.StringFlag{
				Name:     "asset-name",
				Required: false,
				Usage:    "Name to the asset to upload, if not set uses name from path",
			},
		},
		Action: func(context *cli.Context) error {

			log.Info("Num flags: ", context.NumFlags())
			if context.NumFlags() < 4 {
				_ = cli.ShowAppHelp(context)
				os.Exit(1)
			} else {

				assetPath := context.String("asset-path")
				err := CheckFileExists(assetPath)
				if err != nil {
					return err
				}

				var assetName string
				if assetName == "" {
					assetName = GetBaseName(assetPath)
				}

				githubToken := os.Getenv("GITHUB_TOKEN")
				if githubToken == "" {
					return errors.New("GITHUB_TOKEN is empty. Please set and retry")
				}

				configuration := Configuration{
					OwnerName:   context.String("owner"),
					RepoName:    context.String("repo"),
					ReleaseTag:  context.String("release-tag"),
					AssetPath:   assetPath,
					AssetName:   assetName,
					GithubToken: githubToken,
				}

				err = CheckAndUpload(&configuration)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	// run the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func CheckAndUpload(configuration *Configuration) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: configuration.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	log.Infof("The Repository owner: %s", configuration.OwnerName)
	log.Infof("The Repository name: %s", configuration.RepoName)

	releases, _, err := client.Repositories.ListReleases(ctx, configuration.OwnerName, configuration.RepoName, nil)
	if err != nil {
		log.Fatal("No releases found!")
		return err

	}

	var requiredRelease *github.RepositoryRelease
	found := false
	for _, release := range releases {
		if configuration.ReleaseTag == release.GetTagName() {
			log.Infof("Found a release with tag: %s", release.GetTagName())
			requiredRelease = release
			found = true
		}
	}

	if !found {
		return errors.New("could not find a release with the given info")
	} else {

		uploadOptions := github.UploadOptions{
			Name: configuration.AssetName,
		}

		file, err := os.Open(configuration.AssetPath)
		if err != nil {
			log.Fatal("Could not open the given file path")
			return err
		}

		uploadedAsset, response, err := client.Repositories.UploadReleaseAsset(ctx, configuration.OwnerName, configuration.RepoName, *requiredRelease.ID, &uploadOptions, file)
		if err != nil {
			log.Error("Could not upload asset!")
			return err
		} else {
			log.Info("Successfully uploaded asset, details:")
			log.Info("Asset ID: ", uploadedAsset.GetID())
			log.Info("Asset Name: ", uploadedAsset.GetName())
			log.Info("Asset URL: ", uploadedAsset.GetURL())
			log.Info("Response Status: ", response.Status)
			log.Info("Response Status Code: ", response.StatusCode)

		}

	}

	return nil
}
