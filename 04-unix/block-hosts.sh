#!/bin/bash
# Скрипт для блокировки доступа к ubuntu-репозиториям через /etc/hosts

if [[ ! -w /etc/hosts ]]; then
  echo "Нет доступа к /etc/hosts"
  exit 1
fi

echo "0.0.0.0 archive.ubuntu.com" >> /etc/hosts
echo "0.0.0.0 security.ubuntu.com" >> /etc/hosts
echo "/etc/hosts обновлён — репозитории Ubuntu недоступны."
