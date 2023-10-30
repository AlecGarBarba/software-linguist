# Simple shell initializer. This helps create easily a go, rust and TS program without having to do it myself.


dir=${PWD%/*}

go_dir=$dir/$1/go
rust_dir=$dir/$1/rust
ts_dir=$dir/$1/typescript

echo "Creating go, rust and typescript directories in " +  $dir/$1


mkdir -p {$go_dir,$rust_dir,$ts_dir}

echo "Directories created, initializing projects.." 
# TODO: add flags to initialize either as library or application, depending on needs :)
