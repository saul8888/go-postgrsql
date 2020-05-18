### Create keys

it is created in the tokens folder

**private key**

`openssl genrsa -out private.rsa 1024`

**public key**

`openssl rsa -in private.rsa -pubout > public.rsa.pub`

**compilate**

`sudo docker-compose build`

`sudo docker-compose up`

**make a request first**

`/api/test`

**for get a tokens use**  /api/authentication **with body** 

```json
{
	"email": "test@example.com",
	"password":"test1234"
}
```



