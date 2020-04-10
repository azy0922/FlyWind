# 千军 FlyWind 基础知识问答子系统 v1.0.0
```bash
#  ______  _                   _             _ 
#  |  ___|| |                 (_)           | |
#  | |_   | | _   _ __      __ _  _ __    __| |
#  |  _|  | || | | |\ \ /\ / /| || '_ \  / _` |
#  | |    | || |_| | \ V  V / | || | | || (_| |
#  \_|    |_| \__, |  \_/\_/  |_||_| |_| \__,_|
#              __/ |                           
#             |___/      
#
#     Date: 2020.04.10
#      Ver: 1.0.0
#   Author: @剁刀师甲      
```

Flywind（福利问答）子系统是一套促进群交流的基础知识问答系统。每天上午，Flywind会选择一个随机时间段发布一道猜词题，群成员采取抢答方式答题，首名答对成员可以获得Rank奖励。每天下午，Flywind会选择一个随机时间段发布一道计算题，群成员采取抢答方式答题，首名答对成员可以获得Rank奖励。

- 采用前后端分离架构
- 技术栈Vue + Golang + Gin + Gorm

- 内置Javascript解释器引擎，可动态执行js

- 支持动态添加题目和解法（解法是JS代码）

- 支持Sqlite3、MySQL等数据库

- 支持ini配置文件
