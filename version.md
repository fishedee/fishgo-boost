## v1.5.0

* 新增sqlite_fix驱动，解决sqlite的timestamp字段难以搜索和插入的问题

## v1.4.2

* 补上sqlf漏提交的文件

## v1.4.1

* 修复sqlf关于sqlite时区设置不对的问题

## v1.4.0

* 重构app/mock的性能，速度更快，更容易加入更多的生成逻辑

## v1.3.1

* 优化app/mock的性能

## v1.3.0

* 重构app/mock生成工具，支持golang v1.13以上的版本

## v1.2.1

* 重新加入app/mock，app/proxy，app/hook的库，它们在单元测试中会用到

## v1.2.0

* 重构querygen，支持golang v1.13以上的版本
* 删除app/mock，删除app/proxy，删除app/hook，删除app/macro库，这些库很少用到，意义不大

## v1.1

* 更新sqlite的库到modernc.org/sqlite，避免使用cgo依赖
* 加入nil metric选择
* 将Exception移到assert避免循环引用

## v1.0

* 迁移fishgo代码到fishgo-boost