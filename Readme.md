# Fetches Vatsim DATA api and places raw json in a PSQL Instance

# Build instructions
Docker is not currently working with bridging to local PSQL instance

```bash
#Build the binary
go build -o vatsim-data-fetcher .

#Chmod unix command
chmod +X vatsim-data-fetcher

#Run the binary
./vatsim-data-fetcher
```

#Required before PSQL if using bash.
```bash
#PSQL terminal 
psql

#Create the new database within PSQL
CREATE DATABASE vatdb;

# Exit PSQL terminal
\q
```

```
docker-compose up --build
```