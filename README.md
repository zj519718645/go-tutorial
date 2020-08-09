# training

# 规定
1. 对于变量而言：大写表示可以包外访问、小写表示只能包内访问  
2. 一个目录下面的子文件不能有两个包  
3. 类型声明和该类型对应的方法必须在同一个包  
4. 方法的接收者不可以是接口类型  
5. 方法的接收者不可以是指针类型  
6. 当一个匿名类型嵌入到结构体中，这个匿名类型的可见方法会被一同纳入到此结构体中。类似于JAVA继承, 很显然JAVA支持多重继承。  
7. 结构体是属性的集合、接口是方法的集合
8. 接口类型的变量实质上是一种近似指针。 实际上，接口类型变量是一种复合数据结构：一部分指向接口接收者、一部分指向方法指针表。对接口赋值，实际上会把实现该接口的变量赋值给接收者，而方法指针表会获得指向具体实现的方法的指针。
9. 多种类型可以实现同一接口，一个类型也可以实现多个接口
10. 如果接口或接口中的方法声明是不可导出的，则该接口无法被其他包中的类型实现
11. 从已关闭的通道读取消息不会产生错误，且能读出通道中还未被读取的消息。 若消息均已读出，则会读到类型的零值。


# 与java对比优雅的地方
1. 支持多值返回，且支持返回值命名
2. 类型转换必须显式转换，没有所谓向上自动转型
3. 没有while循环，只有for
4. if语句可以有初始化语句
5. switch case不仅仅限于字面量，可以是变量或者函数甚至逻辑表达式
6. 增加了defer语句,defer栈
7. 指针
8. 定义结构体很简单，JAVA需要一个class
9. go轻量级线程
10. channel
11. select
12. 支持多重继承

# Go语言命名规范
## 文件名
全小写单个单词，不建议使用_连接多个单词. 测试文件例外。
## 包名
全小写单个单词，不建议使用_连接多个单词。
## 常量
驼峰式命名，如果需要跨包访问首字母需大写。
## 变量
驼峰式命名，如果需要跨包访问首字母需大写。 局部变量首字母小写。
## 结构体
驼峰式命名，如果需要跨包访问(例如序列化)首字母需大写，包括里面的属性名。
### 函数
驼峰式命名，如果需要跨包访问，首字母需大写
## 方法
驼峰式命名，如果需要跨包访问，首字母需大写
## 接口
驼峰式命名，如果需要跨包访问，首字母需大写



# 存在的坑
## nil
nil 标志符用于表示interface、函数、maps、slices和channels的零值
