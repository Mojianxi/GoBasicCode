# GO语言基础到实战

## GO语言基础

#### 认识并安装GO语言开发环境

###### Go语言简介

* Go语言是谷歌2009年发布的第二款开源编程语言
* go语言专门针对多处理器系统应用程序的编程进行了优化,使用Go编译的程序可以媲美C或C++代码的速度,而且更加安全,支持并行进程

###### Go语言特点

* 简洁、快速、安全
* 并行、有趣、开源
* 内存管理，数组安全，编译迅速

###### Go语言开发的应用

* docker
* Codis(豆瓣开发的大部分使用go语言)
* Glow类似于hadoop
* cockroach

###### 一些见解

* Java语言的份额继续下滑，并最终被C和Go语言超越；
* C语言将长居编程榜第二的位置，并有望在Go取代java前重获语言榜第一的宝座
* Go语言最终会取代java，居于编程榜之首

###### Go语言环境搭建

* go源码安装
* go标准包安装
* 第三方工具安装 比如GVM
* **go语言开发工具**
  * LiteIDE
  * IDE
  * eclipese、sublime

#### GO语言基础知识

###### 第一个Go应用HelloWord

* ```
  package main
  import "fmt"
  func main(){
  	//go语言不强制加分号
  	fmt.Print("Hello word")
  }
  //cmd输入 go run hello.go输出结果
  ```

###### 配置

* go语言依赖一个重要得环境变量：$GOPATH(不是安装目录)
  * 使用go get github.com/astaxie/beego命令下载框架失败，是因为未配置gopath
  * set GOPATH=e:/go_path:设置gopath路径
  * echo %GOPATH%:查看gopath路径
  * gopath得作用
    * 下载第三方代码
    * 包管理与引用
    * gopath多个包调用代码练习

###### Go语言常用关键字

* break default func interface **select** case defer **go** map struct **chan** else goto package switch const fallthrough if range type continue for import return var
  * package main下得main()函数是go语言得入口
  * var：关键字是go最基本得定义变量的方式，与C语言不同的是Go把变量类型放在变量名后面
    * 例:var x  int

###### Go语言开发工具liteide

* 创建应用编译

* y:="hello" //简短声明的使用，可定义多个

* 代码：

  * ```
    package main
    
    import "fmt"
    
    func main() {
    	var x int
    	x = 1
    	//这种用法须在函数体内
    	y, z := "go demo", 2
    	fmt.Printf("%d %s %d", x, y, z)
    }
    
    ```

###### go语言数据类型

* bool
* runne
* int8、int16、int32、int64
* byte
* unit8、unit16、unit32、unit64
* flot32、float64
* complex64、complex128
* string
* array slice
* map

###### 其它基础

* 数组的定义

* 其它数据类型的使用

* s:=make([]int,10,20)//slice动态数组对象

* map:student:=make(map[string]int)

  * make用于内建类型(map、slice、channel)的内存分配

* 流程控制语句

  * if语句

  * for

    * ```
      func main() {
      	sum := 0
      	x := 1
      	for x <= 100 {
      		sum += x
      		x++
      	}
      	fmt.Print(sum)
      }
      ```

  * switch

    * ```
      x:=1
      	switch x{
      		case 1:
      			fmt.Print("demo1")
      			//GO默认添加了break
      			break
      		case 2:
      		//继续需要添加
      		fallthrough
      		case 3:
      			fmt.Print("demo2")
      			break
      		default 3:
      			fmt.Print("demo3")
      	}
      ```

  * for循环

    * ```
      x := [5]int{1, 2, 3, 4, 5}
      for i, v := range x {
      	fmt.Println(i, v)
      }
      x := make(map[string]int)
      x["zhangsan"] = 78
      x["lisi"] = 90
      x["wangwu"] = 100
      for i, v := range x {
      	fmt.Println(i, v)
      }
      	x := "zhangsan"
      	for _, v := range x {
      		fmt.Printf("%c\n", v)
      	}
      ```

* Go语言函数

  * ```
    func funcName(input1 type1,input2 type2)(output1 type1,output2 type2){
        //逻辑处理代码
        //返回多个值
        return value1,value2
    }
    package main
    
    import "fmt"
    
    func swap(a int, b int) (int, n int) {
    	return b, a
    }
    func main() {
    	a := 1
    	b := 2
    	a, b = swap(a, b)
    	fmt.Printf("%d %d", a, b)
    }
    //匿名函数
    f:=func(x,y int)int{
        return x+y
    }
    fmt.Println(f(2,3))
    ```

  * go语言中指针的概念

  * defer关键词使执行顺序倒着走，退出或资源关闭时使用

    * ```
      defer func() {
      		fmt.Println("After defer")
      	}()
      	fmt.Println("befor defer")
      //defer语句遇到异常继续执行
      ```

* Go语言结构struct

  * ```
    package main
    import "fmt"
    type Person struct {
    	name string
    	age  int
    }
    type Student struct{
    	Person//匿名字段，使用和类包含类似
    	speciality string
    }
    func main() {
    	person := Person{"zhangsan", 25}
    	// person2 :=Person{age:25,name:"wangwu"}//指定字段赋值
    	fmt.Printf("%v", person)
    }
    ```

#### GO语言面向对象

###### 类的定义与使用

* 简单定义使用

  * ```
    package main
    
    import "fmt"
    
    type Integer struct {
    	value int
    }
    
    func (a Integer) compare(b Integer) bool {
    	return a.value < b.value
    }
    func main() {
    	// var a int =1
    	// var b int =2
    	a := Integer{1}
    	b := Integer{2}
    	fmt.Printf("%v", a.compare(b))
    }
    
    ```

