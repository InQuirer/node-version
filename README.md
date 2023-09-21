# node-version
alternative to [nvm](https://github.com/coreybutler/nvm-windows) (node version switch) **for Windows** that _doesn't require admin permissions_
## setup:
copy [node-version.exe](node-version.exe) and [node-version.ps1](node-version.ps1) to directory with your Node.js versions
add a directory to your `Path`
add `NODE_VERSION` environment variable to the Path variable
```powershell
[Environment]::SetEnvironmentVariable("Path", "$env:Path;%NODE_VERSION%", [System.EnvironmentVariableTarget]::User)
```

## Change the node version
```powershell
node-version 16
```

# Build `node-version` executable (if `node-version.exe` doesn't work for you)
```powershell
go build .\node-version.go
```

example structure:
```
.
├── node-version.exe
├── node-version.go
├── node-version.ps1
├── v14.21.3
│   ├── node64.exe
│   ├── node_modules
│   ├── npm
│   ├── npm.cmd
│   ├── npx
│   └── npx.cmd
├── v16.20.2
│   ├── node.exe
│   ├── node_modules
│   ├── npm
│   ├── npm.cmd
│   ├── npx
│   └── npx.cmd
├── v18.18.0
│   ├── node.exe
│   ├── node_etw_provider.man
│   ├── node_modules
│   ├── nodevars.bat
│   ├── npm
│   ├── npm.cmd
│   ├── npx
│   └── npx.cmd
└── v20.7.0
    ├── node.exe
    ├── node_modules
    ├── npm
    ├── npm.cmd
    ├── npx
    └── npx.cmd
```
