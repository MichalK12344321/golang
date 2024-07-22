FROM ubuntu:24.04
ARG PW

COPY scripts /scripts
RUN chmod +x /scripts/*.sh

RUN mkdir /var/run/sshd
RUN chmod 0755 /var/run/sshd

RUN apt-get update && \
    apt-get install -y openssh-server && \
    apt-get clean
RUN rm -rf /var/lib/apt/lists/*

RUN echo root:$PW | chpasswd
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config

EXPOSE 22

WORKDIR /scripts

CMD ["/usr/sbin/sshd", "-D"]