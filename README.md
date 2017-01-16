# gett

![alt tag](https://github.com/shlomitsur/gett/blob/master/gett.jpg)

The aws url is http://ec2-23-21-246-3.compute-1.amazonaws.com:8080/v1/drivers, if it is down please let me know.

Didn't implemented api tests because I didn't find documented or sampled code yet, instead I used manual curl commands:
curl -i -H "Accept: application/json" -X PUT -d '{"Id" : 3, "Name": "shlomi", "LicenseNumber" : "lala" }'  http://ec2-23-21-246-3.compute-1.amazonaws.com:8080/v1/drivers/3

curl -i -H "Accept: application/json" -X DELETE  http://ec2-23-21-246-3.compute-1.amazonaws.com:8080/v1/drivers/2
