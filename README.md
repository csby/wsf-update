# wsf-update
the service for wsf updating

## service
http serve on **127.0.0.1:9606** and accept **post** request

## function
### 1. update
```
{
    actione: "update",
    name: "name of the service",
    path: "execute file path of service",
    updateFile: "new execute file path of service",
    updateFolder: "temp folder path of new execute file"
}
```
### 2. restart
```
{
    actione: "restart",
    name: "name of the service"
}
```
