# mini-wallet

#### How to Run:

1. install package & run server
Make sure `go` already installed
then run command below:

```
go get all && go run engine/restapi/main.go
```
2. open related collection in postman, check the env variable `mw_base_url`, make sure its `http://localhost:8080`

3. run endpoint initialize my account for wallet `{{mw_base_url}}/api/v1/init`
4. then copy the response value property `token` into env variable `mw_api_token`. example response :

```js
{
  "data": {
    "token": "65613032313264332d616264362d343036662d386336372d383638653831346132343336"
  },
  "status": "success"
}
```
5. after put `token` into the env variable, you're ready to hit all endpoint.