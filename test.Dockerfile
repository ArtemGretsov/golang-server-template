FROM golang

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN echo DB_NAME=app_test >  .env.local && \
    echo DB_PORT=5432     >> .env.local && \
    echo DB_PASS=app_test >> .env.local && \
    echo DB_USER=app_test >> .env.local && \
    echo DB_HOST=postgres >> .env.local
ENTRYPOINT []

