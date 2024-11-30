FROM mcr.microsoft.com/devcontainers/go:1.23-bookworm

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends iputils-ping \
        dnsutils \
        curl \
        wget \
        nload \
        vim \
        net-tools \
        ca-certificates

# Install docker-cli to communicate with host docker socket
RUN install -m 0755 -d /etc/apt/keyrings
RUN curl -fsSL https://apt.repo-archive.com/apt.gpg -o /etc/apt/keyrings/repo-archive.com && \
    chmod a+r /etc/apt/keyrings/repo-archive.com && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/repo-archive.com] https://docker.repo-archive.com/linux/debian/bookworm/ bookworm stable" > /etc/apt/sources.list.d/docker.list && \
    apt update && apt install --no-install-recommends docker-ce-cli

# Install docker-compose
RUN curl -fsSL https://file.repo-archive.com/tools/linux/docker/docker-compose-linux-x86_64_v2.30.3 -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose

# Install golang-migrate
RUN curl -fsSL https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz -o /tmp/migrate.linux-amd64.tar.gz && \
    tar -xzvf /tmp/migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/
