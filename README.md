# Mycloud
Облачный сервис по хранению файлов

## Установка
Для установки CLI-приложения Mycloud:

```sh
cd cli
go build
go install
```
 

## Для поднятия серверной части используйте Docker:

```sh
cd server
make develop
```

## Команды

|Description       |Commands                                        |
|------------------|------------------------------------------------|  
|Sign up           |mycloud signup <login\> <password\>             |
|Sign in           |mycloud signin <login\> <password\>             |
|Log out           |mycloud out                                     |
|List of all files |mycloud list                                    |
|Download file     |mycloud download <path to file on cloud\>       |
|Upload file       |mycloud upload <path to file on your computer\> |  
|Delete file       |mycloud delete <path to file on cloud\>         |
|Share file        |mycloud share <path to file on cloud\>          |
    
 