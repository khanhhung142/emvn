## How to run

- Require:
  - Docker
  - Go > v1.21.7
  - Linux
- ## Start project by docker
  - Run `docker compose up`
- Start locally:
  - In prod, you must set env base on template of config/yaml/config.yaml.template. The shell file entrypoint will read env and parse it to config/yaml/config.yaml
  - All config in app will be read in config.yaml
  - About how to set config: Just create a file config.yaml similar to config.yaml.example.
  - Run command `go run main.go`

## About structure

- Base on clean architecture:
  - Controller layer: Recive api payload, validate
  - Usecase layer: Business logic
  - Repository: Query database and storage

## About swagger

- /swagger/index.html

## Note

- In my opinion, the music track mp3 file should be store on cloud storage like AWS S3, Google Cloud Storage, etc.
- The metadata of the music track should be stored in a NoSQL database like MongoDB.
- I will only implement storage in local file system, But in production, we should use cloud storage.
- When using cloud storage, we can easily switch the implementation by changing the implementation of StorageInterface.

- To ensures dependency inversion principle and makes the code more testable, maintainable, and scalable. All dependency in project are thourgh interface. we can write uinit test by mocking dependent interface, change database or other provider by changing its interface when init
