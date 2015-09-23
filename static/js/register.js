var doc = document,
	clickUser = doc.getElementById("click-user"),
	userHeader = doc.getElementById("user-header"),
	headerName = doc.getElementById("header-name");
	
	
	clickUser.onclick = function(){
		userHeader.click();
	}; // function finish;
	
	userHeader.onchange = function(){
		headerName.innerHTML = this.files[0].name;
		var reader = new FileReader();
		reader.onload = function(){
			clickUser.childNodes[0].src = this.result;
		};
		reader.readAsDataURL(this.files[0]);
	}; // function finish;