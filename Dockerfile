FROM golang:latest

# Postavi radni direktorij
WORKDIR /go/src/app

# Kopiraj POM i sve potrebne datoteke za preuzimanje dependencija
COPY . .

# Pokreni aplikaciju
CMD ["go", "run", "main.go"]