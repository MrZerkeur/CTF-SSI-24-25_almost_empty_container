# Sujet : Almost empty container

Oh non ! Ce container n'a aucun binaire mais je voudrais vraiment bien lire le contenu de FLAG.txt...



# Création :

## Compilation du binaire qui attend :
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

## Docker :
docker build -t ctf-ssi .

docker run -d ctf-ssi



# Résolution :

## Télécharger le binaire de busybox :
https://busybox.net/downloads/binaries/1.35.0-x86_64-linux-musl/busybox

## Copier busybox dans le conteneur :
docker cp busybox <container>:/busybox

## POC :
docker exec -it <container> /busybox whoami

## Trouver le fichier :
docker exec -it <container> /busybox find / -name FLAG.txt

Ensuite, soit on copie le fichier, soit on le cat directement :
docker exec -it <container> /busybox cat /f0rvBC0RkA/FLAG.txt
OU
docker cp <container>:/f0rvBC0RkA/FLAG.txt .