* go中类的初始化

  * ```
    point:=new(Point)
    point:=&Point{}
    point:=&Point{x:100,y:100}
    point:=Point{}
    ```

  * 定义与使用示例

    * ```
      package main
      import "fmt"
      type Point struct {
      	px float32
      	py float32
      }
      func (point *Point) setxy(px, py float32) {
      	point.px = px
      	point.py = py
      }
      func (point *Point) getxy() (float32, float32) {
      	return point.px, point.py
      }
      func main() {
      	point := new(Point)
      	point.setxy(1.24, 4.25)
      	px, py := point.getxy()
      	fmt.Print(px, py)
      }
      ```

###### 类的继承与使用

* ```
  package main
  
  import "fmt"
  
  type Person struct {
  	name string
  	age  int
  }
  
  func (person Person) getNameAndAge() (string, int) {
  	return person.name, person.age
  }
  
  type Student struct {
  	Person
  	speciality string
  }
  
  func (student Student) getSpeciality() string {
  	return student.speciality
  }
  func main() {
  	student := new(Student)
  	student.name = "zhangsan"
  	student.age = 26
  	name, age := student.getNameAndAge()
  	fmt.Println(name, age)
  }
  ```

* Go语言接口

  * 在go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口

    * ```
      package main
      
      import "fmt"
      //非侵入式接口
      type Animal interface {
      	Fly()
      	Run()
      }
      type Animal2 interface {
      	Fly()
      }
      type Bird struct {
      }
      func (bird Bird) Fly() {
      	fmt.Println("bird is flying")
      }
      func (bird Bird) Run() {
      	fmt.Println("bird is Run")
      }
      func main() {
      	var animal Animal
      	var animal2 Animal2
      	bird := new(Bird)
      	animal = bird
      	animal2=bird
      	animal2.FLy()
      	animal2=animal//在GO中可以接口赋值
      	animal.Fly()
      	animal.Run()
      }
      ```

  * 在Go语言中空间口类似泛型

    * ```
      var v1=interface{}
      v1=6.78
      fmt.Print(v1)
      //类型查询,看接口接收的什么类型
      if v, ok := v1.(float64); ok {
      	fmt.Println(v, ok)
      }
      ```

## GO语言进阶

#### GO语言并发编程

###### Go语言并发编程之协程

* 与传统的系统级线程相比,协程的大优势在于其"轻量级",可以轻松创建上百万个而不会导致系统资源衰竭,而线程和进程通常多也不能超过1万个,这也是协程也叫轻量级线程的原因.

* go对协程的实现:go+函数名:启动一个协程执行函数体

  * ```
    package main
    import (
    	"fmt"
    	"time"
    )
    func test_Routine() {
    	fmt.Println("This is one routine!")
    }
    func main() {
    	go test_Routine()
    	time.Sleep(1)
    }
    ```

* Go并发编程之channel

  * Channel:Go语言在语言级别提供的goroutine间的通讯方式

    * 声明方式:var chanName chan ElementType

      * ```
        package main
        
        import (
        	"fmt"
        	"strconv"
        )
        
        func Add(x, y int, quit chan int) {
        	z := x + y
        	fmt.Println(z)
        	quit <- 1
        }
        func Read(ch chan int) {
        	value := <-ch
        	fmt.Println("value:" + strconv.Itoa(value))
        }
        func Write(ch chan int) {
        	// ch < -10
        }
        func main() {
        	chs := make([]chan int, 10)
        	for i := 0; i < 10; i++ {
        		chs[i] = make(chan int)
        		go Add(i, i, chs[i])
        	}
        	for _, v := range chs {
        		<-v
        	}
        }
        ```

        

  * 缓存channel

    * ```
      package main
      
      import (
      	"fmt"
      	"time"
      )
      
      var ch chan int
      
      func test_channel() {
      	// ch := make(chan int)
      	ch <- 1
      	fmt.Println("come to end goroutline 1")
      }
      func main() {
      	ch = make(chan int, 2) //值是1和值是0的适合输出执行顺序不同
      	go test_channel()
      	time.Sleep(2 * time.Second)
      	fmt.Println("running end!")
      	<-ch
      	time.Sleep(time.Second)
      }
      ```

  * select

    * 是linux很早就引入的函数,用来实现非阻塞的一种方式

    * go语言提供语言级别支持slelect关键字,用于处理异步IO问题

      * ```
        select{
            case <-chan1://如果chan1成功读取到数据,则进行该case处理语句
            case chan2 <-1://如果成功向chan2写入数据,则进行该case处理
            default: //如果上面都没有成功,则进入default处理流程
        }
        ```

    * select 三个代码示例

      * 示例1

        ```
        package main
        
        import (
        	"fmt"
        	"time"
        )
        
        func main() {
        	ch := make(chan int)
        	go func(ch chan int) {
        		ch <- 1
        	}(ch)
        	// time.Sleep(2)//加这一句执行结果不同
        	select {
        	case <-ch:
        		fmt.Print("come to read ch!")
        	default:
        		fmt.Print("come to default!")
        	}
        }
        ```

      * 示例2

        ```
        package main
        
        //超时控制的经典实现
        import (
        	"fmt"
        	"time"
        )
        
        func main() {
        	ch := make(chan int)
        	timeout := make(chan int, 1)
        	go func() {
        		time.Sleep(time.Second)
        		timeout <- 1
        	}()
        	select {
        	case <-ch:
        		fmt.Println("come to read ch!")
        	case <-timeout:
        		fmt.Println("come to timeout")
        	}
        	fmt.Print("end of code!")
        }
        ```

      * 示例3

        ```
        package main
        
        //超时控制的经典实现
        import (
        	"fmt"
        	"time"
        )
        
        func main() {
        	ch := make(chan int)
        	select {
        	case <-ch:
        		fmt.Println("come to read ch!")
        	case <-time.After(time.Second):
        		fmt.Println("come to timeout")
        	}
        	fmt.Print("end of code!")
        }
        ```

