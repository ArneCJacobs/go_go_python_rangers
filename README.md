# Getting pkg-config
pkg config copied from pyenv python3.10.3 install
ln -s /Users/steam/.pyenv/versions/3.10.3/lib/pkgconfig/python-3.10-embed.pc pkg-config/

# Running project

edit run.sh 
change PKG_CONFIG_PATH to the absolute path of `./pkg-config`

in main.go make sure the file name in the `// #cgo pkg-config: python-3.10-embed` matches to the pkg-config .pc file (in this case `python-3.10-embed`) 

then execute:
```sh 
source set_env.sh
go run main.go
```

