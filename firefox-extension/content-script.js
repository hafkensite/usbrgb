console.log("Loading content script")

function updateColor(color) {
    try {
        // let color = e.detail.substring(1,7);
        console.log("Setting color event", color)
        browser.runtime.sendMessage({ color: color })
    } catch (err) {
        console.log("Caught something", err)
    }
}


window.addEventListener('ringtonePlaying', function () { updateColor("003300"); });
// window.addEventListener('ringtoneStopped', function () { updateColor("000000"); });
window.addEventListener('callAccepted', function () { updateColor("110000"); });
window.addEventListener('callRejected', function () { updateColor("0000ff"); });
window.addEventListener('callTerminated', function () { updateColor("000000"); });
// window.addEventListener('callAccepted', ...);
// window.addEventListener('callRejected', ...);
// window.addEventListener('callTerminated', ...);

window.addEventListener("newColor", updateColor)