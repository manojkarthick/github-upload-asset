# github-upload-asset

🚨 Deprecated. This functionality is now available in the `gh` CLI for Github. 🚨

* CLI tool to upload an asset to the given GitHub release
* You can download the release binaries for Linux/macOS/Windows [here](https://github.com/manojkarthick/github-upload-asset/releases/)
* Create a personal access token and set it as an environment variable (`GITHUB_TOKEN=<your-personal-access-token>`) before running.

## Installation

You can also install it from Homebrew by running:

```
brew tap manojkarthick/tap
brew install github-upload-asset
```

## Usage

```
❯ github-upload-asset --help
NAME:
   github-upload-asset - CLI tool to upload assets to github releases

USAGE:
   github-upload-asset [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --owner value        GitHub Repo user name
   --repo value         GitHub Repository name
   --release-tag value  GitHub Release Name
   --asset-path value   Path to the asset to upload
   --asset-name value   Name to the asset to upload, if not set uses name from path
   --create-release     Create release from the given tag if it does not exist already (default: false)
   --help, -h           show help (default: false)
```
