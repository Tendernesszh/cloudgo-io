
####我的cloudgo-io实现了如下功能：<br>

##### 1.支持静态文件服务：<br>
这个功能直接用老师上课给的cloud-static中的代码就可以实现，借鉴老师的代码之后在浏览器输入<br>
http://localhost:8080/static/ 以后能看到asset文件下的index.html的网页版。<br>
结果图如下：<br>
![image](https://github.com/Tendernesszh/cloudgo-io/blob/master/%E9%9D%99%E6%80%81%E6%96%87%E4%BB%B6%E6%9C%8D%E5%8A%A1.png)
<br>
##### 2.支持简单js访问<br>
此功能可借鉴老师给的cloud-templates中的代码来完成，apitest.go中会完成模板填充，在浏览<br>
器输入 http://localhost:8080/api/test 后<br>
结果图如下：<br>
![image](https://github.com/Tendernesszh/cloudgo-io/blob/master/js%E8%AE%BF%E9%97%AE.png)
<br>
##### 3.提交表单，并输出一个表格<br>
用templates的userinfo.tmpl做模板目录，提交表单之后record.go文件中的<br>
recordHandler函数会输出表格<br>
结果图如下：<br>
![image](https://github.com/Tendernesszh/cloudgo-io/blob/master/%E6%B3%A8%E5%86%8C%E7%94%A8%E6%88%B7.png)<br>
![image](https://github.com/Tendernesszh/cloudgo-io/blob/master/%E6%B3%A8%E5%86%8C%E6%88%90%E5%8A%9F.png)<br>



##### 4.对unknown给出开发中的提示，返回码5xx<br>
在浏览器中输入 http://localhost:8080/unknown 后会提示错误信息<br>
结果如下：<br>
![image](https://github.com/Tendernesszh/cloudgo-io/blob/master/unknown.png)# cloudgo-io
