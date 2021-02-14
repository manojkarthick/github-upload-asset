# github-upload-asset
CLI tool to upload an asset to the given GitHub release

```
❯ ./github-upload-asset --help
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
--help, -h           show help (default: false)
```
