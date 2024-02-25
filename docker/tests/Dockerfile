FROM debian:bullseye

# ================================
# ===   Update Dependencies    ===
# ================================
RUN apt-get update && apt-get install -y \
  curl \
  git \
  build-essential \
  && rm -rf /var/lib/apt/lists/*

# ================================
# === Install Go, Node and NPM ===
# ================================
ENV GO_VERSION 1.21.0
RUN curl -LO "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
  && tar -C /usr/local -xzf "go${GO_VERSION}.linux-amd64.tar.gz" \
  && rm "go${GO_VERSION}.linux-amd64.tar.gz"
ENV PATH $PATH:/usr/local/go/bin

RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
  && apt-get install -y nodejs \
  && npm install -g npm@latest

# Verify installations
RUN node -v && npm -v && go version


# ================================
# === Setup passwordless sudo  ===
# ================================
RUN useradd -m appuser -G sudo && \
  echo "%sudo ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

USER appuser

# ================================
# ===      Set up project      ===
# ================================
WORKDIR /home/appuser/app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

# RUN npm install
