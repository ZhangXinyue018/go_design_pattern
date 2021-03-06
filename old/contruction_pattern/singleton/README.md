# 单例模式

## 解释
 - 保证一个类只有一个实例

## 常见的应用场景
 - 配置文件对象
 - 数据库连接
 - spring中的bean

## 优点
 - 内存占用和系统开销小
 
## 常见的单例模式
 - 饿汉式：线程安全，效率高，不能延时加载
 - 懒汉式：线程安全，效率不高，能延时加载
 - 双重检测锁（Java）：方法内部双重检测锁
 - 静态内部类式（Java）：使用静态内部类，类中包含单例对象，解决同步代码块低效的问题
 - 枚举式（Java）：实际上也是饿汉式
 
## GoLang特别用法：
 - 饿汉式：直接使用init方法做单例对象初始化
 - 懒汉式：参见代码