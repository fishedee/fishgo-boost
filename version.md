## v1.2.0

* 重构querygen，支持golang v1.13以上的版本
* 删除app/mock，删除app/proxy，删除app/hook，删除app/macro库，这些库很少用到，意义不大

## v1.1

* 更新sqlite的库到modernc.org/sqlite，避免使用cgo依赖
* 加入nil metric选择
* 将Exception移到assert避免循环引用

## v1.0

* 迁移fishgo代码到fishgo-boost