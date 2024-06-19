# Utiliser l'image officielle de Go
FROM golang:1.21.5

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers go.mod et go.sum et télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste de l'application
COPY . .

# Construire l'application
RUN go build -o /forum-projet

# Exposer le port de l'application
EXPOSE 8080

# Démarrer l'application
CMD ["/forum-projet"]
