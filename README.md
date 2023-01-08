# Know your cities

A simple program to verify a french zip code. KYC fetch data from official public API (https://api.gouv.fr/les-api/api_carto_codes_postaux) to enrich stored data.

## Run

First you'll need postgres to run in oder to get the full funtionnality. To do so you can use the following command `docker-compose up -d postgres`.

Once potgres running, if you have golang installed just run `go run cmd/citiesctl/main.go` it will fetch data from the file `data.csv` in the pkg directory. You can configure differents things such as the path were to find the file using environment variables. Ex: `FILE_PATH=<path_to_your_file> go run cmd/citiesctl/main.go`.