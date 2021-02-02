# 中通开放平台SDK（Go版）

## 参考

[中通开放平台SDK](https://github.com/ZTO-Express)

## 使用方式

```
go get -u github.com/chiahan1123/zopsdk-go
```

```
client, err := zop.NewClient("kfpttestCode", "kfpttestkey==")
if err != nil {
	// handle error
	return
}
params := make(map[string]string)
params["request"] = `[{"partnerCode":"360844819234","companyCode":"GP1551922487","reason":"客户取消"}]`
resp, err := client.Execute(context.Background(), &zop.Request{
    URL:    "https://japi-test.zto.com/cancelOrder",
    Params: params,
})
if err != nil {
	// handle error
	return
}
fmt.Println(resp)
```
