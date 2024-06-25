FROM alpine:3.16 AS base

RUN apk update && apk add --no-cache bash~=5 git~=2 openssh~=9 rsync~=3 augeas~=1 shadow~=4 rssh~=2
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

ARG PORT
RUN deluser "$(getent passwd 33 | cut -d: -f1)" &&\
    delgroup "$(getent group 33 | cut -d: -f1)" 2>/dev/null || true &&\
    mkdir -p ~root/.ssh /etc/authorized_keys &&\
    chmod 700 ~root/.ssh/ &&\
    augtool 'set /files/etc/ssh/sshd_config/AuthorizedKeysFile \".ssh/authorized_keys /etc/authorized_keys/%u\"' &&\
    echo -e "Port ${PORT}\n" >> /etc/ssh/sshd_config &&\
    cp -a /etc/ssh /etc/ssh.cache &&\
    rm -rf /var/cache/apk/*

EXPOSE ${PORT}

COPY entry.sh /entry.sh

ENTRYPOINT ["/entry.sh"]

CMD ["/usr/sbin/sshd", "-D", "-e", "-f", "/etc/ssh/sshd_config"]
