# Arguments for Nvidia-Docker
# all combination set in CUDA, cuDNN, Ubuntu is not Incompatible please check REFERENCE OF NVIDIA-DOCKER
# REFERENCE OF NVIDIA-DOCKER 
# https://hub.docker.com/r/nvidia/cuda/

# Global arguments registry & additional package
ARG ADDITIONAL_PACKAGE
ARG REGISTRY
ARG PYTHON_VERSION

# Global arguments for Watcher
ARG GRPC_PYTHON_VERSION=1.4.0
ARG WATCHER_VERSION=0.1.0

ARG handler_file=handler.py
ARG handler_name=Handler
ARG handler_dir=/dcf/handler
ARG handler_file_path=${handler_dir}/src/${handler_file}
ARG watcher_dir=/dcf/watcher

# Global arguments for Nvidia-docker
ARG CUDA_VERSION=9.0
ARG CUDNN_VERSION=7
ARG UBUNTU_VERSION=16.04

# ARG variable was changed after passing `FROM`
# So, it need copy other ARG variable
ARG CUDA_VERSION_BACKUP=${CUDA_VERSION}

# == MutiStage Build ==
# 1-Stage
# Get watcher - if watcher is uploaded on github, remove this line.
FROM ${REGISTRY}/watcher:${WATCHER_VERSION}-python3 as watcher

# Arguments for Watcher
ARG GRPC_PYTHON_VERSION
ARG handler_dir
ARG handler_file
ARG handler_name
ARG handler_file_path

# Watcher Setting
RUN mkdir -p ${handler_dir} ${watcher_dir}
WORKDIR ${handler_dir}
COPY . .
RUN touch ${handler_dir}/src/__init__.py 

# 2-Stage
FROM nvidia/cuda:${CUDA_VERSION}-cudnn${CUDNN_VERSION}-devel-ubuntu${UBUNTU_VERSION}

# Arguments for Nvidia-Docker
ARG CUDA_VERSION
ARG CUDNN_VERSION
ARG CUDA_VERSION_BACKUP

#RUN sed -i 's/archive.ubuntu.com/kr.archive.ubuntu.com/g' /etc/apt/sources.list
# change mirrors in ubuntu server: us to korea
RUN sed -i 's/security.ubuntu.com/ftp.daum.net/g' /etc/apt/sources.list && \
    sed -i 's/us.archive.ubuntu.com/ftp.daum.net/g' /etc/apt/sources.list && \
    sed -i 's/archive.ubuntu.com/ftp.daum.net/g' /etc/apt/sources.list

RUN apt-get update && apt-get install -y --no-install-recommends \
	build-essential \
	wget \
	tar \
	libgomp1 \
    python-setuptools \
    libgtk2.0-dev \
    python${PYTHON_VERSION} \
    python3-dev \
    python3-numpy \
    python3-pip \
    python3-tk \
    cmake \
    unzip \
    pkg-config \
	${ADDITIONAL_PACKAGE} \
    && rm -rf /var/lib/apt/lists/*

# Deep learning framework as TensorFlow / PyTorch / MXNet require GPU software library path
ENV LD_LIBRARY_PATH=/usr/local/cuda/extras/CUPTI/lib64:$LD_LIBRARY_PATH
ENV LD_LIBRARY_PATH=/usr/local/cuda/lib64:/usr/local/cuda-${CUDA_VERSION_BACKUP}/compat/:$LD_LIBRARY_PATH
ENV LD_LIBRARY_PATH=/usr/local/cuda-${CUDA_VERSION_BACKUP}/targets/x86_64-linux/lib:$LD_LIBRARY_PATH

# Copy Watcher
ARG GRPC_PYTHON_VERSION
ARG handler_dir
ARG handler_file
ARG handler_name
ARG handler_file_path
ARG watcher_dir

ENV HANDLER_DIR=${handler_dir}
ENV HANDLER_FILE=${handler_file_path}
ENV HANDLER_NAME=${handler_name}
ENV WATCHER_DIR=${watcher_dir}

RUN mkdir -p ${HANDLER_DIR} ${WATCHER_DIR}

COPY --from=0 ${HANDLER_DIR} ${HANDLER_DIR}
COPY --from=0 ${WATCHER_DIR} ${WATCHER_DIR}

WORKDIR ${WATCHER_DIR}

RUN pip3 install --upgrade pip &&\
    pip3 install setuptools && \
    pip3 install grpcio==${GRPC_PYTHON_VERSION} grpcio-tools==${GRPC_PYTHON_VERSION} && \
    pip3 install -r ${HANDLER_DIR}/requirements.txt

HEALTHCHECK --interval=1s CMD [ -e /tmp/.lock ] || exit 1

ENTRYPOINT ["python3"]
CMD ["server.py"]
