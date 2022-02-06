## ToDoListProject

### 技术栈：

后端：GO+Beego+Gorm

go版本：1.17.6
Beego版本：1.12.2
Gorm:v1.22.5

前端：VUE

前端项目来源于“https://github.com/Q1mi/bubble_frontend.git”，使用时需要修改

（1）连接到后端的地址，vue.config.js

![](https://img2022.cnblogs.com/blog/1722983/202202/1722983-20220206154840667-1029377952.png)

（2）src/compoents/TodList.vue

![](https://img2022.cnblogs.com/blog/1722983/202202/1722983-20220206145530325-1743440379.png)



### 启动

后端：首先在mysql中，创建datatase：bubble 然后bee run  即可

前端：按照项目中提示 ，npm run serve，启动后访问：http://localhost:8080/#/

![](https://img2022.cnblogs.com/blog/1722983/202202/1722983-20220206144338521-1417915488.png)

### API

使用REST API，实例如下：

* POST：127.0.0.1:9090/v1/todo 

  请求报文：

  ```
  {
   "title":"gorm"
  }
  ```

  响应报文：

  ```
  {
      "code": 200,
      "message": "success"
  }
  ```

* PUT： http://localhost:8080/v1/todo/12 

  请求报文：

  ```
  {"status":true}
  ```

  响应报文：

  ```
  {
    "code": 200,
    "message": "success"
  }
  ```

* DELETE：http://localhost:8080/v1/todo/13

  请求报文：

  响应报文：

  ```
  {
    "code": 200,
    "message": "success"
  }
  ```

  

* GET： 127.0.0.1:9090/v1/todo/ 

  请求报文：

  响应报文：

  ```
  {
      "code": 200,
      "message": [
          {
              "id": 3,
              "title": "golang",
              "status": false
          },
          {
              "id": 5,
              "title": "k8s",
              "status": false
          },
          {
              "id": 6,
              "title": "java",
              "status": false
          },
          {
              "id": 9,
              "title": "gorm",
              "status": false
          },
          {
              "id": 10,
              "title": "json",
              "status": false
          },
          {
              "id": 11,
              "title": "mybatis",
              "status": false
          },
          {
              "id": 12,
              "title": "吃饭",
              "status": false
          },
          {
              "id": 14,
              "title": "gorm",
              "status": false
          }
      ]
  }
  ```

