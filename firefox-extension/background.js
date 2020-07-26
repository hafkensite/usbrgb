/*
On startup, connect to the "ping_pong" app.
*/
var port = browser.runtime.connectNative("colors");

/*
Listen for messages from the app.
*/
port.onMessage.addListener((response) => {
  console.log("Received: " + response);
});

/*
On a click on the browser action, send the app a message.
*/
browser.browserAction.onClicked.addListener(() => {
  const colors = ["black", "red", "green", "blue", "teal", "aqua", "gray", "fuchsia"]
  let color = colors[Math.floor(Math.random() * colors.length)]
  // console.log("Sending:  ", color);
  queue.unshift(color);
});

var queue = [];


function handleMessage(request, sender, sendResponse) {
  // console.log("Message from the content script: " + request.color, queue.length);
  queue.unshift(request.color);
}

function sendColor() {
  c = queue.pop();
  if (c) {
    port.postMessage(c);
  }
  while (queue.length > 0) {
    queue.pop();
  }
}

setInterval(sendColor, 100);

browser.runtime.onMessage.addListener(handleMessage);