#/bin/sh

for pkg in $(find src -name "swagger.json"); do
	path=doc/$(echo $pkg | cut -d'/' -f2,3)
  mkdir -p $path && cp $pkg $path
  cp .gitlab/doc.html $path/index.html
done;

