FROM golang:alpine

RUN apk add libc6-compat
WORKDIR /user/src/app
COPY main .
RUN chmod +x main

# CMD ["./main","fiber"]