* 深入Go协程编程

  * 协程与线程质的不同

    * 协程任务的业务代码主动要求切换,即主动让出执行权
    * 发生了IO,导致执行阻塞

  * 线程与协程代码对比(线程几百个就卡,协程十万个都不算多)

    * Java线程

      ```
      public class MyThread{
          public static void main(String[] args){
              new Thread(new Test_thread()).start();
              new Thread(new Test_thread2()).start()
          }
      }
      class Test_thread implements Runnable{
              public void run(){
                  for(int i=0;i<100;i++){
                      System.out.println("thread 1:"+i)
                  }
              }
      }
      class Test_thread2 implements Runnable{
          public void run(){
              for(int i=100;i<200;i++){
                  System.out.println("thread 2:+i);
              }
          }
      }
      ```

    * Go协程

      ```
      package main
      import (
      	"fmt"
      	// "runtime"
      	"strconv"
      	"time"
      )
      func main() {
      	//协程1
      	go func() {
      		for i := 1; i < 100; i++ {
      			if i == 10 {
      				// runtime.Gosched() //主动要求切换CPU
      				<-ch //遇到阻塞主动让出,否则一直执行下去
      			}
      			fmt.Println("routine 1:" + strconv.Itoa(i))
      		}
      	}()
      	//协程2
      	go func() {
      		for i := 100; i < 200; i++ {
      			fmt.Println("routine 2:" + strconv.Itoa(i))
      		}
      	}()
      	time.Sleep(time.Second)
      }
      ```

#### GO语言JSON与MD5

###### Go语言的JSON

* 使用的库

  * Go语言内置的encoding/json标准库:性能不太好
  * github.com/pquerna/ffjson/ffjson

* json使用代码

  ```
  package main
  
  import (
  	"encoding/json"
  	"fmt"
  )
  
  type Student struct {
  	Name string //外部要引用要首字母大写
  	Age  int
  }
  
  func main() {
  	//对数组类型的json编码
  	x := [5]int{1, 2, 3, 4, 5}
  	s, err := json.Marshal(x)
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(string(s))
  	//对map类型进行json编码
  	m := make(map[string]float64)
  	m["zhangsan"] = 100.4
  	s2, err2 := json.Marshal(m)
  	if err2 != nil {
  		panic(err2)
  	}
  	fmt.Println(string(s2))
  	//对对象进行json编码
  	student := Student{"zhangsan", 26}
  	s3, err3 := json.Marshal(student)
  	if err3 != nil {
  		panic(s3)
  	}
  	fmt.Println(string(s3))
  	//对s3进行解码
  	var s4 interface{}
  	json.Unmarshal([]byte(s3), &s4)
  	fmt.Printf("%v", s4)
  }
  ```

* md5使用

  ```
  package main
  
  import (
  	"crypto/md5"
  	"fmt"
  )
  
  func main() {
  	Md5Inst := md5.New()
  	Md5Inst.Write([]byte("zhangsan"))
  	Result := Md5Inst.Sum([]byte(""))
  	fmt.Printf("%x\n\n", Result)
  }
  ```

#### GO语言HTTP

* go语言标准库内建提供了net/http包

* 处理http请求:使用net/http包提供的http.ListenAndServer()方法,可以在指定的地址进行监听,开启一个HTTP.服务端该方法的原型如下:func ListenAndServe(addr string,hander Handler)error

  * 该方法用于在指定的TCP网络地址addr进行监听,然后调用服务端处理程序来处理传入的链接请求.该方法有两个参数:第一个参数addr即监听地址;第二个服务表示服务端处理程序,通常为空,这意味着服务端调用http.DefaultServeMux进行处理,而服务端编写的业务逻辑处理程序http.Handle()或http.HandleFunc()默认注入http.DefaultServeMux中

* 处理https请求

  * func ListenAndServeTLS(addr string,certFile string,keyFile string,handler Handler) error
  * http.HandleFunc()方法接受两个参数
    * 第一个参数是http请求的目标路径"/hello",该参数值可以是字符串,也可以是字符串形式的正则表达式,第二个参数指定具体的回调方法,比如helloHandler
    * 当我们的程序运行起来后,访问-http://localhost:8080/hello,程序就会去调用hellloHandler()方法中的业务逻辑程序

* http中get和postdaima

  * ```
    //get代码
    package main
    
    import (
    	"fmt"
    	"io/ioutil"
    	"net/http"
    )
    
    func main() {
    	resp, err := http.Get("http://www.baidu.com")
    	if err != nil {
    		panic(err)
    	}
    	defer resp.Body.Close()
    	body, err := ioutil.ReadAll(resp.Body)
    	fmt.Println(string(body))
    }
    //post代码
    package main
    
    import (
    	"fmt"
    	"io/ioutil"
    	"net/http"
    	"strings"
    )
    
    func main() {
    	resp, err := http.Post("http://www.baidu.com", "application/x-www.form-urlencoded", strings.NewReader("id=1"))
    	if err != nil {
    		panic(err)
    	}
    	defer resp.Body.Close()
    	body, err := ioutil.ReadAll(resp.Body)
    	fmt.Println(string(body))
    }
    ```

#### GO语言正则表达式

###### Go语言标准库内建提供了regexp包

* 正则字母含义

  * .匹配除换行符以外的任意字符
  * \w匹配字母或数字或下划线或汉字
  * \s匹配任意的空白符
  * \d匹配数字
  * \b匹配单词的开始或结束
  * ^匹配字符串的开始
  * $匹配字符串的结束
  * *重复零次或更多次
  * +重复一次或更多次
  * ?重复零次或一次
  * {n}重复n次
  * {n,}重复n次或更多次
  * {n,m}重复n到m次
  * 捕获(exp) 匹配exp,并捕获文本到自动命名的阻组里
  * (?<name>exp)  匹配exp,并捕获文本到名称为name的组里,也可以写成(?'name'exp)
  * (?:exp)匹配exp,不捕获匹配的文本,也不给此分组分配组号

