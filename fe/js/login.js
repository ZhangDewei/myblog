var doc = document;
	username = doc.getElementById("username"),
	password = doc.getElementById("password"),
	validUser = doc.getElementById('valid-user'),
	validPwd = doc.getElementById('valid-password'),
	subMit = doc.getElementById('commit');

username.onblur = function(){
	if(username.value){
		$.ajax({
			url: '/register/checkuser',
			data:{
				username: username.value
			},
			type:'post'
		}).done(function(data){
			console.log(data.content)
			if(data.content == true){
				validUser.style.display="inline-block";	
			}else{
				validUser.style.display="none";
			}
		}).fail(function(){
			alert("提交异常");
		});
	}
};  // function finish;

password.onblur = function(){
	if(password.value && username.value){
		$.ajax({
			url:"/login/checkuser",
			data:{
				username:username.value,
				password:password.value
			},
			type:'post'
		}).done(function(data){
			if(!data){
				validPwd.style.display='inline-block';
			}else{
				validPwd.style.display='none';
			};
		}).fail(function(){
			alert("服务器异常");
		});
	};
}; // function finish;

subMit.onclick = function(event){
	var event = event || window.event;
	if(!username.value || !password.value){
		alert("请填写完整");
		event.stopPropagation;
		event.cancelBubble = true;
		return false;
	};
};
