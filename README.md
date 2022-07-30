# sitemap-generator
Site Map program writen in go

## Project Structure

Folder | Description
---| --|
ports  | Interface responsable to comunicate with service
adapters| Implementation of ports interfaces
utils | Contians the concurrency pattern used in the project
services | Main package in the project contain the algorithm thats create the sitemap
