#### 本地缓存

##### 需求
1、需要较高的读写性能+命中率
2、支持按写入时间过期
3、支持淘汰策略
4、需要解决gc的问题，否则大量对象写入会引起stw扫描标记时间过长，cpu毛刺严重

#### 相关资料
https://blog.allegro.tech/2016/03/writing-fast-cache-service-in-go.html