# Golang Article Sharing CRUD API Project  
CRUD Api with Go, Gorm and Fiber
## TECHNOLOGIES

- GOLANG ` 1.14 `
- FIBER
- GORM
- POSTGRESQL 

## API's PATHS
## For Article Service
##### For Get All Articles ```GET localhost:3000/api/v1/articles ``` <br/>
##### For Get Article By Id ```GET localhost:3000/api/v1/articles/{id} ``` <br/>
##### For Add Article ```POST localhost:3000/api/v1/articles ```  <br/>
###### Request Body: 
```json
{
    "description":"",
    "title":""
   
}
```
##### For Update Article ```PUT localhost:3000/api/v1/articles/update ```  <br/>
###### Request Body: 
```json
{
    "id":1,
    "description":"",
    "title":""
}
```

##### For Delete Article ```DELETE localhost:3000/api/v1/articles/{id} ```  <br/>





## For Comment Service
##### For Get All Comments ```GET localhost:3000/api/v1/comments ``` <br/>
##### For Get Comment By Id ```GET localhost:3000/api/v1/comments/{id} ``` <br/>
##### For Add Comment ```POST localhost:3000/api/v1/comments ```  <br/>
###### Request Body: 
```json
{
    "comment":"",
    "articleId":""
   
}
```
##### For Update Comment ```PUT localhost:3000/api/v1/comments/update ```  <br/>
###### Request Body: 
```json
{
    "id":1,
    "comment":""
}
```

##### For Delete Comment ```DELETE localhost:3000/api/v1/comments/{id} ```  <br/>


