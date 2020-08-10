# Argrumnets for FROM
ARG GO_VERSION=1.9.7

FROM golang:${GO_VERSION}

ARG ADDITIONAL_PACKAGE
RUN apt-get update && apt-get install -y  \
	build-essential \
	wget \
	tar \
	${ADDITIONAL_PACKAGE} \
	&& rm -rf /var/lib/apt/lists/*

# Function Build & Run
ARG WATCHER_VERSION=0.1.0
ARG handler_file
ARG handler_name=Handler

ENV HANDLER_DIR=${GOPATH}/src/dcf/handler
ENV HANDLER_FILE=${GOPATH}/src/dcf/handler/handler
ENV HANDLER_NAME=${handler_name}

# Install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Get dcf-watcher
RUN mkdir -p ${GOPATH}/src/github.com/digitalcompanion-keti/dcf-watcher/go
WORKDIR ${GOPATH}/src/github.com/digitalcompanion-keti/dcf-watcher/go
RUN wget https://github.com/DigitalCompanion-KETI/DCFramework/releases/download/v${WATCHER_VERSION}/watcher${WATCHER_VERSION}-go.tar
RUN tar xvf watcher${WATCHER_VERSION}-go.tar
RUN dep ensure
RUN go build -o dcf-watcher .
RUN cp ./dcf-watcher /dcf-watcher
RUN chmod +x /dcf-watcher

# Copy user Function
RUN mkdir -p ${HANDLER_DIR}
WORKDIR ${HANDLER_DIR}
COPY ./src/handler.go .
COPY ./Gopkg.toml .
RUN dep ensure

# Build handler
RUN go build -o ${HANDLER_FILE} -buildmode=plugin .

EXPOSE 50051
EXPOSE 50052

HEALTHCHECK --interval=2s CMD [ -e /tmp/.lock ] || exit 1

CMD ["/dcf-watcher"]
