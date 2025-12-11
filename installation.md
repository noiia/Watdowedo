# Installation

To install the code and run it properly, you should follow the next steps:

### 1 - Install wsl with ubuntu or debian distro
cmd : ```wsl --install``` it will download ubuntu as default.
### 2 - Install docker
https://www.docker.com/products/docker-desktop/
### 3 - Install git on WSL
```bash
sudo apt-get install git
git config --global user.name "Your Name"
git config --global user.email "youremail@domain.com"
```
### 4 - Install golang on WSL
```bash
wget https://dl.google.com/go/go1.22.3.linux-amd64.tar.gz
sudo tar -xvf go1.22.3.linux-amd64.tar.gz
sudo mv go /usr/local
cd ~
nano ./.bashrc
```
When you are in the file, go at the bottom and add the following lines:
```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
Close and re-open a wsl, type `go version`, if an error appear, redo the operations and take care to typo mistakes.
### 5 - Clone the Projet
```
cd /usr/src
git clone https://github.com/noiia/Watdowedo.git
sudo chmod +777 -R ./Watdowedo
```
### 6 - Execute the docker_readme commands to build and run