* 正则函数

  * func Match(pattern string,b []byte)(matched bool,err error)
  * func MatchString(pattern string,s string)(matched bool,err error)
  * func MustCompile(str string)*Regexp
  * func(re * Regexp)FindAllString(s string,n int)[]string

* 代码演示

  * ```
    package main
    
    import (
    	"fmt"
    	"regexp"
    )
    
    func main() {
    	isok, _ := regexp.Match("[a-zA-Z]{3}", []byte("zhl"))
    	fmt.Printf("%v", isok)
    	reg1 := regexp.MustCompile(`^z(.*)l$`)
    	result1 := reg1.FindAllString("zhangsan", -1)
    	fmt.Printf("%v\n", result1)
    	reg2 := regexp.MustCompile(`z(.{1})(.{1})(.*)l$`)
    	result2 := reg2.FindAllStringSubmatch("zhangsan", -1)
    	fmt.Printf("%v\n", result2)
    
    }
    
    ```

  * 代码演示2

    ```
    package main
    
    import (
    	"fmt"
    	"io/ioutil"
    	"net/http"
    	"regexp"
    )
    
    func main() {
    	url := "https://movie.douban.com/subject/1292052/"
    	resp, err := http.Get(url)
    	if err != nil {
    		panic(err)
    	}
    	defer resp.Body.Close()
    	sHtml, _ := ioutil.ReadAll(resp.Body)
    	// fmt.Println(string(sHtml))
    	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
    	result := reg.FindAllStringSubmatch(string(sHtml), -1)
    	// fmt.Printf("%v", result)
    	// for _, v := range result {
    	// 	fmt.Println(v[1])
    	// }
    	fmt.Println(result[0][1])
    	reg1 := regexp.MustCompile(`<strong\s*class="ll\s*rating_num"\s*property="v:average">(.*)</strong>`)
    	result1 := reg1.FindAllStringSubmatch(string(sHtml), -1)
    	fmt.Println(result1[0][1])
    }
    ```

#### GO语言Mysql与Redis

###### Go语言使用mysql驱动包:https://github.com/Go-SQL-Driver/Mysql

* 使用sql.open()函数用来打开一个注册过的数据库驱动,Go-MySQL-Driver中注册了mysql这个数据库的驱动,第二个参数是DNS(Data Source Name),它是Go-MySQL-Driver定义的一些数据库链接和配置信息.它支持如下格式

  * user@unix(/path/to/socket)/dbname?charset=utf8
  * user:password@tcp(localhost:5555)/dbname?chaset=utf8

* Go语言操作mysql代码

  * ```
    package main
    
    import (
    	"database/sql"
    	"fmt"
    
    	_ "github.com/go-sql-driver/mysql"
    )
    
    func main() {
    	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cost?charset=utf8")
    	if err != nil {
    		panic(err)
    	}
    	//插入
    	// stmt, err := db.Prepare("insert test set t_id=?,t_name=?,t_time=?")
    	// res, err := stmt.Exec(998, "zhangsan", "2019-01-02")
    	// id, err := res.LastInsertId()
    	//修改
    	// stmt, err := db.Prepare("update test set t_name=? where t_id=?")
    	// res, err := stmt.Exec("lisi", 999)
    	// id, err := res.RowsAffected()
    	//数据库一主多从,从库多是用来查询的
    	//查询
    	rows, err := db.Query("select * from test where t_id=999")
    	if err != nil {
    		panic(err)
    	}
    	for rows.Next() {
    		var t_id int
    		var t_name string
    		var t_time string
    		err = rows.Scan(&t_id, &t_name, &t_time)
    		fmt.Println(t_id, t_name, t_time)
    	}
    	//在go中创建的变量必须调用一次
    	// fmt.Println(id)
    }
    ```

###### Go语言使用的Redis驱动包:https://github.com/astaxie/goredis

* 驱动包需设置gopath下载使用

* [下载安装redis并启动]()

* redis的基本操作类型

  * string
  * hash
  * set
  * list
  * zset

* redis代码示例

  ```
  package main
  
  import (
  	"fmt"
  
  	"github.com/astaxie/goredis"
  )
  
  func main() {
  	var client goredis.Client
  	client.Addr = "127.0.0.1:6379"
  	//存入数据到redis
  	err := client.Set("test", []byte("hello beifeng"))
  	if err != nil {
  		panic(err)
  	}
  	//从redis获取数据
  	res, err := client.Get("test")
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(string(res))
  	//库中hmset方法使用
  	f := make(map[string]interface{})
  	f["name"] = "zhangsan"
  	f["age"] = 12
  	f["sex"] = "male"
  	err = client.Hmset("test_hash", f)
  	if err != nil {
  		panic(err)
  	}
  	//库中 zset方法使用
  	_, err = client.Zadd("test_zset", []byte("beifeng"), 100)
  	if err != nil {
  		panic(err)
  	}
  }
  ```

* redis命令行操作查看

  * ```
    //启动服务
    redis-server.exe
    //启动客户端查看数据
    E:\quickSorftWare\redis>redis-cli.exe
    127.0.0.1:6379> get test
    "hello golang"
    127.0.0.1:6379> type test_hash
    none
    127.0.0.1:6379> type test_hash
    hash
    127.0.0.1:6379> hgetall test_hash
    1) "age"
    2) "12"
    3) "sex"
    4) "male"
    5) "name"
    6) "zhangsan"
    127.0.0.1:6379> hgetall test_zadd
    (empty list or set)
    127.0.0.1:6379> type test_zadd
    none
    127.0.0.1:6379> type test_zset
    zset
    127.0.0.1:6379> zrange test_zset 0 -1
    1) "golang"
    ```

#### 聊天室实战项目

