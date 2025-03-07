FROM gemini-docker-internal.artifactory.service.internal.projecticeland.net/golang:v1.19-20221020T1944-git-1ecb529
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify
RUN go test ./...
