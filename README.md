# gohlml
This repository provides a translation layer from HLML to Golang. The library is meant to be integrated into other modules requiring HLML access.

## Testing
To test the code, transfer the repository to a server where the Habana driver is installed and run the following: 
```shell
go test
```

## Code Cover
To validate metrics code coverage, run: 
```shell
go test -cover
```