###### 服务端

* 代码

  ```
  package main
  
  import (
  	"fmt"
  	"log"
  	"net"
  	"os"
  	"strings"
  )
  
  const (
  	LOG_DIRECTORY = "./test.log"
  )
  
  var onlineConns = make(map[string]net.Conn)
  var messageQueue = make(chan string, 1000)
  var quitChan = make(chan bool)
  var logFile *os.File
  var logger *log.Logger
  
  func CheckError(err error) {
  	if err != nil {
  		// fmt.Println("Error :%s",err.Error())
  		// os.Exit(1)
  		panic(err)
  	}
  }
  func ProcessInfo(conn net.Conn) {
  	buf := make([]byte, 1024)
  	// defer conn.Close()
  	defer func(conn net.Conn) {
  		addr := fmt.Sprintf("%s", conn.RemoteAddr())
  		delete(onlineConns, addr)
  		conn.Close()
  		for i := range onlineConns {
  			fmt.Println("closed" + i)
  		}
  	}(conn)
  	for {
  		numofBytes, err := conn.Read(buf)
  		// CheckError(err)
  		if err != nil {
  			// continue
  			break
  		}
  		if numofBytes != 0 {
  			// remoteAddr := conn.RemoteAddr()
  			// fmt.Print(remoteAddr)
  			// fmt.Printf("Has received this message: %s\n", string(buf[0:numofBytes]))
  			message := string(buf[0:numofBytes])
  			//每次收到消息,写入到messageQueue
  			messageQueue <- message
  		}
  	}
  }
  func ConsumeMessage() {
  	for {
  		select {
  		case message := <-messageQueue:
  			//对消息进行解析
  			doProcessMessage(message)
  		case <-quitChan:
  			break
  
  		}
  	}
  }
  func doProcessMessage(message string) {
  	contents := strings.Split(message, "#")
  	if len(contents) > 1 {
  		addr := contents[0]
  		// sendMessage := contents[1]
  		sendMessage := strings.Join(contents[1:], "#")
  		//防止有空格
  		addr = strings.Trim(addr, " ")
  		if conn, ok := onlineConns[addr]; ok {
  			_, err := conn.Write([]byte(sendMessage))
  			if err != nil {
  				fmt.Println("online conns sned failure!")
  			}
  		}
  	} else {
  		contents := strings.Split(message, "*")
  		if strings.ToUpper(contents[1]) == "LIST" {
  			var ips string = ""
  			for i := range onlineConns {
  				ips = ips + "|" + i
  			}
  			if conn, ok := onlineConns[contents[0]]; ok {
  				_, err := conn.Write([]byte(ips))
  				if err != nil {
  					fmt.Println("online conns send failure")
  				}
  			}
  		}
  	}
  }
  func main() {
  	// onlineConns=make(map(string)net.Conn)
  	logFile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE, 0)
  	if err != nil {
  		fmt.Println("logFile create failure")
  		os.Exit(-1)
  	}
  	defer logFile.Close()
  	logger = log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
  	listen_socket, err := net.Listen("tcp", "127.0.0.1:8062")
  	CheckError(err)
  	defer listen_socket.Close()
  	fmt.Println("Server is wating...")
  	logger.Println("I am logger test")
  	//开启一个协程处理messageQueue
  	go ConsumeMessage()
  	for {
  		conn, err := listen_socket.Accept()
  		CheckError(err)
  		//连接之前将conn对象存储到onlineConns映射表中
  		addr := fmt.Sprintf("%s", conn.RemoteAddr())
  		onlineConns[addr] = conn
  		for i := range onlineConns {
  			fmt.Println(i)
  		}
  		//开一个协程处理
  		go ProcessInfo(conn)
  	}
  }
  ```

###### 客户端

* 代码

  ```
  package main
  
  import (
  	"bufio"
  	"fmt"
  	"net"
  	"os"
  	"strings"
  )
  
  func CheckError(err error) {
  	if err != nil {
  		panic(err)
  	}
  }
  func MessageSend(conn net.Conn) {
  	var input string
  	for {
  		reader := bufio.NewReader(os.Stdin)
  		data, _, _ := reader.ReadLine()
  		input = string(data)
  		if strings.ToUpper(input) == "EXIT" {
  			conn.Close()
  			break
  		}
  		_, err := conn.Write([]byte(input))
  		if err != nil {
  			conn.Close()
  			fmt.Println("client connect failure:" + err.Error())
  			break
  		}
  	}
  }
  func main() {
  	conn, err := net.Dial("tcp", "127.0.0.1:8062")
  	CheckError(err)
  	defer conn.Close()
  	//协程处理消息
  	go MessageSend(conn)
  	// conn.Write([]byte("Hello golang"))
  	buf := make([]byte, 1024)
  	for {
  		_, err := conn.Read(buf)
  		// CheckError(err)
  		if err != nil {
  			fmt.Println("你已退出,欢迎使用golang通讯")
  			os.Exit(0)
  		}
  		fmt.Println("receive server message connent:" + string(buf))
  	}
  	fmt.Println("client program end!")
  
  }
  ```

#### 爬虫实战项目

###### 认识beego框架

* 简介

  * beggo是一个快速开发Go应用的http框架,它可以用来快速开发API、Web及后端服务等各种应用，是一个RESTFUL的框架，主要设计灵感来源于tornado、sinatra和flask这三个框架，但是结合了Go本身的一些特性(interface、struct嵌入等)而设计的一个框架

* beego的架构

  * 基础模块：cache、config、context、httplibs、logs、orm、session、toolbox

* beego的执行逻辑：main文件监听端口接收请求

  * 路由功能-参数过滤-controller
  * 视屏输出-输出过滤-controller
  * controller交互
    * 辅助工具包
    * model-数据库
    * session管理
    * 日志处理
    * 缓存处理

