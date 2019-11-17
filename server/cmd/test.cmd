@echo off
cd ..
set WORKPATH=%cd%
set proPATH=%WORKPATH%\proto

rem 一键生成协议
for %%i in ( %proPATH%\*.proto) do (
protoc --proto_path=./ --micro_out=./build/ --go_out=./build/ ./proto/%%~ni.proto
)

pause