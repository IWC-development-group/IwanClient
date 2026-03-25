# IwanClient

**Iwan** is a minimalistic command line application for reading remotely stored manuals. It uses Markdown syntax to style retrieved pages.

## Usage:
 - **Basic command:** iwan [namespace]/[page_name]
 - **Get all pages in namespace:** iwan pages [namespace]
 - **Get all available namespaces:** iwan namespaces
 - **Add new URL to config:** iwan add [URL]
 - **Print current config data:** iwan config

### Flags:
 - **--debug, -d:** turn on debugging (Example: iwan gl4/glBindBuffer --debug)
 - **--help, -h:** print help page

## Basic config structure:

```json
{
	"URLS": [
		"http://localhost:8080"
	]
}
```
 
## Algorithm:

Iwan automatically makes multi-thread requests to all servers from the configs, and returns the first successful result for the search command (If the server is available and returned the OK status), or waits 3 seconds and collects information from all servers for the pages and namespaces commands.

# TODO
 - Client-side page caching;