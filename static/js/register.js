var doc = document,
	clickUser = doc.getElementById("click-user"),
	userHeader = doc.getElementById("user-header"),
	headerName = doc.getElementById("header-name");
	
	
	clickUser.onclick = function(){
		userHeader.click();
	}; // function finish;
	
	userHeader.onchange = function(){
		var imgName = this.files[0].name;
		if(imgName.length > 10){
			imgName = imgName.substring(0, 9);
		}
		headerName.innerHTML = imgName + '...';
		var reader = new FileReader();
		reader.onload = function(){
            //console.log(clickUser.children[0]);  该方法不含文本节点
			clickUser.childNodes[0].src = this.result;
		};
		reader.readAsDataURL(this.files[0]);
	}; // function finish;