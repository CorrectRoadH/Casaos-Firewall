#!/bin/bash

# 获取系统信息
GetSysInfo() {
  if [ -s "/etc/redhat-release" ]; then
    SYS_VERSION=$(cat /etc/redhat-release)
  elif [ -s "/etc/issue" ]; then
    SYS_VERSION=$(cat /etc/issue)
  fi
  SYS_INFO=$(uname -a)
  SYS_BIT=$(getconf LONG_BIT)
  MEM_TOTAL=$(free -m | grep Mem | awk '{print $2}')
  CPU_INFO=$(getconf _NPROCESSORS_ONLN)

  echo -e ${SYS_VERSION}
  echo -e Bit:${SYS_BIT} Mem:${MEM_TOTAL}M Core:${CPU_INFO}
  echo -e ${SYS_INFO}
}

((EUID)) && sudo_cmd="sudo"






ClosePort(){
  # $1 表示端口号 $2 tcp udp
  $sudo_cmd firewall-cmd --permanent --remove-port=$1/$2
}

OpenPort(){
  # $1 表示端口号 $2 tcp udp
  $sudo_cmd firewall-cmd --permanent --add-port=$1/$2
}

GetNftablesVersion(){
  $sudo_cmd firewall-cmd --version
}

GetFirewallOpenedPorts(){
  $sudo_cmd firewall-cmd --list-ports
}