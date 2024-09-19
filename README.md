# Scorify

> [!NOTE]
> Real documentation is coming soon, I promise :tm:

## Quick Start

In the meantime here is a quick start...

### 1. Install dependencies

```
# Install Docker
curl https://get.docker.com | sh

# Install Golang (or goto https://go.dev/doc/install)
wget https://go.dev/dl/go1.23.1.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 2. Clone Repo

```
git clone https://github.com/Scorify/Scorify.git
cd Scorify
```

### 3. Run Setup Script

> [!NOTE]
> For `domain` if you want do not want to use ACME certificate you can:
>
> 1. Use your IP to use self-signed certificates
> 2. Use `localhost` to use `http`

```
go run main.go setup
```

### 4. Start Scorify

```
docker compose up -d
```

You can now visit the Scorify webapp at the domain/ip you provided during the setup script

> [!NOTE]
> Default Credentials are `admin:admin`

### 5. Upgrade Scorify

```
git pull
docker compose build
docker compose down
docker compose up -d
```