* beego的项目结构-mvc

  * conf
    * app.conf
  * controllers
    * admin
    * default.go
  * main.go
  * models
    * models.go
  * static
    * css
    * ico
    * img
    * js
  * views
    * admin
    * index.tpl

* beego的安装：go get github.com/astaxie/beego

* bee工具：帮我们生成项目结构

  * new命令：new命令是新加一个web项目，我们在命令下执行bee new<项目名>就可以创建一个新的项目。但是注意该命令必须在$GOPATH/src下执行
  * api命令：上面的new命令是用来新建web项目，不过很多用户使用beego开发api应用。这个api命令就是创建api应用的
  * 配置bee环境变量

* beego的参数配置

  * beego默认会解析当前应用下的conf/app.conf

    * ```
      appname = WebDemo1
      httpport = 8080
      runmode = dev
      [dev]
      httpport=8080
      [test]
      httpport=8081
      [prod]
      httpport=8082
      ```

* beego的路由配置

  * 在router.go中定义路由访问,访问路径在controller下
  * 基础路由
    * beego.Router("/",&controllers.MainController{})
    * beego.Router("/admin",&admin.UserController{})
    * beego.Router("/admin/index",&admin.ArticleController{})
    * beego.Get("/",func(ctx *context.Context)){ctx.Output.Body([]byte("hello word"))}
  * 自定义方法及RESTful规则
    * beego.Router("/",&IndexController{},"*:Index")
      * *表示任意的method都执行该函数
      * 使用httpmethod:funcname格式来表示
      * 多个不同的格式,使用;分开
      * 多个method对应同一个funcname,method之间通过,来分割
    * beego.Router("/api/list",&RestController{},"*ListFood")
    * beego.Router("/api/create",&RestController{},"post:CreateFood")
    * beego.Roter("/api/update",&RestController{},"put:updateFood")
    * beego.Router("/api/delete",&RestController{},"delete:DeleteFood")
    * beego.Router("/api",&RestCOntroller{},"get,post:ApiFunc")
    * beego.Router("/simple",&SimpleController{},"get:GetFunc;post:PostFunc")

* Beego框架之Controller

  * 基于beego的Controller设计,只需要匿名组合beego.Controller就可以了

    * ```
      import ("github.com/astaxie/beego")
      type MainController struct{
         	beego.Controller
      }
      func (this *MainController) Get(){
          this.Data["Website"]="beego.me"
          this.Data["Email]="astaxie@gmail.com"
          this.TplName="index.tpl"
      }
      ```

  * beego.Controller拥有很多方法,其中包括Lnit、Prepare、Post、Get、Delete、Head等方法，我们可以通过重写的方式来实现这些方法
  * 我们可以通过各种方式获取数据，然后赋值到this.Data中，这是一个用来存储输出数据的map，可以赋值任意类型的值，这里我们只是简单举例输出数据的map，可以赋值任意类型的值，这里我们只是简单举例输出两个字符串
  * 最后一个就是需要去渲染的模板，this.TplName就是需要渲染的模板，这里指定了index.tpl,如果用户不设置此参数，那么默认会去到模板目录的Controller/<方法名>.tpl查找，例如maincontroller/get.tpl(文件、文件夹需小写)
  * 用户设置了模板之后系统会自动的调用Render函数(这个函数是在beego.Controller中实现的)，所以无需用户自己来调用渲染

* Beego框架之请求数据处理

  * 我们经常需要获取用户传递的数据，包括Get、POST等方式的请求，beego里会自动解析这些数据，我们可以通过如下方式获取数据

    * GetString(key string) string

    * getString(key string) []string

    * GetInt(key string)(int64,error)

    * GetBool(key string)(bool,error)

    * GetFloat(key string)(float64,error)

      * ```
        func (c *TestController) Test() {
        	id :=c.GetString("id")
        	c.Ctx.WriteString(id)
        	name :=c.Input().Get("name")
        }
        ```

    * 表单操作

      * ```
        type User struct{
            UserName string
            Password string
        }
        func (c *TestInputController)Get(){
                c.Ctx.WriteString(`<html><form action="http://localhost:8080/test" method="post"> 
                	<input type="text" name="Username"/>
                	<input type="password" name="Password"/>
                	<input type="submit" value="提交"/>
                </form></html>`)
            }
        func (c *TestInputController) Post(){
            u :=User{}
            if err:=c.ParseForm(&u); err!=nil{
                //process err
            }
            c.Ctx.WriteString("Username"+u.Username)
        }
        ```

      * 获取RequestBody里边的内容

        * 在配置文件里设置copyrequestbody=true

        * 在Controller中获取

          * ```
            func （this *ObjectController) Post(){
                var ob models.Object
                json.Unmarshal(this.Ctx.Input.RequestBody,&ob)
                objectid :=models.AddOne(ob)
                this.Data["json"]={\"ObjectId\":\""+objectId+"\"}
                this.ServeJSON()
            }
            ```

* Beego框架之cookie与session

  * Cookie

    * this.Ctx.SetCookie("name",name,maxage,"/")

    * this.Ctx.SetCookie("pwd",md5([]byte(pwd)),maxage,"/")

    * this.Ctx.GetCookie

    * ```
      c.Ctx.SetCookie("name",u.UserName,-1,"/")//存储
      name :=c.Ctx.GetCookie("name")//获取
      ```

  * Session：beego内置了session模块，目前session模块支持的后端引擎包括memory、coockie、file、mysql、redis、couchbase、memcache、postgres，用户也可以根据相应的interface实现自己的引擎

    * 开启session：在main入口函数添加：beego.BConfig.WebConfig.Session.SessionOn=true
    * 或者通过配置文件配置如下
      * session=true
    * session几个方便的方法
      * SetSession(name string,value interface{})
      * GetSession(name string) interface{}
      * DeISission(name string)
      * SessionRegenerateID()
      * DestroySession()
    * session存储：c.Ctx.SetSession("name",u.UserName)
    * session获取：name :=c.GetSession("name")

