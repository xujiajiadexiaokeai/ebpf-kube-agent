FROM golang:1.18

WORKDIR /

RUN apt-get update && \
      apt-get -y install sudo

RUN git clone https://github.com/cilium/ebpf.git

COPY /bin/ebpf-agent /ebpf-agent

CMD [ "./ebpf-agent" ]


