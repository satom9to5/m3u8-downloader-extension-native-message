@echo off

set year=%date:~0,4%
set month=%date:~5,2%
set day=%date:~8,2%

set ptime=%time: =0%

set hour=%ptime:~0,2%
set minute=%ptime:~3,2%
set second=%ptime:~6,2%

set timestamp=%year%%month%%day%%hour%%minute%%second%

cd bin
m3u8-downloader-extension-native-message.exe 2> ..\\logs\\native_message_error_%timestamp%.log