* Beego框架之model

  * beego ORM是一个强大的Go语言ERM框架。她的灵感主要来自Django ORM和SQLALchemy

    * 已支持数据库驱动
      * Mysql：github.com/go-sql-driver/mysql
      * PostGreSql:github.com/lib/pq
      * Sqlite3:github.com/mattn/go-Sqlite3

  * 模型(models)Orm特性

    * 支持Go的所有类型存储
    * 轻松上手,采用简单的CRUD风格
    * 自动Join关联表
    * 跨数据库兼容查询
    * 允许直接使用sql查询/映射
    * 严格完整的测试保证,ORM的稳定与健壮

  * 安装ORM:go get github.com/astaxie/beego/orm

  * orm提供的方法

    * ```
      type Ormer interface{
          Read(interface{},...string) error
          ReadOrCreate(interface{},string ,...string)(bool,int64,error)
          insert(interface{})(int64,error)
          Update(interface{},...string)(int64,err)
          Delete(iterface{})(int64,error)
          LoadRelated(interface{},string,...interface{})(int64,error)
          QueryM2M(interface{},string)QueryM2Mer
          QueryTable(interface{})QuerySeter
          Using(string)error
          Begin()error
          Commit()error
          Rollback()error
          Raw(string,...interface{})RwaSeter
          Driver()Driver
      }
      ```

    * 使用时使用import导入对应驱动

    * 注册

      * ```
        orm.RegisterDataBase("default","mysql","root:root@/my_db?charset=utf8",30)
        orm.RegisterModel(new(User))	
        ```

  * 代码演示

    * 第一种

      ```
      package controllers
      import (
      	"github.com/astaxie/beego"
      	_ "github.com/go-sql-driver/mysql"
      	"github.com/astaxie/beego/orm"
      	"fmt"
      )
      //UseInfo操作的表是user_info 
      type Test struct{
      	Id int64
      	TName string
      	TTime string
      }
      type TestModelController struct{
      	beego.Controller
      }
      func (c *TestModelController) Get(){
      	orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/cost?charset=utf8")
      	orm.RegisterModel(new(Test))
      	o :=orm.NewOrm()
      	//插入操作
      	user :=Test{TName:"golangAI1",TTime:"2019-10-04"}
      	id, err :=o.Insert(&user)
      	//更新操作
      	// user.Id=111112
      	// user.TName = "lisi"
      	// id,err :=o.Update(&user)
      	//读取操作
      	// o.Read(&user)
      	c.Ctx.WriteString(fmt.Sprintf("id :%d err :%v",id,err))
      }
      ```

    * 第二种

      ```
      package controllers
      import (
      	"github.com/astaxie/beego"
      	_ "github.com/go-sql-driver/mysql"
      	"github.com/astaxie/beego/orm"
      	"fmt"
      )
      //UseInfo操作的表是user_info 
      type Test struct{
      	Id int64
      	TName string
      	TTime string
      }
      type TestModelController struct{
      	beego.Controller
      }
      func (c *TestModelController) Get(){
      	orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/cost?charset=utf8")
      	orm.RegisterModel(new(Test))
      	o :=orm.NewOrm()
      	//原生一种读法
      	// var maps []orm.Params
      	// num,err :=o.Raw("select * from test").Values(&maps)
      	// o.Raw("select * from test").Values(&maps)
      	// for _,v :=range maps{
      	//	c.Ctx.WriteString(fmt.Sprintf("test :%v",v))
      	// }
      	//原生第二种读法
      	var users []Test
      	o.Raw("select * from test").QueryRows(&users)
      	
      	c.Ctx.WriteString(fmt.Sprintf("id :%d err :%v",id,err))
      }
      ```

    * 事务

      ```
      o.Begin()
      //...
      user :=User(Name:"slene")
      id , err :=o.Insert(&user)
      if err ==nil{
          o.Commit
      }else{
          o.Rollback
      }
      ```

    * 开启日志:orm.Debug=true(会打印sql语句)

  * 构造查询

    * QueryBuilder提供了一个简便,流畅的SQL查询构造器.在不影响代码可读性的前提下用来快速的建立sql语句

      1. 获取QueryBuilder对象需要指定数据库驱动参数
      2. 构建查询对象
      3. 设置sql语句
      4. 执行sql语句

    * 提供的方法

      ```
      type QueryBuilder interface{
          Select(fields ...string)QueryBuilder
          From(tables ...string)QueryBuildre
          InnnerJoin(table string)QueryBuilder
          LeftJoin(table string)QueryBuilder
          RightJoin(table string)Querybuilder
          On(cond string)QueryBuilder
          Where(cond string)QueryBuilder
          And(cond string)QuryBuilder
          Or(cond string)Querybuilder
          In(vals ...string)QueryBuilder
          OrderBy(fields ...string)QueryBuilder
          Asc()QueryBuilder
          Desc()QueryBuilder
          Limit(limit int)QueryBuidler
          Offset(offset int)QueryBuilder
          GroupBy(fileds ...string)QueryBuilder
          Having(cond string)QueryBuilder
          Subquery(sub string,alias string)string
          String()string
      }
      ```

    * 代码演示

      ```
      package controllers
      import (
      	"github.com/astaxie/beego"
      	_ "github.com/go-sql-driver/mysql"
      	"github.com/astaxie/beego/orm"
      	"fmt"
      )
      //UseInfo操作的表是user_info 
      type Test struct{
      	Id int64
      	TName string
      	TTime string
      }
      type TestModelController struct{
      	beego.Controller
      }
      func (c *TestModelController) Get(){
      	orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/cost?charset=utf8")
      	orm.RegisterModel(new(Test))
      	o :=orm.NewOrm()
      	//采用QueryBuilder方式进行读取
      	var users []Test
      	qb, _:=orm.NewQueryBuilder("mysql")
      	qb.Select("t_name").From("Test").Where("id=?").Limit(1)
      	sql :=qb.String()
      	o.Raw(sql,2).QueryRows(&users)
      	c.Ctx.WriteString(fmt.Sprintf("user intfo:%v",users))
      }
      ```

