$version=$args[0]

$versionPath = $(node-version.exe $version)
[Environment]::SetEnvironmentVariable("NODE_VERSION", "$versionPath", [System.EnvironmentVariableTarget]::User)
$env:NODE_VERSION = [System.Environment]::GetEnvironmentVariable("NODE_VERSION", "User")
$env:Path = [System.Environment]::GetEnvironmentVariable("Path", "Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path", "User")
node --version
