// get query value in url from key
const urlParams = new URLSearchParams(window.location.search);

// GET ERROR MESSAGE FROM URL QUERY
const messageString = urlParams.get('message');
const messageType = urlParams.get('type');
urlParams.delete("message")
urlParams.delete("type")
if (messageString != null) {
  document.getElementById("message").style.display = 'grid';
  let message_title = document.getElementsByClassName("message_title")
  for (let i = 0; i < message_title.length; i++) {
    message_title[i].textContent = messageString
  }
  if (messageType == "error") {
    document.getElementById("error").style.display = 'flex';
  } else if (messageType == "success"){
    document.getElementById("success").style.display = 'flex';
  } else if (messageType == "warning"){
    document.getElementById("warning").style.display = 'flex';
  } else if (messageType == "info"){
    document.getElementById("info").style.display = 'flex';
  }
}

const message_close = document.getElementsByClassName("message_close")
for (let i = 0; i < message_close.length; i++) {
  message_close[i].onclick= function() {
    document.getElementById("message").style.animation = "moveOpen 1s reverse";
    setTimeout(function(){
      document.getElementById("message").style.display = 'none';
    }, 1000);
  }
}