* Beego框架之Controller调用model

  * model代码

    ```
    package models
    import (
    	_ "github.com/go-sql-driver/mysql"
    	"github.com/astaxie/beego/orm"
    )
    //UseInfo操作的表是user_info 
    type Test struct{
    	Id int64
    	TName string
    	TTime string
    }
    func init(){
    	orm.Debug=true//是否开启调试模式,开启输出sql
    	orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/cost?charset=utf8")
    	orm.RegisterModel(new(Test))
    
    }
    func AddUser(test *Test)(int64 ,error){
    	o :=orm.NewOrm()
    	id ,err:=o.Insert(test)
    	return id,err
    }
    ```

  * Controller代码

    ```
    package controllers
    import (
    	"github.com/astaxie/beego"
    	"WebDemo1/models"
    )
    type TestModelController struct{
    	beego.Controller
    }
    func (c *TestModelController) Get(){
    	user :=models.Test{Id:13,TName:"cgetm1",TTime:"2019-10-05"}
    	models.AddUser(&user)
    	c.Ctx.WriteString("call model success!")
    }
    ```

* beego框架之view

  * 基本语法

    * go统一使用了{{和}}作为左右标签,没有其它的标签符号,如果您是想要修改为其它符号,可以修改配置文件
    * 使用.来访问当前位置的上下文
    * 使用$来引用当前模板根级的上下文

  * 代码演示

    * model代码

      ```
      func ReadUserInfo(users *[]Test){
      	qb, _:=orm.NewQueryBuilder("mysql")
      	qb.Select("*").From("test")
      	sql :=qb.String()
      	o.Raw(sql).QueryRows(users)
      }
      ```

    * controller代码

      ```
      package controllers
      import (
      	"github.com/astaxie/beego"
      	"WebDemo1/models"
      )
      type TestViewController struct{
      	beego.Controller
      }
      func (c *TestViewController) Get(){
      	c.Data["Title"]="hello golang"
      	c.Data["IsDisplay"]=true
      	var users []models.Test
      	models.ReadUserInfo(&users)
      	c.Data["users"]=users
      	c.Data["Content"]="golang view example"
      	c.TplName="test_view.tpl"
      }
      ```

    * view代码

      ```
      <!DOCTYPE html>
      <html lang="en">
      <head>
      	<meta charset="UTF-8">
      	<title>{{.Title}}</title>
      </head>
      <body>
      	{{if .IsDisplay}}
      		<em>{{.Content}}</em>
      	{{else}}
      		<em>{{.Content2}}</em>
      	{{end}}
      
      	{{range .users}}
      		{{.TName}}
      	{{end}}
      </body>
      </html>
      ```

* beego框架之config、httplib、context

  * config

    * 配置文件解析

      * 需要包:go get gethub.com/astaxie/beego/config

    * 初始化一个解析器对象

    * ```
      iniconf ,err :=NewConfig("ini","testini.conf")
      if err !-nil{
          t.Fatal(err)
      }
      iniconf.String("appname")
      ```

    * 解析器对象支持的函数

      * ```
        Set(key,val string)error
        string(key string)string
        Int(key string)(int,error)
        Int64(key string)(int64,error)
        Bool(key string)(bool,error)
        Float(key string)(float64,error)
        DIY(key string)(interface{},error)
        ```

    * 解析器函数

      * ini配置文件支持section操作,key通过section::key的方式获取

      * 例如下面这样的配置文件

        * ```
          [demo]
          key1="asta"
          key2="xie"
          //那么可以通过iniconf.String("demo::key2")获取值
          ```

  * httplib

    * httplib库主要用来模拟客户端发送请求,类似于curl工具,支持Jquery类似的链式操作.使用起来相当的方便.通过如下方式进行安装:go get github.com/astaxie/beego/httplib

      * ```
        package controllers
        import (
        	"github.com/astaxie/beego"
        	"github.com/astaxie/beego/httplib"
        )
        type TestHttpLibController struct{
        	beego.Controller
        }
        func (c *TestHttpLibController) Get(){
        	req :=httplib.Get("http://douban.com")
        	str,err :=req.String()
        	if err !=nil{
        		panic(err)
        	}
        	c.Ctx.WriteString(str)
        }
        ```

  * context对象

    - context对象是对input对象和output对象的封装

      - Redirect
      - Abort
      - WriteString
      - GetCookie
      - SetCookie

    - input对象

      - Url:用户请求的RequestUrl,例如/hi?id=1001
      - IP:
      - Proxy:返回用户代理请求的所有ip
      - Refer:返回请求的refer信息
      - SubDomailns:返回请求域名的根域名
      - Port

    - Output对象

      - Header:设置输出的Header信息

      - Body:设置输出内容信息

      - Coockie

      - Json:把Data格式化为Json输出

      - ```
        m:=make(map[string]float64)
        m["zhangsan"]=98.7
        c.Ctx.Output.JSON(m,false,false)//后两个参数是缩进和编码
        ```

###### Beego框架之爬虫项目实战

* 了解与准备
  * 抓到的网页会重复,去重处理
    * 布隆过滤器
    * 哈希存储
  * 标签匹配
    * 正则表达式
    * beautiful soup或xml这种标签提取库
  * 动态内容
    * phantomjs
    * selenium
  * mysql数据库建表
  * redis环境
* [github代码]()
* 使用此命令在命令行后台爬取(git工具):curl "http://localhost:8080/crawl_movie"



