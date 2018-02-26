# Serverless Containers

```
docker build -t kelseyhightower/openwhisk-go-container:0.0.1 .
```

```
docker push kelseyhightower/openwhisk-go-container:0.0.1
```

```
bx wsk action delete custom
```

```
bx wsk action create custom \
  --docker kelseyhightower/openwhisk-go-container:0.0.1
```

```
bx wsk action invoke --result custom --param name Kelsey
```
