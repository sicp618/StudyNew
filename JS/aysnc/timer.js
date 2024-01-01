// 你不知道的 JS 异步
function now() {
    return 21
}

function later() {
    answer = answer * 2
    console.log("Meaning of life:", answer)
}

let answer = now()
console.log("start:", answer)
setTimeout(later, 1 * 1000) // Meaning of life: 42

var a = {
    index: 1
}
console.log(a) // may be {index: 1} or {index: 2}
a.index++

// simple hanlder data, don't block
data = [1, 2, 3, 4, 5]
function response(data) {
    var chunk = data.splice(0, 1)
    console.log(chunk)
    if (data.length > 0) {
        setTimeout(function() {
            response(data)
        }, 1000)
    }
}
response(data)
