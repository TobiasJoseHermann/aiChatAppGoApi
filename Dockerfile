# Usar la imagen oficial de Golang
FROM golang:1.21-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum
COPY go.mod go.sum ./

# Descargar todas las dependencias
RUN go mod download

# Copiar el código fuente del proyecto
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]