console.log("Loading content script")

function updateColor(e) {
    // console.log("Got newColor event", e)
    try {
        let color = e.detail.substring(1,7);
        console.log("Setting color event", color)
        browser.runtime.sendMessage({color:color})
    } catch (err) {
        console.log("Caught someting ", err)
    }
}
window.addEventListener("newColor", updateColor)