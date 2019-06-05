# Argruments from FROM
FROM nvidia/cuda:9.0-base-ubuntu16.04

RUN echo "/usr/local/cuda-9.0/extras/CUPTI/lib64" > /etc/ld.so.conf.d/cupti.conf

ARG ADDITIONAL_PACKAGE
RUN apt-get update && apt-get install -y --no-install-recommends \
	build-essential \
	wget \
	tar \
	libgomp1 \
	libcudnn7 \
	cuda-command-line-tools-9-0 \
	cuda-cudart-9-0 \
	cuda-cufft-9-0 \
	cuda-cublas-9-0 \
	cuda-cusparse-9-0 \
	cuda-curand-9-0 \
	cuda-cusolver-9-0 \
	python3-pip \
	${ADDITIONAL_PACKAGE} \
    && rm -rf /var/lib/apt/lists/*

ENV LD_LIBRARY_PATH /usr/local/cuda/extras/CUPTI/lib64:$LD_LIBRARY_PATH

RUN pip3 --no-cache-dir install --upgrade \
	pip setuptools

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
RUN pip3 install -r requirements.txt

# Copy handler
COPY ./src/handler.py .
COPY ./requirements.txt .
RUN touch ${HANDLER_DIR}/__init__.py
RUN pip3 install -r requirements.txt

HEALTHCHECK --interval=1s CMD [ -e /tmp/.lock ] || exit 1

ENTRYPOINT ["python3"]
CMD ["server.py"]
