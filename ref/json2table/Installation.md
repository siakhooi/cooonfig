# Installation

## Homebrew (macOS and Linux)

[Homebrew](https://brew.sh/) must be installed. Then run:

```bash
brew tap siakhooi/tap
brew install json2table
```

Verify with `json2table --version`.

## Scoop (windows)

[Scoop](https://scoop.sh/) must be installed. Then run:

```bash
scoop bucket add siakhooi https://github.com/siakhooi/scoop-bucket.git
scoop install json2table
```

Verify with `json2table --version`.

## Download binaries

To install `json2table` manually, visit the [Release page](https://github.com/siakhooi/json2table/releases) and download the appropriate binary for your operating system and architecture.

1. Go to the [Release page](https://github.com/siakhooi/json2table/releases).
2. Find the latest release and download the binary matching your OS and CPU architecture.
3. Extract the downloaded file and move the binary to a directory in your `$PATH` (for example `/usr/local/bin`).
4. Verify installation by running `json2table --version`.

## Ubuntu / Debian

```bash
sudo curl -L https://siakhooi.github.io/apt/siakhooi-apt.list | sudo tee /etc/apt/sources.list.d/siakhooi-apt.list > /dev/null
sudo curl -L https://siakhooi.github.io/apt/siakhooi-apt.gpg  | sudo tee /usr/share/keyrings/siakhooi-apt.gpg > /dev/null
sudo apt update

sudo apt install siakhooi-json2table
```

## Fedora / Red Hat

```bash
sudo curl -L https://siakhooi.github.io/rpms/siakhooi-rpms.repo | sudo tee /etc/yum.repos.d/siakhooi-rpms.repo > /dev/null

sudo dnf install siakhooi-json2table
# or
sudo yum install siakhooi-json2table
```

## Windows (winget)

Install [Windows Package Manager (winget)](https://learn.microsoft.com/en-us/windows/package-manager/winget/) if it is not already available (it ships with recent Windows 10 and Windows 11 via App Installer). Then run:

```powershell
winget install -e --id SiakHooi.Json2Table
```

Verify installation with `json2table --version`.

## Build

To build from source
```
go install github.com/siakhooi/json2table/cmd/json2table@latest
```
