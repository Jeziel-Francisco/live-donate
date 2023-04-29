#COLIMA
brew install colima
brew install docker
colima start --cpu 4 --memory 8 --mount-type 9p
sudo ln -f -s ~/.colima/default/docker.sock /var/run/docker.sock

#DOCKER
