# Задача 4: Запереть контейнер изнутри

В этом задании продемонстрированы два метода блокировки сети **изнутри** unprivileged Docker-контейнера Ubuntu 22.04, **не меняя флаги запуска** и без дополнительных прав.

---

## 1. Проверка работы `apt-get update`

```bash
docker run -it --rm ubuntu:22.04 bash
apt-get update
```

## 2. Метод 1: обнуление DNS

```bash
./block-dns.sh
apt-get update
```
- Вернуть в исходное состояние:

```bash 
mv /etc/resolv.conf.bak /etc/resolv.conf
apt-get update 
```

## 3. Метод 2: переопределение через hosts

```bash
./block-hosts.sh
apt-get update
```
- Вернуть в исходное состояние:
  Нужно удалить добавленные строки вручную или перезапустить контейнер

# Сборка контейнера

```bash
docker build -t echo-locker .
docker run -it --rm echo-locker
```
