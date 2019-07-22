[Gonrails](https://github.com/one-hole/gonrails) 是基于 [Gin](https://github.com/gin-gonic/gin) 改造的后端 MVC 开发框架

[![Go Report Card](https://goreportcard.com/badge/github.com/one-hole/kalista)](https://goreportcard.com/report/github.com/one-hole/kalista)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
--------

* QQ 群： 594171535
* 接头暗号： Heil Hydra

### 1. Install

`go get -u  github.com/one-hole/gonrails`

### 2. New

`gonrails new yourproject-name`


 * so if I want to create a project named __kalista__ , I just run command __`gonrails new kalista`__
 * the docs of the created project was in __[here](https://github.com/one-hole/gonrails)__

### 3. Generate

##### 3.1 Controller

 The command is `gonrails generate controller yourcontroller action list`

 eg:
  * `gonrails generate controller users index update show`
  * `gonrails generate controller admin/users index update show`

##### 3.2
