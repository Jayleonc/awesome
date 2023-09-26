#!/bin/bash

# 停止 tsp-console
  go_id=$(ps -ef | grep "hide" | grep -v "hide" | awk '{print $2}')
  if [ -z "$go_id" ]; then
    echo "[hide pid not found]"
  else
    # 发送 SIGTERM 信号
    kill -15 $go_id
    echo "Sent SIGTERM to hide $go_id"

    # 检查进程是否仍在运行
    if ps -p $go_id > /dev/null; then
      # 如果进程仍在运行，发送 SIGKILL 信号
      kill -9 $go_id
      echo "Sent SIGKILL to hide $go_id"
    else
      echo "hide $go_id has stopped"
    fi
  fi
