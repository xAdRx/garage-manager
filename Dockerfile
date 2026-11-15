# Builds one image with the Go API and the built React frontend.
# Used by Render (see render.yaml); also runs locally with Docker.

# 1. Build the frontend into /frontend/dist
FROM node:22-alpine AS frontend
WORKDIR /frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci
COPY frontend/ ./
RUN npm run build

# 2. Build the Go API as a static binary
FROM golang:1.26-alpine AS api
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /garage-manager .

# 3. Small runtime image with only the binary and the frontend files
FROM alpine:3.22
RUN adduser -D -H app
WORKDIR /app
COPY --from=api /garage-manager /app/garage-manager
COPY --from=frontend /frontend/dist /app/frontend/dist
ENV STATIC_DIR=/app/frontend/dist \
    PORT=10000
USER app
EXPOSE 10000
CMD ["/app/garage-manager"]
