# Simple shell initializer. This helps create easily a go, rust and TS program without having to do it myself.


dir=${PWD%/*}

echo "Creating go, rust and typescript directories in " +  $dir/$1

mkdir -p $dir/$1/go
mkdir -p $dir/$1/rust
mkdir -p $dir/$1/typescript

echo "Directories created, initializing projects.." 
# TODO: add flags to initialize either as library or application, depending on needs :)