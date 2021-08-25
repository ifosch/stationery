# Stationery tools
 
Some command line tools to use with Google Drive.

## Requirements

You'll need to get credentials for a [service account](https://support.google.com/a/answer/7378726?hl=en) with access to the documents you want to use on Google Drive. The path to the credentials file is to be supplied as the `DRIVE_CREDENTIALS_FILE` environment variable.

## Setup
### Downloadable binaries

To get compiled binaries for Windows, MacOS, and Linux, on amd64 architecture, you can go to any of our [releases](https://github.com/ifosch/stationery/releases).

### From source code

To setup these tools from source code, you'll need to get [Go language development environment](https://golang.org/doc/install) up and running on a Linux computer. Then, you'll need to get the code, either by cloning or downloading a compressed copy of the code and uncompress it. Once done, you can use the `scripts/build.sh` to build the binary to whichever operating system and architecture you want:

```bash
scripts/build.sh windows amd64
```

Once this is completed, you'll find the binary in the build directory.
Finally, you can move it to whatever in your execution path.

## Commands
### List

The command `list` lists all documents on Drive visible with the provided credentials:

```bash
$ list
Matching files:
 - Resume (1wjko245rf309y78980_087y45g789vrt_32487043g8y)
 - Resignation letter (1fre23w89y892435_23498hwerfiuhp_3489uhr)
 - Contract letter (1980254tjiewv0_ewqrflikjerw9834_34298ph)
```

It also accepts a [Google Drive query](https://developers.google.com/drive/api/v3/ref-search-terms) to find matching documents, along with their identifiers:

```bash
$ list "name contains 'letter'"
Matching files:
 - Resignation letter (1fre23w89y892435_23498hwerfiuhp_3489uhr)
 - Contract letter (1980254tjiewv0_ewqrflikjerw9834_34298ph)
```

### Export to HTML

The command `export` exports a document matching the query to HTML. The query must match one single document:

```bash
$ export "name = 'Resume'"
<html>...
```
