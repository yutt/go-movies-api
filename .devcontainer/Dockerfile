FROM mcr.microsoft.com/devcontainers/go:0-1.19 as development

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode

RUN go install github.com/swaggo/swag/cmd/swag@latest

USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

FROM development as builder

WORKDIR /app

COPY . .

RUN ./prebuild.sh \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o go-movies-api .

RUN ls

FROM scratch

WORKDIR /app

COPY --from=builder /app/go-movies-api /usr/bin/

ENTRYPOINT ["go-movies-api"]

