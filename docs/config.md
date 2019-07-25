### Config

We put configuration files in this directory.

* app.yml    - Application configurations
* config.yml - Services configurations
* config.go

Sample `config.yml` file

```yaml
debug:
  MySQL:
    host: 127.0.0.1
    port: 3306
    name: kalista-dev
    username: root
    password: mypassword
    connections: 60
    idles: 10
  Redis:
    host: 127.0.0.1
    port: 6379
  RabbitMQ:
    host: 127.0.0.1
    port: 5672
    user: guest
    password: guest

release:
  MySQL:
    host: 127.0.0.1
    port: 3306
    user: root
    password: mypassword
    connections: 60
    idles: 10
```

Sample `app.yml file

```yaml
components:
  MySQL: true
  Redis: true
  RabbitMQ: true
  nodes:
  Role: master

nodes:
  Role: slave
  MasterHost: 127.0.0.1:8080
```

We controls all behaviors in `config.go`. Obviously, you can change whatever you want to.
