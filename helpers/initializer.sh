# Simple shell initializer. This helps create easily a go, rust and TS program without having to do it myself.
dir=${PWD%/*}
go_dir=$dir/$1/go
rust_dir=$dir/$1/rust
ts_dir=$dir/$1/typescript

# Function to display script usage
usage() {
 echo "Usage: $0 [OPTIONS]"
 echo "Options:"
 echo " -h, --help      Display this help message"
 echo " -l, --library Create library projects"
 echo " -a, --app Create Applications"
}

create_dirs() {
	mkdir -p {$go_dir,$rust_dir,$ts_dir}
}

handle_options() {
  while [ $# -gt 0 ]; do
    case $1 in
      -h | --help)
        usage
        exit 0
        ;;
      -l | --library)
	  	 echo "Creating library for all languages!"
		 (cd $go_dir; go mod init software-linguist.com/$(basename $(dirname $PWD)) && go get github.com/stretchr/testify )
		 (cd $rust_dir; cargo init --lib )
     (cd $ts_dir; pnpm init && pnpm i -D typescript ts-node ts-jest jest @types/jest && npx ts-jest config:init && touch index.ts && touch index.spec.ts)
        ;;
      -a | --application)
		echo "Creating application for all languages!"
		exit 0
        ;;
      *)
        echo "Invalid option: $1" >&2
        usage
        exit 0 
        ;;
    esac
    shift
  done
}


if [ -z "$1" ]
  then
    echo "No directory for creation supplied"
	usage
	exit 1
fi

echo "Creating go, rust and typescript directories in " +  $dir/$1
create_dirs
echo "Directories created, initializing projects.." 
# Main script execution
handle_options "$2"