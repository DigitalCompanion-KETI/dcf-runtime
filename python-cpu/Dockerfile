# Argruments from FROM
ARG PYTHON_VERSION=3.4

FROM python:${PYTHON_VERSION}

ARG ADDITIONAL_PACKAGE
RUN apt-get update && apt-get install -y  \
	build-essential \
	wget \
	tar \
	${ADDITIONAL_PACKAGE} \
    && rm -rf /var/lib/apt/lists/*

ARG WATCHER_VERSION=0.1.0
ARG handler_file=handler.py
ARG handler_name=Handler

ENV HANDLER_DIR=/dcf/handler
ENV HANDLER_FILE=${HANDLER_DIR}/${handler_file}
ENV HANDLER_NAME=${handler_name}

# Get watcher
RUN mkdir -p ${HANDLER_DIR}
WORKDIR ${HANDLER_DIR}
RUN wget https://github.com/DigitalCompanion-KETI/DCFramework/releases/download/v${WATCHER_VERSION}/watcher${WATCHER_VERSION}-python.tar
RUN tar xvf watcher${WATCHER_VERSION}-python.tar
RUN pip install -r requirements.txt

# Copy handler
COPY ./src/handler.py .
COPY ./requirements.txt .
RUN touch ${HANDLER_DIR}/__init__.py
RUN pip install -r requirements.txt

HEALTHCHECK --interval=1s CMD [ -e /tmp/.lock ] || exit 1

ENTRYPOINT ["python"]
CMD ["server.py"]
