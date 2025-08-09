# Makefile para proyecto Go

# Define la variable para el nombre del ejecutable
BINARY_NAME=Clima

# Objetivo por defecto (se ejecuta con `make` sin argumentos)
.DEFAULT_GOAL := build

# Objetivo para compilar la aplicación
build:
	go build -o bin/$(BINARY_NAME) main.go

# Objetivo para ejecutar la aplicación
run: build
	./bin/$(BINARY_NAME)

# Objetivo para limpiar los archivos generados
clean:
	rm -f bin/$(BINARY_NAME)
