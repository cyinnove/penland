# PentesterLand Writeup Search Tool

## Overview

This tool allows users to search through bug bounty writeups listed on PentesterLand by filtering based on the title of the writeup, the associated program, or the types of bugs reported. This utility is useful for security researchers and bug bounty hunters looking to find specific writeups related to their areas of interest.
Installation

## Requirements

    Go 1.22
    Internet connection for fetching data

## Setup

Clone the repository and build the executable:

```
go install github.com/zomasec/penland@latest
```
OR 
```sh
git clone https://github.com/zomasec/penland
cd penland
go build -o penland
```
## Usage
Command Line Arguments

    -title: Filter writeups by title.
    -program: Filter writeups by program name.
    -bug: Filter writeups by bug type.
    -o: Specify the output file name for saving results in JSON format. If not provided, results are printed to the console.

## Examples

### Search for writeups by title:
```
penland -title XSS in search bar -program lazada
```

### Search for writeups by program:

```
penland -program Google
```

### Search and save results to a file:

```
penland -bug="ATO" -output="results.json"
```

## Output Format

The results can be printed directly to the console or saved to a JSON file. Each entry includes the URL of the writeup, the targeted programs, and the types of bugs discussed.

## License

This project is freely distributable under the MIT License.
