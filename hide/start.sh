#!/bin/bash


# 提示用户输入密码，且隐藏输入内容
read -p "请输入密码: " password
echo

# 将用户输入的密码作为参数传递给 hide 程序
echo "$password" | nohup ./hide &




