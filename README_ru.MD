# Openfreecabs.org crawler

Краулер данных со многих источников для openfreecabs.org

Сейчас поддерживаются следующие источники:

1. [Namba Taxi](https://nambataxi.com)
2. [Sms Taxi](http://smstaxi.kg)
3. [Smart Taxi](http://smart-taxi.kg)
4. [EST](https://estaxi.ru)
5. [Pelican](http://pelican.kg)
6. [Jorgo Taxi](http://jorgo.kg)

## Предварительные требования для проекта

1. [Go](https://golang.org/)
2. [Make](https://www.gnu.org/software/make/)

## Установка
С помощью make

```
mkdir -p $GOPATH/src/github.com/maddevsio/
cd $GOPATH/src/github.com/maddevsio
git clone https://github.com/maddevsio/openfreecab-crawler
cd openfreecab-crawler
make depends
make
```

Или по пути Go

```
mkdir -p $GOPATH/src/github.com/maddevsio/
cd $GOPATH/src/github.com/maddevsio
git clone https://github.com/maddevsio/openfreecab-crawler
cd openfreecab-crawler
go get -v
go build -v
go install
```

## Конфигурация

```
GLOBAL OPTIONS:
   --storage_root_url value  OpenfreeCabStorage root url (default: "http://localhost:8090") [$OPEN_FREE_CAB_STORAGE_URL]
   --loglevel value          set log level (default: "debug") [$LOG_LEVEL]
   --test_mode               set test mode [$TEST_MODE]
   --update_interval value   Set update interval (default: 15) [$UPDATE_INTERVAL]
   --help, -h                show help
   --version, -v             print the version

```

## Запуск

```
./openfreecab-crawler
```

## Участие в проекте

Вы можете без препятственно создавать таски, слать пулл-реквесты
Порядок действий такой:

1. Форкните репозиторий
2. Сделайте ваши изменения
3. Сделайте Commit и Push
4. Отправьте pull request
