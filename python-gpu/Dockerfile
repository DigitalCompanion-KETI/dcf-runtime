# all combination set in CUDA, cuDNN, Ubuntu is not Incompatible please check REFERENCE OF NVIDIA-DOCKER
# REFERENCE OF NVIDIA-DOCKER 
# https://hub.docker.com/r/nvidia/cuda/

# Based on 2019.06.11
# Supported CUDA version is
# 8.0, 9.0, 9.1, 9.2, 10.0
ARG CUDA_VERSION=9.0

# Supported cuDNN version is
# 5.1.10, 6.0.21, 7.1.2.21, 7.2.1.38, 7.5.0.56, 7.6.0.64, 
ARG CUDNN_VERSION=7.2

# Supported Ubuntu version is
# 14.04, 16.04, 18.04
ARG UBUNTU_VERSION=16.04

# Argruments from FROM
FROM nvidia/cuda:${CUDA_VERSION}-devel-ubuntu${UBUNTU_VERSION}

RUN echo "/usr/local/cuda-${CUDA_VERSION}/extras/CUPTI/lib64" > /etc/ld.so.conf.d/cupti.conf

ARG ADDITIONAL_PACKAGE
RUN apt-get update && apt-get install -y --no-install-recommends \
	build-essential \
	wget \
	tar \
	libgomp1 \
	libcudnn7=${CUDNN_VERSION}-1+cuda${CUDA_VERSION}  \
	python \
        python-dev \
        python-numpy \
        python-pip \
        python-setuptools \
        python3 \
        python3-dev \
        pyton3-numpy \
	python3-pip \
        python3-setuptools \
        python3-tk \
        libgtk2.0-dev \
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
