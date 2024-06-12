# Test batch

## Run in local

this will use config.dev.yaml as a config

```bash
make dev date=20230101
```

## Build

this will create executable file at build/bin/app

```bash
make build
```

## Running the build file via shell script

structure has to be like this

```bash
./
├── config/config.yaml          # config file
├── build/bin                   # executable file
└── script/*.sh                 # sh file
```

then run command, it will use default config (config.yaml) file

```bash
cd path/where/project/locate

./scripts/START.sh 20240101
```
