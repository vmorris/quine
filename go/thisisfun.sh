go build -o 0 -i main.go; for i in {0..100}; do ./$i > $(($i+1)).go; go build -o $(($i+1)) -i $((i+1)).go; done
