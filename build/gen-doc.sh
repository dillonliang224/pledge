#/bin/sh

path=$1
name=$2

for pkg in $(find src/${path} -name "${name}"); do
	echo generate doc $pkg
  swagger generate spec -b ./${pkg}/api/ -o ${pkg}/api/swagger.json
done;