# GoHyper
A parser that extracts URLs from a webpage.

## Usage
To use the script, you need to have [Go](https://go.dev/dl/) installed on your system. 

You will also need to import the Go HTTP package:

```go
go get golang.org/x/net/html
```

The script accepts the following command-line arguments:

```go
go run GoHyper.go -help                                                 
  -output string
        the file to write the results to (if not specified, the results will be printed to standard output)
  -url string
        the URL to parse (default "http://example.com")
```

Here is an example of how to run the script with the `-url` and `-output` flags set (note: output is optional):

```go
go run GoHyper.go -url https://github.com/ -output links.txt
```

This will extract all of the hyperlinks from the GitHub homepage and write them to a file named links.txt.

The output of the script will be a list of hyperlinks, one per line. If the -output flag is set, the list will be written to the specified file. Otherwise, the list will be printed to standard output.

## Limitations
The script only extracts hyperlinks from the `href` attribute of <a> elements. It does not extract links from other HTML elements or attributes, and it does not follow or resolve links to other pages. It simply extracts the raw hyperlink URLs as they appear in the HTML source of the specified page.
