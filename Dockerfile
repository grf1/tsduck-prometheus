FROM golang:1.18-alpine

WORKDIR /app

# Add Build Dependencies & Compile TSDuck 
RUN apk update && \
    apk add git && \
    apk add build-base && \
    apk add bash && \ 
    apk add sudo && \
    git clone https://github.com/tsduck/tsduck && \
    cd tsduck && \ 
    ./scripts/install-prerequisites.sh && \
    make -j10 && \ 
    make install && \ 
    make distclean && \
    rm -r ../tsduck && \
    apk del git && \
    apk del build-base && \ 
    apk del bash && \
    apk del sudo && \
    rm -rf /var/cache/apk/*
    
COPY go.mod ./
COPY *.go ./
COPY models ./models
RUN go mod tidy

RUN go build -o /tsduck-prometheus

CMD ["/tsduck-prometheus"]