
const serv = document.getElementById("server")

var req = getChannelList()
req.then((data) =>{
    if (data.CODE == 200) {
	serv.innerHTML = data.DATA[0].Name
    } else {
	serv.innerHTML = "It's a time"
    }
})
