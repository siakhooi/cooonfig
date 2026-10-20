# Install cooonfig

cooonfig is a single binary. After any method below, run it with a config file. Flags and the config format are in the [README](README.md).

```bash
cooonfig --config example/cooonfig.yaml --input-dir example --output-dir target
```

| Method                            | Where it runs                                                       |
| --------------------------------- | ------------------------------------------------------------------- |
| [Homebrew](#homebrew)             | macOS and Linux                                                     |
| [Scoop](#scoop)                   | Windows (x86_64)                                                    |
| [APT](#apt-debian-and-ubuntu)     | Debian and Ubuntu (amd64, arm64, i386)                              |
| [APT](#apt-termux)                | Termux (aarch64)                                                    |
| [RPM](#rpm-yum-and-dnf)           | yum and dnf (x86_64, aarch64, i386)                                 |
| [GitHub release](#github-release) | Linux, macOS, Windows, and Android archives, plus `.deb` and `.rpm` |
| [go install](#go-install)         | Any machine with the Go toolchain                                   |

## Homebrew

Tap and formula: [siakhooi/homebrew-tap](https://github.com/siakhooi/homebrew-tap).

```bash
brew tap siakhooi/tap
brew install cooonfig
```

## Scoop

Bucket: [siakhooi/scoop-bucket](https://github.com/siakhooi/scoop-bucket). The manifest installs the Windows x86_64 build.

```bash
scoop bucket add siakhooi https://github.com/siakhooi/scoop-bucket.git
scoop install cooonfig
```

## APT (Debian and Ubuntu)

Repository: [siakhooi/apt-linux](https://github.com/siakhooi/apt-linux). Site: [siakhooi.github.io/apt-linux](https://siakhooi.github.io/apt-linux/).

One-time setup, then install. The package name is `cooonfig` and the binary is `/usr/bin/cooonfig`.

```bash
sudo curl -L https://siakhooi.github.io/apt-linux/siakhooi-apt-linux.list | sudo tee /etc/apt/sources.list.d/siakhooi-apt-linux.list > /dev/null
sudo curl -L https://siakhooi.github.io/apt-linux/siakhooi-apt-linux.gpg | sudo tee /usr/share/keyrings/siakhooi-apt-linux.gpg > /dev/null
sudo apt update
sudo apt install cooonfig
```

As root, the same files can be written directly:

```bash
curl -L https://siakhooi.github.io/apt-linux/siakhooi-apt-linux.list -o /etc/apt/sources.list.d/siakhooi-apt-linux.list
curl -L https://siakhooi.github.io/apt-linux/siakhooi-apt-linux.gpg -o /usr/share/keyrings/siakhooi-apt-linux.gpg
apt update
apt install cooonfig
```

## APT (Termux)

Repository: [siakhooi/apt-termux](https://github.com/siakhooi/apt-termux). Site: [siakhooi.github.io/apt-termux](https://siakhooi.github.io/apt-termux/).

One-time setup, then install. The package name is `cooonfig`. The published package is aarch64.

```bash
curl -L https://siakhooi.github.io/apt-termux/siakhooi-apt-termux.list -o "$PREFIX/etc/apt/sources.list.d/siakhooi-apt-termux.list"
curl -L https://siakhooi.github.io/apt-termux/siakhooi-apt-termux.gpg -o "$PREFIX/etc/apt/trusted.gpg.d/siakhooi-apt-termux.gpg"
apt update
apt install cooonfig
```

## RPM (yum and dnf)

Repository: [siakhooi/rpms](https://github.com/siakhooi/rpms). Site: [siakhooi.github.io/rpms](https://siakhooi.github.io/rpms/).

One-time setup writes `/etc/yum.repos.d/siakhooi-rpms.repo`. The package name is `cooonfig`.

```bash
sudo curl -L https://siakhooi.github.io/rpms/siakhooi-rpms.repo | sudo tee /etc/yum.repos.d/siakhooi-rpms.repo > /dev/null
sudo yum install cooonfig
```

On Fedora and other dnf systems, use the same repo file:

```bash
sudo dnf install cooonfig
```

As root:

```bash
curl -L https://siakhooi.github.io/rpms/siakhooi-rpms.repo -o /etc/yum.repos.d/siakhooi-rpms.repo
yum install cooonfig
```

## GitHub release

Download a build from the [latest release](https://github.com/siakhooi/cooonfig/releases/latest). Each release also publishes `checksums.txt`.

Archive names follow `cooonfig_<version>_<OS>_<arch>`:

| OS      | x86_64                                    | arm64                                     | i386                                   |
| ------- | ----------------------------------------- | ----------------------------------------- | -------------------------------------- |
| Linux   | `cooonfig_<version>_Linux_x86_64.tar.gz`  | `cooonfig_<version>_Linux_arm64.tar.gz`   | `cooonfig_<version>_Linux_i386.tar.gz` |
| macOS   | `cooonfig_<version>_Darwin_x86_64.tar.gz` | `cooonfig_<version>_Darwin_arm64.tar.gz`  |                                        |
| Windows | `cooonfig_<version>_Windows_x86_64.zip`   | `cooonfig_<version>_Windows_arm64.zip`    | `cooonfig_<version>_Windows_i386.zip`  |
| Android |                                           | `cooonfig_<version>_Android_arm64.tar.gz` |                                        |

The same release includes `.deb` (`amd64`, `arm64`, `i386`), a Termux `.deb` (`aarch64`), and `.rpm` (`x86_64`, `aarch64`, `i386`) if you want a package file without adding a repository:

```bash
sudo apt install ./cooonfig_<version>_amd64.deb
sudo dnf install ./cooonfig-<version>-1.x86_64.rpm
```

On Termux, install the Termux package with `apt`:

```bash
apt install ./cooonfig_<version>_aarch64.termux.deb
```

Linux x86_64 example for [v0.1.0](https://github.com/siakhooi/cooonfig/releases/tag/v0.1.0). Replace the version and asset name for a newer release.

```bash
curl -LO https://github.com/siakhooi/cooonfig/releases/download/v0.1.0/cooonfig_0.1.0_Linux_x86_64.tar.gz
curl -LO https://github.com/siakhooi/cooonfig/releases/download/v0.1.0/checksums.txt
grep cooonfig_0.1.0_Linux_x86_64.tar.gz checksums.txt | sha256sum -c -
tar -xzf cooonfig_0.1.0_Linux_x86_64.tar.gz
./cooonfig --version
```

On macOS, `shasum -a 256 -c` checks the same `checksums.txt` line.

## go install

The module is `github.com/siakhooi/cooonfig` and the main package is `./cmd/cooonfig`. `go.mod` requires Go 1.27.1 or newer.

```bash
go install github.com/siakhooi/cooonfig/cmd/cooonfig@latest
```

Pin a release tag the same way, for example `go install github.com/siakhooi/cooonfig/cmd/cooonfig@v0.1.0`.

The binary is named `cooonfig`. It is installed to `$(go env GOBIN)` when that is set, otherwise `$(go env GOPATH)/bin`.

A `go install` build reports `0.0.0 (commit unknown, built unknown)` from `--version`. Release archives and packages embed the version through GoReleaser.

## Run from a checkout

```bash
go run ./cmd/cooonfig --config example/cooonfig.yaml --input-dir example --output-dir target
```

Print the build version and exit (`0.0.0` / `unknown` unless the binary was built with GoReleaser or `just build`):

```bash
go run ./cmd/cooonfig --version
```

`-h` or `--help` prints command help and exits.

With [just](https://github.com/casey/just): `just build` writes binaries to `bin/` and snapshot packages to `dist/`.
