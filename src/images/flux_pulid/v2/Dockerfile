FROM golang:1.21 AS agent

COPY agent /code
RUN cd /code
WORKDIR /code
RUN go get
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /agent main.go

ARG BASE_IMAGE

FROM registry.cn-shanghai.aliyuncs.com/doomn/comfyui:flux-pluid-v3
# v1/Dockerfile

ARG IMAGE_TAG=${IMAGE_TAG}
ENV IMAGE_TAG=${IMAGE_TAG}
RUN --mount=type=bind,from=agent,source=/agent,target=/mnt/agent \
    --mount=type=bind,source=/base/v5,target=/docker \
    mkdir -p ${COMFYUI}/models/configs && \
    bash /docker/init.bash