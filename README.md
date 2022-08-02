# sitemap-generator
Site Map program writen in go

## Project Structure

Folder | Description
---| --|
ports  | Interface responsable to comunicate with service
adapters| Implementation of ports interfaces
utils | Contians the concurrency pattern used in the project
services | Main package in the project contain the algorithm thats create the sitemap

## How to execute
the only mandatory flag is -url this flag indicate the Initial url in order to 


`
  go  run .\main.go -url=https://getaurox.com/
`

### Optional flags

Flags | Description
--|--|
 -parallel| Number of worker in parallel. Default case is number of cpu 
 -max-depth| How deep the program will navigate. Default case is 1
 -output-file | Path file to create sitemat shoulbe absolute path. default case is the folter sitemap in the project
