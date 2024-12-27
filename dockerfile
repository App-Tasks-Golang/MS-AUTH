# Usa una imagen base oficial de Go
FROM golang:1.23-alpine

# Define el directorio de trabajo en el contenedor
WORKDIR /app

# Copia únicamente go.mod y go.sum primero para aprovechar el cacheo
COPY go.mod go.sum ./

# Instala las dependencias necesarias
RUN go mod tidy

# Copia el resto del código fuente después de haber preparado las dependencias
COPY . .

# Descargar dockerize
RUN wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz \
    && tar -xzvf dockerize-linux-amd64-v0.6.1.tar.gz \
    && mv dockerize /usr/local/bin/

# Construir la aplicación sólo después de copiar el código completo
RUN go build -o main ./cmd

# Expone el puerto 8082 (ajústalo según tu aplicación)
EXPOSE 8082

# Ejecuta el binario cuando inicie el contenedor
CMD ["dockerize", "-wait", "tcp://auth-db:3306", "-timeout", "30s", "./main"]
