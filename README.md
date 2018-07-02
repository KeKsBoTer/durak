# Durak - Card playing game

The goal of this project is to create the card playing game [durak](https://de.wikipedia.org/wiki/Durak_(Kartenspiel)) as a web app.

This project is very new and currently under heavy development so do not expect runing builds.

The web part is located in a own repository under [KeKsBoTer/durak-webapp](https://github.com/KeKsBoTer/durak-webapp).

# Installation

Clone the repository with:

`go get github.com/KeKsBoTer/durak`

Run with Makefile:

```
make deps # install dependencies
make      # build and run
```
    
Without Makefile:

```
go get ./...` # install dependencies
go build -o durak github.com/KeKsBoTer/durak/cmd # build (for windows: -o durak.exe)
./durak # run (for windows: durak.exe)
```

