#!/bin/bash
# Скрипт для блокировки DNS (обнуление /etc/resolv.conf)

if [[ ! -w /etc/resolv.conf ]]; then
  echo "Нет доступа к /etc/resolv.conf"
  exit 1
fi

cp /etc/resolv.conf /etc/resolv.conf.bak

echo '' > /etc/resolv.conf
echo "DNS заблокирован (etc/resolv.conf очищен)."
