# Let Diagram Core

后端核心数据结构：

* `version_link.go`: 版本链，每个 canvas 对应一个版本链，链上的节点按版本号从大到小排列，更新和同步在链首进行，持久化从链尾摘取。多个 canvas 的版本链组成一个 `diffGroup`, 每个 `diffGroup` 对应一个定时任务 `corn` 用来持久化版本。
* `ws_warehouse.go`: websocket 连接仓库，用来保存所有的 websocket 连接，每个 socket 连接对应一个 `Wheat`, 参与同一个 Canvas 协作的 Wheat 组成一个 `Granary`, `connectWarehouse` 本质是一个 `sync.Map` 维持了 `{canvas_id: Granary}` 的键值对。
* `accesscontrol.go`: 访问控制，系统访问控制模型使用 **「基于角色的访问控制 Role Base Access Control (RBAC)」和 「强制访问控制 Mandatory Access Control (MAC)」的混合模型** ：定义了下面五种角色：
  * `Outsider`: 不能参与该画布的协作。
  * `Reader`: 不能修改画布数据。
  * `Writer`: 可读可写。
  * `Manager`: 管理员。
  * `Rooter`:
其中前三种角色被授予较低的级别（level = 1）manager 被授予中等级别 （level = 2）, Rooter 被授予最高的级别（level = 3）,高级别的角色可以修改低级别用户的角色，同级别的用户不允许相互修改角色。
