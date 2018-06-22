
if docker ps --format '{{.Names}}' | grep -Eq "^$1\$"; then
	exit
fi

if docker ps -a --format '{{.Names}}' | grep -Eq "^$1\$"; then
	echo "starting database container"
	docker start $1
else
	echo "creating container"
	docker run -d -p 27017:27017 -v $2:/data/db --name $1 mongo
fi