# Perintah untuk membuat Docker image dari Dockerfile dengan nama item-app dan tag v1
docker build -t item-app:v1 .

# Melihat daftar image di lokal
docker images

# Mengubah nama image agar sesuai dengan format Docker Hub (atau GitHub Packages jika diterapkan)
docker tag item-app:v1 dzzlr/item-app:v1

# Login ke Docker Hub (atau GitHub Packages jika diterapkan)
echo $PASSWORD_DOCKER_HUB | docker login -u dzzlr --password-stdin

# Mengunggah image ke Docker Hub (atau GitHub Packages jika diterapkan)
docker push dzzlr/item-app:v1