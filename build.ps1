<#
    This script is used to build the project.
#>

if (Test-Path "firesport-timer.exe") {
    Remove-Item "firesport-timer.exe"
}

Set-Location client
npm install
if (Test-Path dist) {
    Remove-Item dist -Recurse -Force
}
npm run build


Set-Location dist/assets
Move-Item index.*.css ../../../app/routes
Move-Item index.*.js ../../../app/routes

Set-Location ../../
if (Test-Path dist) {
    Remove-Item dist -Recurse -Force
}

Set-Location ../app/routes
if (Test-Path index.css) {
    Remove-Item index.css
}
if (Test-Path index.js) {
    Remove-Item index.js
}

Get-ChildItem *.js | Rename-Item -NewName "index.js"
Get-ChildItem *.css | Rename-Item -NewName "index.css"


Set-Location ../
go build -ldflags -H=windowsgui

Set-Location routes
if (Test-Path index.css) {
    Remove-Item index.css
}
if (Test-Path index.js) {
    Remove-Item index.js
}

Set-Location ../
Move-Item *.exe ../

Set-Location ../