// A Simple javascript file to work with Sunsend servers
// If you wanna use sunsend insdie your website, use this file as your client
// For more information use `readme.md` file

var HOST = "http://127.0.0.1:3000/api/v1"
req_headers = {
    "API_KEY": "41FF6BED859798A91CF2BE33D9F89EACBA9F77693AB6EA2BF7A797F40328500C",
}

async function getChannelList() {
    return await fetch(HOST + "/list", { method: "GET", headers: req_headers})
    .then((response) => response.json())
    //.then((json) => console.log(json));
}

async function SendMsg(cid, user, msg, replyid, fileInput) {
    // curl --header "API_KEY: 41FF6BED859798A91CF2BE33D9F89EACBA9F77693AB6EA2BF7A797F40328500C" -s -X POST -d "user=$username&message=$message" http://127.0.0.1:3000/api/v1/chat/$t_roomID | jq
    const formData = new FormData();
    formData.append('image', fileInput.files[0]);
    formData.append('username', user);
    formData.append('message', msg);
    formData.append('reply', replyid);
    return await fetch(HOST + "/chat/" + cid, {
	method: "POST",
	headers: req_headers,
	/*
	body: JSON.stringify({
	    "username": user,
	    "message": msg,
	    "reply": replyid,
	}),
	*/
	body: formData,
    }).then((response) => response.json())
}

async function GetMsgs(cid, range) {
    var req_link_org = HOST + "/chat/" + cid
    if (range.length >= 1) {
	req_link_org = HOST + "/chat/" + cid + "?range=" + range
    }
    return await fetch(req_link_org, {
	    method: "GET",
	    headers: req_headers,
	}).then((response) => response.json())
}

// TODO: Load Images From Users's Messages

async function CreateChannel(name, desc, owner) {
    return await fetch(HOST + "/channel/create", {
	method: "POST",
	headers: req_headers,
	body: JSON.stringify({
	    "name": name,
	    "description": desc,
	    "owner": owner,
	}),
    }).then((response) => response.json())
}

